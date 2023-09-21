# 湖北大学校园网自动认证

## 1，介绍

这是一款由Go语言编写的小型命令行程序，能够自动地完成湖北大学校园网的登录认证操作。

目前仅支持Windows操作系统，后续可能会开发其它操作系统版本。

## 2，使用说明

在仓库页面右侧**发行版/Releases**处即可下载该程序的压缩包。

### (1) 前置要求

在使用该程序之前，首先需要保证你的电脑已经连接了湖北大学校园网`HUBU-STUDENT`或者`HUBU-WLAN`，才能够正常使用程序，建议将校园网设置为自动连接。

除此之外，该程序依赖Google Chrome浏览器，若电脑上没有安装Google Chrome，该程序也无法正常运行，可以到[官网](https://google.cn/chrome/)下载安装。

### (2) 配置

下载并解压缩，其中`config.cfg`就是该程序的配置文件，初始时其内容如下：

```
学号
密码
10
```

配置文件中：

- 第一行表示你的学号，需要**替换成你自己的学号**
- 第二行为校园网登录密码，同样需要**替换成你自己的校园网登录密码**
- 第三行为失败重试次数，可以保持默认也可以自己修改为一个不小于`0`的整数值

编写配置完成后，该程序即可使用。

### (3) 手动认证

双击可执行文件`hubu-wlan.exe`即可启动程序，程序就会帮助你完成认证操作，显示登录成功即可。

### (4) 自动认证（开机自启动）

下载的压缩包中除了程序和配置文件之外，还有`add-auto-start.bat`和`remove-auto-start.bat`这两个脚本，双击`add-auto-start.bat`即可将该程序添加至开机自启动列表，这样开机的时候程序就会自动启动完成校园网认证操作。

如果不想再开机自启了，则双击`remove-auto-start.bat`移除开机自启动项即可。

需要注意的是，添加开机自启动后，就不要再改变可执行文件和配置文件的所在位置了！否则会导致开机自启时找不到程序。