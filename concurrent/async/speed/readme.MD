# 测试异步的性能

测试异步调用方式的性能。

## 测试用例
### case1：BenchmarkGoroutine
直接启动新的goroutine去跑运行任务。

### case2:BenchmarkChan
启动一个后台goroutine消费，通过chan通信发送。

# 测试结果

| name | count | 耗时 |
| --- | ---- | ---- |
| BenchmarkGoroutine-8 | 10000000 | 229ns/op |
| BenchmarkChanNoBuffer-8 | 5000000 | 348ns/op |
| BenchmarkChanBuffer100-8 | 20000000 | 85.5ns/op |
| BenchmarkChanBuffer1000-8 | 20000000 | 64.6ns/op |
| BenchmarkChanBuffer10000-8 | 30000000 | 61.5ns/op |

区别是有没有buff和buff大小。