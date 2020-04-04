1.放在http://localhost/info.php也就是xampp根目录下

<?php
    echo phpinfo();
?>

从xampp自带的httpd.conf找到DocumentRoot属性，并且找到配置的根目录

2.windows下安装memcached（memcached.exe -d install），报错：Failed to ignore SIGHUP: No error
以管理员运行cmd，执行以下命令：
sc create "Memcached11211" binPath= "D:\memcached-amd64\memcached.exe -d runservice -p 11211" DisplayName= "Memcached11211" start= auto

3.phpstorm
ctrl+alt+left/right 浏览位置上一个和下一个

4.写api的时候参数校验至关重要

5.[如何在CentOS上安装phpMyAdmin](https://linux.cn/article-4526-1.html)

6.adminer.php（可视化数据库）直接放在项目跟目录下，例如/usr/local/var/www/BackApi/frontend/adminer.php

7.打印数组 echo json_encode($list);

8.PHP 也提供了另外一种方式给变量赋值：引用赋值。这意味着新的变量简单的引用（换言之，“成为其别名” 或者 “指向”）了原始变量。改动新的变量将影响到原始变量，反之亦然。
使用引用赋值，简单地将一个 & 符号加到将要赋值的变量前（源变量）。
PHP也有引用的概念

9.判断某字符串中是否包含某字符串的方法
if(strpos('www.idc-gz.com','idc-gz') !== false){
　　    echo '包含';
}else{
  echo '不包含';
}

10.psr2标准，mac上phpstorm支持：
https://oomusou.io/php/php-psr2/
https://www.jianshu.com/p/4cf8a1548b95
https://blog.csdn.net/hevenue/article/details/77679340

11.两张表查询：

$big_sql = 'select a.id,a.item_code,a.order_no,a.period_no,b.pick_number,b.prize_num,b.position,a.amount_per_note,a.note_count,b.unit,a.status,a.is_valid,b.mult_odds as odds,a.rakeback,a.rakeback_amount,a.sale_rakeback,a.win_note_count,a.win_amount,a.profit,a.is_follow,a.is_win_stop,a.created_at,a.amount,b.name as item_name,a.layout_code FROM ' . $order_table . ' a join ' . $order_slave_table . ' b on a.order_no = b.order_no where ' . $where
            . ' order by a.created_at desc, a.period_no desc, a.id desc limit ' . $index . ',' . $size ;
            
https://www.cnblogs.com/geaozhang/p/6753190.html
关于数据库的连表查询 

12.安装php7.2  https://newsn.net/say/centos-php72-yum.html

通过php -v 来查看php版本

13.https://wiki.swoole.com/wiki/page/479.html
swoole websocket服务器

14.linux服务器配置：（php开发）
https://www.cnblogs.com/liaojie970/p/6253404.html

https://blog.csdn.net/u011323949/article/details/73379146 （nginx）

https://blog.csdn.net/hu_zhe_kan/article/details/79368169   （安装php 7.2）

15.php Code Sniffer 与 php代码规范
https://blog.csdn.net/qq_32737755/article/details/97528394
https://docs.phpcomposer.com/00-intro.html
https://www.jianshu.com/p/5b697ccf80fc

//命名规范
//目录名称首字母大写，复数形式，多个单词以大写分割，如：/Models,/Controllers,/Services,/RedisQueues
 
//文件名称，以大写开头，资源集合类以复数形式，多个单词已大写分割，如：UsersResource.php, UsersService.php, UsersRepostiroy.php,
//对象实体类以单数形式，如：UserModel.php
 
//命名空间名称和文件名称一致，已大写字母开头，如：
App/Http/Resources
 
//类，一个文件一个类，类名应该与文件名称一致，如：UsersResource.php
class UsersResource
{
}
 
//变量名称以小写字母开头，单词已大写分割，变量名称禁止用 “中文拼音” 表示，禁止使用无意义的“简写”，如：$xinbie, $gp(goods Price),
//运算符前后需要一个空格隔开
//性别
$xinbie = 1;//错误的格式，1在代码中不可读
$gp = 20;//错误的格式，不能用无意义的简写，除了专业名称（goods price）
 
$prince='广东';//错误
$city  ='深圳';//错误，不需要为了等号对齐，添加不必要的空格
 
$gender = 'male';//正确
$goodsPrice = 20;//正确
$jwt = new JWT();//正确，json web token，专有名称
 
$province = '广东';//正确，赋值运算符前后以一个空格分割
$city = '深圳';//正确

//常用的命名格式
//分页
$total = 100;//总数据条数
$page = 1;//分页，页码，默认1
$pageSize = 18;//每页条数，默认18
 
//日期，时间
$startDate = '2018-10-12';//开始日期，精确到天，格式为YYYY-MM-DD
$endDate = '2018-10-13';//结束日期，精确到天，格式为YYYY-MM-DD
 
$startTime = '2018-10-12 12:14';//开始日期，精确到天，格式为YYYY-MM-DD HH:MM
$endTime = '2018-10-12 21:00';//开始日期，精确到天，格式为YYYY-MM-DD HH:MM
 
//count，和amount，次数和数量类型 $*Count, $*Amount
$loginCount = 10;
$depositAmount = 100.00;
 
//min, 和max, 最小/低 最大/高，类型$min*, $max*
$mixPrice = 0.00;//最低价格
$maxPrice = 100.00;//最高价格
 
//bool类型变量 $is*
//true ,false, null 小写，不能大写，如：
$isExist = false;
$isRemember = true;
$user = null;

//异常处理
//方法的返回值，除了true，false，null，需要的数据外，不能放回其他数据
//方法的返回值类型不能超过2种，如:
function myFuntion($option)
{  
    if (!$option) {
        return false;
    }
     
    if ($option == 'list') {
        return -1;//错误，和其他返回状态不统一
    }
    return true;
}
//一个方法里如果有多种异常情况，应该用异常类去处理，比如常用的HTTP请求
try {
    request('http://example.com');
} catch(HttpConnectException $e) {//连接不上异常
     
} catch(HttpRequestTimeoutException $e) {//超时异常
     
} //更多
