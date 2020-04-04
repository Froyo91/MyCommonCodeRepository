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

15.php Code Sniffer
https://blog.csdn.net/qq_32737755/article/details/97528394
https://docs.phpcomposer.com/00-intro.html
https://www.jianshu.com/p/5b697ccf80fc
