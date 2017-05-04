#!/bin/bash
P1=`ps -ef | grep java |grep scinan| grep -v grep | awk '{print $2}'`
P2=`ps -ef | grep java |grep scinan| grep -v grep | awk '{print $2}'| wc -l`
if [ "$P2" != "0" ]
then
echo $P1
kill -9 $P1
fi