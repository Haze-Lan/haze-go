package standard

//组件标准 TODO 后续使用
type Component interface {
	//初始化
	Init() error
	//销毁
	Destroy() error
	//中断运行
	Interrupt()
}
