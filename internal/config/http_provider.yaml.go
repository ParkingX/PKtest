package config

import "go.danale.net/be/be/go/srpc"

type HttpProvider struct {
	PvtServer *srpc.HttpServerConf `yaml:"pvt"`
}
