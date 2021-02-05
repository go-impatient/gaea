package xnsq

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
		Host: conf.Get("NSQ_HOST"),
		Port: conf.Get("NSQ_PORT"),
	}
}
