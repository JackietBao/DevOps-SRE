#!/usr/bin/bash
#批量创建脚本
for i in {1..100}
do 
	useradd JackietBao-$i
	echo "JackietBao-$i" is create ok ..
done