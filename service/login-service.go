package service

import (
	"context"
	"gitee.com/swsk33/sclog"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"hubu-wlan-connect/config"
	"strings"
	"time"
)

// 自动填写登录表单逻辑

// 清空全部输入框
func clearInput() chromedp.QueryAction {
	return chromedp.Evaluate(`
		const inputDOMs = document.querySelectorAll('input');
		for (let item of inputDOMs) {
			if (item.name === 'username' || item.name === 'password') {
				item.value = '';
			}
		}
	`, nil)
}

// 填写不同的值到不同的输入框
func fillInput(node *cdp.Node) chromedp.QueryAction {
	getName, _ := node.Attribute("name")
	if getName == "username" {
		return chromedp.SendKeys(node.FullXPath(), config.GlobalConfig.Username)
	} else {
		return chromedp.SendKeys(node.FullXPath(), config.GlobalConfig.Password)
	}
}

// 执行模拟登录操作
func mockLogin() bool {
	// 登录地址
	const url = "http://202.114.144.21/"
	// 两个输入框选择器
	const input = "#login-form .control-group .controls .input-block-level"
	// 登录按钮选择器
	const loginButton = "#login-form .row-fluid .span4 .btn-primary"
	// 注销按钮选择器
	const logoutButton = "#login-form .row-fluid .span4 .btn-danger"
	// 登录成功标题选取
	const successSelector = "#success-form .form-signin-heading"
	// 创建浏览器上下文
	timeoutContext, cancel1 := context.WithTimeout(context.Background(), 6*time.Second)
	webContext, cancel2 := chromedp.NewContext(timeoutContext)
	defer cancel1()
	defer cancel2()
	// 查找得到的输入框节点对象
	var inputNodes []*cdp.Node
	// 登录完成后的成功提示
	var successTip string
	// 开始执行
	// 先导航至登录页并等待页面加载完成，并获取用户名和密码框
	e := chromedp.Run(webContext, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("body"),
		chromedp.Nodes(input, &inputNodes),
	})
	if e != nil {
		sclog.Error("加载登录页时出现错误！%s\n", e)
		return false
	}
	// 最后填写账户密码并登录
	e = chromedp.Run(webContext, chromedp.Tasks{
		// 填写用户名和密码
		fillInput(inputNodes[0]),
		fillInput(inputNodes[1]),
		// 先点击注销按钮，否则重复登录会导致无法上网
		chromedp.Click(logoutButton, chromedp.NodeVisible),
		// 清空输入框
		clearInput(),
		// 再次填写用户名和密码
		fillInput(inputNodes[0]),
		fillInput(inputNodes[1]),
		// 点击登录按钮
		chromedp.Click(loginButton, chromedp.NodeVisible),
		// 获取成功字样
		chromedp.WaitVisible(successSelector, chromedp.ByQuery),
		chromedp.Text(successSelector, &successTip),
	})
	if e != nil {
		sclog.Error("进行登录时出现错误！%s\n", e)
		return false
	}
	return strings.Contains(successTip, "成功")
}

// DoLoginRetry 执行带重试的登录
func DoLoginRetry() {
	// 当前是否已登录成功
	success := false
	delay := config.GlobalConfig.Delay
	if delay > 0 {
		sclog.Info("将在%d秒后开始执行登录操作...\n", delay)
		time.Sleep(time.Duration(delay) * time.Second)
	}
	// 先执行一次登录操作
	sclog.InfoLine("开始执行登录操作...")
	success = mockLogin()
	// 如果失败则重试
	for count := 1; !success && count <= config.GlobalConfig.Retry; count++ {
		sclog.Warn("登录失败！正在进行第%d次重试...\n", count)
		success = mockLogin()
	}
	if success {
		sclog.InfoLine("登录成功！")
	} else {
		sclog.ErrorLine("登录失败！请确保已连接无线网络并且用户名密码正确！")
	}
}