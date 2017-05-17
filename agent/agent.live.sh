#!/bin/sh
#author yuchangfu
#date 2017年5月15日
#检测agent存活
date=`date +"%Y-%m-%d-%H-%M-%S"`
ip=`ifconfig | grep -oP '10(\.[0-9]{1,3}){3}'  | grep -v 255`
curl http://127.0.0.1:1988/health
if [ $? != 0 ]
then
cp /usr/local/open-falcon/agent/var/app.log /usr/local/open-falcon/agent/var/app.log.$date
curl -d "content= $ip agent failed ,please check log" "http://10.0.0.27:82/alert.php?type=weixin|rtx&user=yuchangfu"
/usr/local/open-falcon/agent/control start
fi
