#!/bin/bash

version=$1
if [ -z "$version" ]; then
	echo 请指定版本号！
	exit
fi

mkdir ./target
echo 清理上次构建...
rm -rf ./target/*
echo 开始构建程序...
cd ./src
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o ../target/hubu-wlan.exe
echo 开始压缩可执行文件...
cd ../target
upx -9 ./hubu-wlan.exe
echo 创建默认配置...
echo "学号" >./config.cfg
echo "密码" >>./config.cfg
echo "10" >>./config.cfg
echo 进行打包...
7z a -t7z -mx9 "./hubu-wlan-${version}.7z" ./hubu-wlan.exe ./config.cfg ../batch/*.bat
echo 构建完成！构建压缩包位于：target/hubu-wlan-${version}.7z
