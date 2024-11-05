# 湖北大学校园网自动认证

## 1，介绍

这是一款由Go语言编写的小型命令行程序，能够自动地完成湖北大学校园网的登录认证操作。

目前支持Windows和Linux操作系统。

## 2，使用说明

在仓库页面右侧**发行版/Releases**处即可下载该程序的压缩包，根据你自己的操作系统和架构选择对应的版本。

对于Debian系的操作系统，例如Ubuntu、Linux Mint、Deepin、Kali等等，还提供了`deb`格式的安装包，可以下载后安装：

```bash
sudo dpkg -i hubu-wlan-x.x.x-debian-xxx.deb
```

上述`hubu-wlan-x.x.x-debian-xxx.deb`替换成你自己下载的`deb`文件路径。

### (1) 前置要求

在使用该程序之前，首先需要保证你的电脑已经连接了湖北大学校园网`HUBU-STUDENT`或者`HUBU-WLAN`，才能够正常使用程序，建议将校园网设置为自动连接。

除此之外，该程序依赖Google Chrome浏览器，若电脑上没有安装Google Chrome，该程序也无法正常运行，可以到[官网](https://google.cn/chrome/)下载安装。

下载了`7z`或者`tar.xz`包并解压后，请将可执行文件所在目录加入到`Path`系统环境变量，使得在终端中可以直接调用`hubu-wlan`命令，配置环境变量后，可通过下列命令测试是否配置成功：

```bash
hubu-wlan version
```

如果能够输出版本号信息，说明配置环境变量成功。

如果是通过Linux的`deb`安装包安装，则无需配置环境变量，直接在终端里就可以调用该命令。

### (2) 自动补全脚本

下载解压后，除了命令程序本身`exe`文件，还有下列脚本文件，是用于命令自动补全的，将该脚本配置到对应的终端后，使用`hubu-wlan`命令时即可按下Tab键自动补全命令。

在Windows环境下：

- `hubu-wlan.fish` 用于Fish Shell的自动补全脚本，在Windows中通常是在Msys2环境下运行Fish Shell，将该文件放到`你的Msys2安装目录\etc\fish\completions`目录下即可，请勿修改文件名
- `hubu-wlan-completion.bash` 用于Bash Shell的自动补全脚本，在Windows中可以在Git Bash或者Msys2环境中运行Bash Shell，这里分别说明：
	- 使用Git Bash时，将`hubu-wlan-completion.bash`文件的扩展名改成`sh`，然后放到`你的Git安装目录\etc\profile.d`目录下
	- 使用Msys2时，直接把`hubu-wlan-completion.bash`文件放在`你的Msys2安装目录\etc\bash_completion.d`目录下

在Linux环境下，使用Bash Shell时，请先确保安装了`bash-completion`包：

```bash
sudo apt install bash-completion
```

然后把`hubu-wlan-completion.bash`放置到`/etc/bash_completion.d`目录下即可，若该目录不存在则创建。

在Linux下使用Fish Shell时，将`hubu-wlan.fish`直接放到`/etc/fish/completions`目录下即可，若该目录不存在则创建。

将文件放在对应位置后，配置就完成了，重启终端，后续在使用`hubu-wlan`命令时，即可使用Tab键补全命令。

> 使用`deb`安装包安装时会包含自动补全脚本配置，无需再手动复制补全脚本。

### (3) 配置

如果你下载的是Windows系统对应的`7z`压缩包或者是Linux系统的`tar.xz`包，那么只需下载并解压缩，在解压缩得到的文件中就包含名为`config.yaml`的文件，即为该程序的配置文件，初始时其内容如下：

```yaml
# 你的学号
username: "202300000000000"
# 校园网登录密码
password: "000000"
# 重试次数
retry: 15
# 初次启动延迟（秒）
delay: 3
```

配置文件中：

- `username` 表示你的学号，需要**替换成你自己的学号**，值需要用英文双引号`"`包围
- `password` 为校园网登录密码，同样需要**替换成你自己的校园网登录密码**，值需要用英文双引号`"`包围
- `retry` 为失败重试次数，可以保持默认也可以自己修改为一个不小于`0`的整数值
- `delay` 表示初次延迟秒数，防止开机自启时还没有连接网络就开始进行认证

如果你是下载的Linux的`deb`安装包，那么配置文件位于`/etc/hubu-wlan/config.yaml`，使用文本编辑器修改即可：

```bash
sudo vim /etc/hubu-wlan/config.yaml
```

编写配置完成后，该程序即可使用。

除了`deb`安装的情况，解压后的配置文件必须和可执行文件位于同一目录下！

### (4) 手动认证

修改完配置文件后，即可运行程序：

- Windows操作系统上，双击解压的可执行文件`hubu-wlan.exe`即可启动程序
- 如果是在Linux系统上并且你下载的是`tar.xz`压缩包，那么运行其中可执行文件`hubu-wlan`即可
- 如果是使用的`deb`安装包安装，则安装后可以直接打开终端执行`hubu-wlan`命令进行手动认证

启动程序后，程序就会帮助你完成认证操作，显示登录成功即可。

此外，也可以在终端中调用`hubu-wlan`命令一键登录：

```bash
hubu-wlan auth
```

### (5) 自动认证（开机自启动）

通过终端调用`hubu-wlan`进行操作，执行下列命令：

```bash
# 启用开机自启动
hubu-wlan auto-start add

# 禁用开机自启动
hubu-wlan auto-start remove
```

需要注意的是，添加开机自启动后，就不要再改变可执行文件和配置文件的所在位置了！否则会导致开机自启时找不到程序。

若启用了开机自启动，在Windows操作系统上开机时会显示一个命令行窗口并显示程序的执行状态，请勿手动关闭，在认证完成后程序会自动退出，在Linux系统上开机自启动是在后台运行的，因此你看不到程序运行的命令行窗口。

### (6) 卸载

如果你是通过`deb`安装包安装，可以使用下列命令进行卸载：

```bash
sudo apt remove hubu-wlan-connect
```

卸载之前，请先执行移除开机自启动操作！