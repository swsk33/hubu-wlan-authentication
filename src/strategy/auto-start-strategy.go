package strategy

// 不同操作系统下的自启动策略抽象

// AutoStart 开机自启动操作接口
type AutoStart interface {
	// AddAutoStart 添加开机自启动
	//
	// exePath 程序自身的路径
	//
	// 出现错误时返回错误对象
	AddAutoStart(exePath string) error
	// RemoveAutoStart 移除开机自启动
	//
	// 出现错误时返回错误对象
	RemoveAutoStart() error
}