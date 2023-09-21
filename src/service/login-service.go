package service

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/spf13/viper"
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
		return chromedp.SendKeys(node.FullXPath(), viper.GetString(config.USERNAME))
	} else {
		return chromedp.SendKeys(node.FullXPath(), viper.GetString(config.PASSWORD))
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
	timeoutContext, cancel1 := context.WithTimeout(context.Background(), 8*time.Second)
	webContext, cancel2 := chromedp.NewContext(timeoutContext)
	defer cancel1()
	defer cancel2()
	// 查找得到的输入框节点对象
	var inputNodes []*cdp.Node
	// 登录完成后的成功提示
	var successTip string
	// 开始执行
	// 先导航至登录页并等待页面加载完成，并获取用户名和密码框
	e1 := chromedp.Run(webContext, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible("body"),
		chromedp.Nodes(input, &inputNodes),
	})
	if e1 != nil {
		color.Red("加载登录页时出现错误！\n%s", e1)
		return false
	}
	// 最后填写账户密码并登录
	e2 := chromedp.Run(webContext, chromedp.Tasks{
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
	if e2 != nil {
		color.Red("进行登录时出现错误！\n%s", e2)
		return false
	}
	return strings.Contains(successTip, "成功")
}

// DoLoginRetry 执行带重试的登录
func DoLoginRetry() {
	// 最大重试次数
	maxCount := viper.GetInt(config.RETRY)
	// 当前是否已登录成功
	success := false
	// 执行延迟数
	delay := viper.GetInt(config.DELAY)
	if delay > 0 {
		color.Yellow("将在%d秒后开始执行登录操作...", delay)
		time.Sleep(time.Duration(delay) * time.Second)
	}
	// 先执行一次登录操作
	color.HiBlue("开始执行登录操作...")
	success = mockLogin()
	// 如果失败则重试
	for count := 1; !success && count <= maxCount; count++ {
		color.HiRed("登录失败！正在进行第%d次重试...", count)
		success = mockLogin()
	}
	if success {
		color.Green("登录成功！")
	} else {
		color.HiRed("登录失败！请确保已连接无线网络并且用户名密码正确！")
	}
}