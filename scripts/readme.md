
crontab -l 查看定时任务

crontab -e 编辑定时任务


```
* * * * * 命令
minute hour day month week command
分钟  小时  天  月    周   命令

分钟（0-59）
小时（0-23）
天（1-31）
月（1-12）
周（0-7）
```

#### 例子

```
* * */1 * * /usr/local/gopath/src/github.com/MDGSF/Blog/scripts/syncPostFromGithub.sh > /dev/null 2>&1 &
每天执行一次

00 00 * * * command
每天00:00执行一次

* * * * * command
每分钟执行一次

*/10 * * * * command 
每 10 分钟执行一次

3,15 * * * * command
每小时的第 3 分钟和第 15 分钟执行一次

3,15 8-11 * * * command 
在每天上午的 8 点到 11 点的第 3 分钟和第 15 分钟执行一次

3,15 8-11 */2 * * command 
每隔两天的上午的 8 点到 11 点的第 3 分钟和第 15 分钟执行一次

3,15 8-11 * * 1 command 
每个星期一的上午的 8 点到 11 点的第 3 分钟和第 15 分钟执行一次


* */1 * * * command
每小时执行一次

* 23-7/1 * * * command
在晚上11到早上7点之间，每小时执行一次

```



