@echo off

cd /d %~dp0
cd ..
set JAVA_OPTS=-Dfile.encoding="UTF-8"
java -Dfile.encoding=UTF-8 -Xms256m -Xmx1024m -jar start.jar
