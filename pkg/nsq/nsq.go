package nsq

import (
	"moocss.com/gaea/pkg/conf"
)

type Nsq struct {
	Topic string
	Host  string
	Port  string
}

func NewNsq() Nsq {
	return Nsq{
		Host: conf.Get("features.nsq.host"),
		Port: conf.Get("features.nsq.port"),
	}
}
