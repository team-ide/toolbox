#!/bin/sh
cd `dirname $0`
cd ..
THIS_HOME=`pwd`
LOGS_DIR=$THIS_HOME/logs

if [ ! -d $LOGS_DIR ]; then
    mkdir $LOGS_DIR
fi
PID_FILE=$THIS_HOME/bin/start.pid
START_LOG_FILE=$LOGS_DIR/start.log

echo -e "Starting Team IDE Toolbox \c"
nohup java -Dfile.encoding=UTF-8 -Xms256m -Xmx1024m -jar $TEAMIDE_HOME/start.jar >$START_LOG_FILE 2>&1 & echo $! > $PID_FILE
sleep 1
echo "Started Team IDE Toolbox "

