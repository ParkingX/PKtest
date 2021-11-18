package config

import "go.danale.net/be/be/go/srpc"

type SrpcProvider struct {
	DipDemoPvtSvc srpc.SrpcServerConf `yaml:"demo.DipDemoPvtSvc"`
}
