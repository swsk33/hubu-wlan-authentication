#!/bin/bash

# 检查命令行参数
version=$1
if [ -z "$version" ]; then
	echo 请指定版本号！
	exit
fi

# 创建文件夹并清理
mkdir -p ./target
echo 清理上次构建...
rm -rf ./target/*

# 进行资源文件和可执行文件构建
cd ./src
echo 开始构建资源...
go-winres make --in "../winres/winres.json"
echo 开始构建程序...
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o "../target/hubu-wlan.exe"
rm ./*.syso

# 压缩可执行文件
cd ../target
echo 开始压缩可执行文件...
upx -9 ./hubu-wlan.exe

# 创建配置文件
echo 创建默认配置...
echo "# 你的学号" >./config.yaml
echo "username: \"202300000000000\"" >>./config.yaml
echo "# 校园网登录密码" >>./config.yaml
echo "password: \"000000\"" >>./config.yaml
echo "# 重试次数" >>./config.yaml
echo "retry: 15" >>./config.yaml
echo "# 初次启动延迟（秒）" >>./config.yaml
echo "delay: 3" >>./config.yaml

# 打包
echo 进行打包...
7z a -t7z -mx9 "./hubu-wlan-${version}.7z" ./hubu-wlan.exe ./config.yaml ../batch/*.bat
echo 构建完成！构建压缩包位于：target/hubu-wlan-${version}.7z
