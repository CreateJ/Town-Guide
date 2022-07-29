package base

import (
	"github.com/tietang/props/kvs"
	"reflect"
)

//应用程序
type BootApplication struct {
	IsTest     bool
	conf       kvs.ConfigSource
	starterCtx StarterContext
}

//构造系统
func New(conf kvs.ConfigSource) *BootApplication {
	e := &BootApplication{conf: conf, starterCtx: StarterContext{}}
	e.starterCtx.SetProps(conf)
	return e
}

func (b *BootApplication) Start() {
	//1. 初始化starter
	b.init()
	//2. 安装starter
	b.setup()
	//3. 启动starter
	b.start()
}

//程序初始化
func (e *BootApplication) init() {
	for _, v := range GetStarters() {
		_ = reflect.TypeOf(v)
		v.Init(e.starterCtx)
	}
}

//程序安装
func (e *BootApplication) setup() {

	for _, v := range GetStarters() {
		_ = reflect.TypeOf(v)
		v.Setup(e.starterCtx)
	}

}

//程序开始运行，开始接受调用
func (e *BootApplication) start() {
	for i, v := range GetStarters() {

		_ = reflect.TypeOf(v)
		if v.StartBlocking() {
			if i+1 == len(GetStarters()) {
				v.Start(e.starterCtx)
			} else {
				go v.Start(e.starterCtx)
			}
		} else {
			v.Start(e.starterCtx)
		}

	}
}

//程序开始运行，开始接受调用
func (e *BootApplication) Stop() {

	for _, v := range GetStarters() {
		_ = reflect.TypeOf(v)
		v.Stop(e.starterCtx)
	}
}
