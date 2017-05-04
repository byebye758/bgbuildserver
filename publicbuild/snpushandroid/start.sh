#!/bin/sh
#CLASSPATH="$JAVA_HOME"/lib/
JAVA_OPTS="-verbose:gc -XX:+PrintGCDetails  -Xdebug -Xss512k -Xloggc:gc.log -Xms6240m -Xmx6240m -XX:PermSize=128m -XX:MaxPermSize=128m  -XX:NewSize=1524m -XX:MaxNewSize=1524m -XX:SurvivorRatio=8 -XX:ParallelGCThreads=20 -XX:CMSFullGCsBeforeCompaction=5 -XX:CMSInitiatingOccupancyFraction=70 -XX:MaxTenuringThreshold=10  -XX:+UseConcMarkSweepGC -XX:+UseParNewGC -Duser.timezone=GMT+8 -Dfile.encoding=UTF8"
#JAVA_OPTS="$JAVA_OPTS -Djava.rmi.server.hostname=192.168.10.107 -Dcom.sun.management.jmxremote -Dcom.sun.management.jmxremote.port=12345 -Dcom.sun.management.jmxremote.ssl=false -Dcom.sun.management.jmxremote.authenticate=true" 
S8100_HOME=`pwd`
CLASSES="$S8100_HOME"/bin
LIB="$S8100_HOME"/lib
echo "Using Lib:   $LIB"
MAINCLASS=bin/pushAndroid.jar
LOGBACK=-Dlogback.configurationFile=bin/logback.xml
MAINARGS=$1
CLASSPATH=.:"$CLASSES"
echo "Using CLASSPATH:  $CLASSPATH"
echo "Using MAINCLASS:  $MAINCLASS"
echo "Using JAVA:       $_RUNJAVA"
echo "Using JAVA_OPTS:  $JAVA_OPTS"
mkdir -p /var/log/scinan/snpushandroid
NIOServer_OUT=/var/log/scinan/snpushandroid/Server.`date +%Y%m%d`.out
touch "$NIOServer_OUT"
pid=`ps -ef | grep -i $MAINCLASS | grep -v grep | awk '{print $2}'| xargs`
if [ "$pid" != ""  ]; then
        echo "have pid, will restart"
        kill -9 $pid     
fi
java $JAVA_OPTS -jar $LOGBACK $MAINCLASS $MAINARGS  >> "$NIOServer_OUT" 2>&1

