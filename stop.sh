#!/usr/bin bash

ps aux | grep -v grep | grep spamcheck | awk '{print $2}' | xargs kill -9
rm /tmp/spamcheck.sock
echo "spamcheck服务已关闭"
