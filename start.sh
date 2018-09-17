#!/usr/bin bash
PWD=`pwd`
cd $PWD
# 检查是否已经开启过spamcheck服务
PID=`ps aux | grep -v grep | grep spamcheck | awk '{print $2}'`
if [ "$PID" != "" ];then
    echo "服务已开启，不能重复启动。"
    exit -1
fi
# 开启服务
nohup go run spamcheck.go -c config.json -a sock >> nohup.out &
