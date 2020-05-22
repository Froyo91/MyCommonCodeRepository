1.defer 是后进先出的。

2.sync.waitgroup
我经常在main里边直接写协程的测试demo,main退出会结束主协程，之后会强制结束子协程

3.正向代理隐藏真实客户端，反向代理隐藏真实服务端。
先说正向代理，比如要访问youtube,但是不能直接访问，只能先找个翻墙软件，通过翻墙软件才能访问youtube. 翻墙软件就叫做正向代理。所谓的反向代理，指的是用户要访问youtube,但是youtube悄悄地把这个请求交给x-art来做，那么x-art就是反向代理了。

4.goland要配置下go安装路径。

5.golang代码规范
包名应为小写单词
不应有下划线或者混合大小写。正确示例 controllers, models, routers, views
全局变量即参数采用驼峰式命名
全局变量：驼峰式，首字母大写（如果不可导出，则首字母小写）
参数传递：驼峰式，首字母小写
采用全部大写或者全部小写来表示缩写单词
比如对于url这个单词，不要使用Url，而要使用url或者URL
单个函数的接口名以 er 为后缀
如Reader, Writer，其具体的实现则去掉 er，如
type Reader interface {
Read(p []byte) (n int, err error)
}
两个函数的接口名综合两个函数名，如
type WriteFlusher interface {
Write([]byte) (int, error)
Flush() error
}
三个以上函数的接口名类似于结构体名，如
type Server interface {
Start()
Stop() error
Receive()
}
流程控制
if 接受初始化语句
if err := file.Chmod(0664); err != nil {
return err
}
for 采用短声明建立局部变量
sum := 0
for i := 0; i < 10; i++ {
sum += i
}
range 如果只需要第一项（key），就丢弃第二个
for key := range m {
if key.expired() {
delete(m, key)
}
}
range 如果只需要第二项，则把第一项置为下划线
sum := 0
for _, value := range array {
sum += value
}
错误处理
一旦有错误发生，返回，外层调用对错误进行处理，对于错误处理 2.0版本会有很会的改进
f, err := os.Open(name)
if err != nil {
return err
}
不要滥用 panic 
不要使用 panic 来做正常的错误处理，应当使用 error 和 多个返回值来进行

不要忽略错误 
如果一个函数的返回值包括 err，那么不要使用 _ 来忽略它，而应该去检查函数是否执行成功，如果不成功则执行对应的错误处理并返回，只有在确实不希望出现的情况下才使用 panic

无论是参数列表还是返回值，最好加上名称，方便理解（尤其是在有同类型的多个参数的时候） 
比如

func (f *Foo) Location() (float64, float64, error)
就不如

func (f *Foo) Location() (lat, long float64, err error)
来得清晰

参数传递
对于大量数据的struct可以考虑使用指针
传入参数是map，slice，chan不要传递指针，因为map，slice，chan是引用类型，一般情况不需要传递指针的指针
注释
每个包都应该有一个包注释，位于 package 之前
/*
* @ Author:
* @ For: 做什么
* @ Date: 2018年10月8日
* @ Des: 详细描叙
*/
package models
每个以大写字母开头（即可以导出）的方法应该有注释，大写表示是开放的 对外可调用

单行注释使用 //, 多行注释使用 /* xxx */

程序中每一个被导出的（大写的）名字，都应该有一个文档注释

注释必须是完整的句子，以需要注释的内容作为开头，句点作为结尾

行长
一行最长不超过80~100个字符，超过的使用换行展示，尽量保持格式优雅，方便阅读。
IDE推进
Visual Studio Code 或者 Goland
声明切片
使用 var t []string 而不是 t := []string{} 来声明一个切片，好处是切片从来没有使用的话，就不会对其分配内存
常量
常量均需使用全部大写字母组成,或者驼峰首字母大写
package级的Error变量，通常会把自定义的Error放在package级别中，统一进行维护:
var (
ErrCacheMiss = errors.New("memcache: cache miss")
ErrCASConflict = errors.New("memcache: compare-and-swap conflict")
ErrNotStored = errors.New("memcache: item not stored")
)
常量 状态码也是一样的代理

Imports
当import多个包时，应该对包进行分组。同一组的包之间不需要有空行，不同组之间的包需要一个空行。标准库的包应该放在第一组。
package main
import (
"fmt"
"hash/adler32"
"os"
"appengine/foo"
"appengine/user"
"code.google.com/p/x/y"
"github.com/foo/bar"
)
断言,Interface类型一定要做防御处理
var target int
if interInt, ok := inter.(int); ok {
}
不要写成

interInt, _ := inter.(int)
直接就调用转化后的interInt，这是不安全的，特别是Interface到string的处理

复制slice
var b1, b2 []byte
for i, v := range b1 {
b2[i] = v
}
for i := range b1 {
b2[i] = b1[i]
}
使用内建的copy函数：copy(b2, b1)

正则表达式中使用raw字符串避免转义字符
regexp.MustCompile("\.") 
而是直接使用raw字符串，可以避免大量的\出现 
regexp.MustCompile(\.)

单元测试
单元测试文件名命名规范为 example_test.go
测试用例的函数名称必须以 Test 开头，例如：TestExample
对于单元测试，go 天生就友好的支持,一些工具库都要写单元测试代码测试 go test 去测试
包管理
需要安装 golang 版本为 go1.11.1 以及以上版本
使用Go自带的 go module和 go get 进行包的管理
使用 vendor 存放第三方库 
初始化一个包，例如项目名为 testgo
cd xxxx/testgo
go mod init testgo
vendor 存放第三方库

go mod vendor
导入新的第三方库

go mod tidy
Recover
recover用于捕获runtime的异常，禁止滥用recover，在开发测试阶段尽量不要用recover 
recover一般放在有不可预期的异常地方
func server(workChan <-chan *Work) {
for work := range workChan {
go safelyDo(work)
}
}
func safelyDo(work *Work) {
defer func() {
if err := recover(); err != nil {
log.Println("work failed:", err)
}
}()
// do 函数可能会有不可预期的异常
do(work)
}
项目的代码结构
分层架构的一个重要原则是：每层只能与位于其下方的层发生耦合。严格分层架构，某层只能与直接位于其下方的层发生耦合 
还有面向接口编程 多用interface
PROJECT_NAME 
├── README.md 介绍软件及文档入口,程序运行方式，部署环境 
├── changehistory.txt 每次版本迭代更新的更新记录，比如配置，DB字段修改，防止部署时候遗漏 
├── bin 编译好的二进制文件,执行./build.sh自动生成，该目录也用于程序打包 
├── build.sh 自动编译的脚本 
├── doc 该项目的文档 
├── Dockerfile Docker部署 
├── controllers 层 
├── login 登录模块 
├──loginctl.go 登录控制器 
├── 表报 表报模块 
├──reportctl.go 表报控制器 
├── service 业务层 根据项目的复杂程度，简单的可以只是MVC结构 复杂的可以切换到 DDD(领域驱动设计) 
├── repository 层 根据项目的复杂程度，简单的可以只是MVC结构 复杂的可以切换到 DDD(领域驱动设计) 
├── middleware 中间件 统一管理 
├── Model层 model层 对应一个表结构，结构体定义 
├── conf 层 配置文件 
├── cache 层 缓存层，统一管理 
├── logs 本地日志代码 
├── es 基于es的日志 
├── file 基于文件流的日志 
├── router 
├── grpc grpc 路由层 
├── api http 路由层 
├── utils 常用的工具库 封装，减少耦合 代码复用，与业务无关的写到这里 
├── vendor 存放go的库 利用go module来管理package

代码分析工具
工具 
demo地址 
* 包括代码行数 
* 单元测试覆盖量 
* 代码质量分析

API 接口定义
http 状态码

6000	正常返回
6001	表示登录状态失效（类似401）
6002	表示账号或密码错误
6003	表示访问被拒绝（类似403）
6004	表示找不到接口地址（类似404）
6005	表示接口服务器产生错误（类似500）
6006	表示无效网关（网关与接口服务器通讯失败，类似502）
6007	表示接口请求失败（例如表单提交内容非法，后台无法接收）
6008	表示违法的参数
6009	表示数据已经存在
6010	表示文件上传失败
规范要素
controller 层 MVC结构 负责组装各层的业务，参数检验，鉴权，都封装好
代码要注重可复用
状态的枚举 两种方式要有其一：第一种，就是在数据返回字段加一个状态字段描叙，第二种，就是在接口文档里面说明清楚
缓存写到cache层，admin/cache里面
http 返回状态码 定义枚举常量，不要直接在代码写数字
返回给前端的枚举值 要有对应的描叙des字段
接口返回数据遵守单一 最简原则，不要把不需要的数据要一起返回
分页业务代码规范
参数说明
page	请求的的页码，为空时默认处理成1
pageSize	每页的数据条数，为空是默认为10
Json统一结构返回

{
"status" : "6000",
"message" : "success"
"result" : {
"data" : [],
"total" : 0,//总的记录数，int类型
"page"： 1,//当前页数
"pageSize":18//每页的条数
}
}
常用的时间参数
参数说明
startDate	开始日期，精确到天，格式YYYY-MM-DD ，如2018-10-11
endDate	结束日期，精确到天，格式YYYY-MM-DD ，如2018-10-12
startTime	开始时间，精确到分，格式YYYY-MM-DD HH:MM ，如2018-10-11 15:45
endTime	结束时间，精确到分，格式YYYY-MM-DD HH:MM ，如2018-10-11 18:45

6.下载依赖库
执行 go mod tidy

7.json字符串转map（参考）
//对json数据进行处理，把不符合当前游戏类型的游戏都剔除
	jsonData := out.(map[string]interface{})
	for k, v := range jsonData {
		temp := v.([]interface{})
		if len(temp) > 0 {
			seq := make([]interface{}, 0)
			for _, v1 := range temp {
				temp2 := v1.(map[string]interface{})
				if temp2["game_type"] == gameType {
					seq = append(seq, temp2)
				}
			}
			if len(seq) > 0 {
				jsonData[k] = seq
			} else {
				delete(jsonData, k)
			}
		}
	}

8.go flag参数
--game=./game_test.json --config=./config_test.json --gamename=./GameNameMap_test.json
