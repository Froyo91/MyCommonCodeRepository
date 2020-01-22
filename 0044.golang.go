1.defer 是后进先出的。

2.sync.waitgroup
我经常在main里边直接写协程的测试demo,main退出会结束主协程，之后会强制结束子协程
