package base

import (
	"bufio"
	"bytes"
	"centralData/db/mdata"
	"centralData/logs"
	userser "centralData/service/user"
	"centralData/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const versionOne = "/central/v1" //这里是当前项目的路由

func HttpHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		routeArr := strings.Split(c.Request.URL.Path, "/")
		if len(routeArr) > 3 { //符合条件的路由为：/central/v1/fds/***
			if routeArr[3] == mdata.AdminServerHost ||
				routeArr[3] == mdata.AgentServerHost ||
				routeArr[3] == mdata.ActiveSySConfigHost ||
				routeArr[3] == mdata.FdAdminHost {
				platform := ""
				var postData []byte
				var err error
				if c.Request.Method == "GET" {
					platform = c.Query("platform")
					if platform == "" {
						WebResp(c, 200, utils.ErrInvalidParams, nil, "platform为空")
						return
					}
				} else if c.Request.Method == "POST" {
					//还要判断是json格式还是form-data格式
					if strings.ToLower(c.Request.Header.Get("Content-Type")) == "application/json" { //json格式
						postData, err = c.GetRawData()
						if err != nil {
							WebResp(c, 200, utils.ErrInvalidParams, nil, err.Error())
							return
						}

						platformData := mdata.Platform{}
						err = mdata.Cjson.Unmarshal(postData, &platformData)
						if err != nil {
							WebResp(c, 200, utils.ErrInvalidParams, nil, "platform配置有问题")
							return
						}
						platform = platformData.Platform
						if platform == "" {
							WebResp(c, 200, utils.ErrInvalidParams, nil, "platform为空")
							return
						}
					} else { //form-data格式
						platform = c.PostForm("platform")
						if platform == "" {
							WebResp(c, 200, utils.ErrInvalidParams, nil, "platform为空")
							return
						}
					}
				}

				if platform != "" {
					proxyAddr := ""
					pathRoute := "" //这里的路由可以根据实际情况进行配置
					switch routeArr[3] {
					case mdata.AdminServerHost:
						proxyAddr = mdata.DomainConfigs[platform].AdminServer
						pathRoute = "/admin/v2"
					case mdata.AgentServerHost:
						proxyAddr = mdata.DomainConfigs[platform].AgentServer
						pathRoute = "/admin/v2/agent"
					case mdata.ActiveSySConfigHost:
						proxyAddr = mdata.DomainConfigs[platform].ActiveSySConfig
						pathRoute = "/activity/api/v3"
					case mdata.FdAdminHost:
						proxyAddr = mdata.DomainConfigs[platform].FdAdmin
						pathRoute = "/fd/v1"
					}

					proxy, err := url.Parse(proxyAddr)
					if err != nil {
						glog.Error(err)
						WebResp(c, 200, utils.ErrInternal, nil, utils.Error)
						return
					}
					c.Request.URL.Scheme = proxy.Scheme
					c.Request.URL.Host = proxy.Host
					c.Request.URL.Path = pathRoute + strings.ReplaceAll(c.Request.URL.Path, versionOne+"/"+routeArr[3], "")

					if c.Request.Method == "POST" { //这里要把操作者和ip带进参数;还要判断是json格式还是form-data格式
						tokenData, err := userser.ParseToken(c.GetHeader(mdata.HeaderToken))
						if err != nil {
							WebResp(c, 200, utils.ErrInvalidParams, nil, "token失效，请重新登录")
							return
						}
						if strings.ToLower(c.Request.Header.Get("Content-Type")) == "application/json" { //json格式
							var finalPostData map[string]interface{}
							err = mdata.Cjson.Unmarshal(postData, &finalPostData)
							if err != nil {
								WebResp(c, 200, utils.ErrInvalidParams, nil, utils.Error)
								return
							}
							finalPostData["bigDataUserName"] = tokenData.Name
							finalPostData["bigDataIp"] = c.ClientIP()

							jsonString, err := mdata.Cjson.Marshal(finalPostData)
							if err != nil {
								WebResp(c, 200, utils.ErrInvalidParams, nil, utils.Error)
								return
							}

							c.Request.ContentLength = int64(len(jsonString))
							c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(jsonString))
						} else {

						}
					}

					transport := http.DefaultTransport
					resp, err := transport.RoundTrip(c.Request)
					if err != nil {
						glog.Error(err)
						WebResp(c, 200, utils.ErrInternal, nil, utils.Error)
						return
					}

					// 将响应结果返回
					for key, value := range resp.Header {
						for _, v := range value {
							c.Writer.Header().Add(key, v)
						}
					}
					defer resp.Body.Close()
					bufio.NewReader(resp.Body).WriteTo(c.Writer)
				}
			}
		}
	}
}
