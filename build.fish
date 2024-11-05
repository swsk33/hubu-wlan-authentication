#!/bin/fish

# 检查命令行参数
set app_version $argv[1]

if test -z "$app_version"
    echo 请指定版本号！
    exit
end

# 基本文件名等
set basename hubu-wlan
set base_output target

# 清理上次构建
if test -d $base_output
    echo 清理上次构建...
    rm -r $base_output
end

mkdir -p $base_output

# 构建资源
echo 开始构建资源...
go-winres make --in "./winres/winres.json"

# 生成自动补全脚本
set script_output $base_output/script
mkdir -p $script_output
echo 开始构建自动补全脚本...
go run main.go completion bash >$script_output/$basename-completion.bash
go run main.go completion fish >$script_output/$basename.fish
sed -i 1d $script_output/$basename-completion.bash
sed -i 1d $script_output/$basename.fish

# 构建Windows程序
set win_i386_output $base_output/win/i386
set win_amd64_output $base_output/win/amd64
mkdir -p $win_i386_output
mkdir -p $win_amd64_output
echo 正在构建Windows i386...
GOOS=windows GOARCH=386 go build -ldflags "-w -s" -o "$win_i386_output/$basename.exe"
echo 正在构建Windows amd64...
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o "$win_amd64_output/$basename.exe"
rm *.syso

# 构建Linux程序
set linux_i386_output $base_output/linux/i386
set linux_amd64_output $base_output/linux/amd64
mkdir -p $linux_i386_output
mkdir -p $linux_amd64_output
echo 正在构建Linux i386...
GOOS=linux GOARCH=386 go build -ldflags "-w -s" -o "$linux_i386_output/$basename"
echo 正在构建Linux amd64...
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o "$linux_amd64_output/$basename"

# 压缩可执行文件
echo 开始压缩可执行文件...
upx -9 $win_i386_output/$basename.exe
upx -9 $win_amd64_output/$basename.exe
upx -9 $linux_i386_output/$basename
upx -9 $linux_amd64_output/$basename

# 创建配置文件
set config_output_file $base_output/config/config.yaml
mkdir -p (dirname $config_output_file)
echo 创建默认配置...
echo "# 你的学号" >"$config_output_file"
echo "username: \"202300000000000\"" >>"$config_output_file"
echo "# 校园网登录密码" >>"$config_output_file"
echo "password: \"000000\"" >>"$config_output_file"
echo "# 重试次数" >>"$config_output_file"
echo "retry: 15" >>"$config_output_file"
echo "# 初次启动延迟（秒）" >>"$config_output_file"
echo "delay: 3" >>"$config_output_file"

# 打包
echo 进行打包...
7z a -t7z -mx9 $base_output/$basename-$app_version-windows-i386.7z ./$win_i386_output/$basename.exe ./$config_output_file ./$script_output/*
7z a -t7z -mx9 $base_output/$basename-$app_version-windows-amd64.7z ./$win_amd64_output/$basename.exe ./$config_output_file ./$script_output/*
7z a -ttar $base_output/$basename-$app_version-linux-i386.tar ./$linux_i386_output/$basename ./$config_output_file ./$script_output/*
7z a -ttar $base_output/$basename-$app_version-linux-amd64.tar ./$linux_amd64_output/$basename ./$config_output_file ./$script_output/*
xz -z -e $base_output/$basename-$app_version-linux-i386.tar
xz -z -e $base_output/$basename-$app_version-linux-amd64.tar

# 构建deb安装包
# 参数1：构建架构，可以是i386或者amd64
function build_deb
    set arch $argv[1]
    echo 构建Debian {$arch}安装包...
    mkdir -p ./deb-build/$arch/opt/hubu-wlan-connect
    mkdir -p ./deb-build/$arch/etc/hubu-wlan
    mkdir -p ./deb-build/$arch/usr/bin
    set exe_path
    if test "$arch" = i386
        set exe_path $linux_i386_output/$basename
    else
        set exe_path $linux_amd64_output/$basename
    end
    cp -f $exe_path ./deb-build/$arch/opt/hubu-wlan-connect/
    cp -f ./winres/icon.png ./deb-build/$arch/opt/hubu-wlan-connect/
    cp -f $config_output_file ./deb-build/$arch/etc/hubu-wlan/
    set link_path ./deb-build/$arch/usr/bin/hubu-wlan
    if test -f $link_path
        rm $link_path
    end
    ln -s /opt/hubu-wlan-connect/$basename $link_path
    dpkg -b ./deb-build/$arch/ $base_output/$basename-$app_version-debian-$arch.deb
end

build_deb i386
build_deb amd64

echo 构建完成！结果位于{$base_output}目录下

# 清理构建
echo 清理构建...
rm -r $script_output
rm -r (dirname $config_output_file)
rm -r ./target/win
rm -r ./target/linux
