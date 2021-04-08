package logger

import (
	"fmt"
	"os"
)

//Logger生成器
type Factory struct {
	name string
}

var cache = map[string]*Factory{}

//实现GRPC LoggerV2  接口

func (c *Factory) Info(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Warning(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Error(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Fatal(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (c *Factory) Warningf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (c *Factory) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (c *Factory) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args)
}

func (c *Factory) Infoln(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Warningln(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Errorln(args ...interface{}) {
	fmt.Println(args)
}

func (c *Factory) Fatalln(args ...interface{}) {
	fmt.Println(args)
	os.Exit(0)
}

func (c *Factory) V(l int) bool {
	return c.V(l)
}

func LoggerFactory(name string) *Factory {

	if cData, ok := cache[name]; ok {
		return cData
	}
	c := &Factory{name}
	cache[name] = c
	return c
}
