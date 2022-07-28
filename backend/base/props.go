package base

import (
	log "github.com/sirupsen/logrus"
	"github.com/tietang/props/kvs"
)

var props kvs.ConfigSource

func Props() kvs.ConfigSource {
	Check(props)
	return props
}

type PropsStarter struct {
	BaseStarter
}

func (p *PropsStarter) Init(ctx StarterContext) {
	props = ctx.Props()
	log.Info("初始化配置.")
}
