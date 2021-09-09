#!/bin/sh
cd `dirname $0`
cd ..
THIS_HOME=`pwd`

PID_FILE=$THIS_HOME/bin/start.pid

echo -e "Stopping Team IDE Toolbox \c"
sudo kill `cat $PID_FILE`  
    rm -rf $PID_FILE
sleep 1
echo "Stopped Team IDE Toolbox "

