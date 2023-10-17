#!/bin/bash

# 检查命令行参数
version=$1
if [ -z "$version" ]; then
	echo 请指定版本号！
	exit
fi

basename=hubu-wlan

# 创建文件夹并清理
mkdir -p ./target
echo 清理上次构建...
rm -rf ./target/*

# 进行资源文件和可执行文件构建
cd ./src
echo 开始构建资源...
go-winres make --in "../winres/winres.json"

# 构建Windows程序
mkdir -p ../target/win/i386
mkdir -p ../target/win/amd64
echo 正在构建Windows i386...
GOOS=windows GOARCH=386 go build -ldflags "-w -s" -o "../target/win/i386/${basename}.exe"
echo 正在构建Windows amd64...
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o "../target/win/amd64/${basename}.exe"
rm ./*.syso

# 构建Linux程序
mkdir -p ../target/linux/i386
mkdir -p ../target/linux/amd64
echo 正在构建Linux i386...
GOOS=linux GOARCH=386 go build -ldflags "-w -s" -o "../target/linux/i386/${basename}"
echo 正在构建Linux amd64...
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o "../target/linux/amd64/${basename}"

# 压缩可执行文件
cd ../target
echo 开始压缩可执行文件...
upx -9 ./win/i386/${basename}.exe
upx -9 ./win/amd64/${basename}.exe
upx -9 ./linux/i386/${basename}
upx -9 ./linux/amd64/${basename}

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
mkdir -p ./out
echo 进行打包...
7z a -t7z -mx9 "./out/${basename}-${version}-windows-i386.7z" ./win/i386/${basename}.exe ./config.yaml ../batch/*.bat
7z a -t7z -mx9 "./out/${basename}-${version}-windows-amd64.7z" ./win/amd64/${basename}.exe ./config.yaml ../batch/*.bat
tar -cJvf "./out/${basename}-${version}-linux-i386.tar.xz" --transform "s=linux/i386/==" ./linux/i386/${basename} ./config.yaml
tar -cJvf "./out/${basename}-${version}-linux-amd64.tar.xz" --transform "s=linux/amd64/==" ./linux/amd64/${basename} ./config.yaml

# 构建deb安装包
# 参数1：构建架构，可以是i386或者amd64
function build_deb() {
	echo 构建Debian ${1}安装包...
	mkdir -p ../deb-build/${1}/opt/hubu-wlan-connect
	mkdir -p ../deb-build/${1}/etc/hubu-wlan
	mkdir -p ../deb-build/${1}/usr/bin
	cp -f ./linux/${1}/${basename} ../deb-build/${1}/opt/hubu-wlan-connect/
	cp -f ../winres/icon.png ../deb-build/${1}/opt/hubu-wlan-connect/
	cp -f ./config.yaml ../deb-build/${1}/etc/hubu-wlan/
	rm ../deb-build/${1}/usr/bin/hubu-wlan
	ln -s /opt/hubu-wlan-connect/${basename} ../deb-build/${1}/usr/bin/hubu-wlan
	dpkg -b ../deb-build/${1}/ ./out/${basename}-${version}-debian-${1}.deb
}

build_deb i386
build_deb amd64

echo 构建完成！结果位于./target/out目录下
