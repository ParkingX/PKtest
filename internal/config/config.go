package config

import (
	"fmt"

	"go.danale.net/be/be/go/ddas-kit/module/dconf"
	"go.danale.net/be/be/go/logx"
	"go.danale.net/be/be/go/srpc"

	"gopkg.in/yaml.v2"
)

func init() {
	dconf.SetConfigName("config")
	dconf.SetConfigPath(".")
	dconf.SetConfigType("ddas")
}

// Config program config
type Config struct {
	Log logx.Config `ddas:"log"`

	SrpcClient srpc.SrpcClientConf `ddas:"srpc_client"`
}

// NewConfig config
func NewConfig() (opts *Config, err error) {
	opts = new(Config)
	if err := dconf.Unmarshal(opts); err != nil {
		return nil, err
	}
	l, err := logx.GetLoggerByConf(&opts.Log)
	if err != nil {
		return nil, err
	}
	logx.X = l
	return
}

// GetConfig obtain pretty config
func GetConfig() (config string, err error) {
	opts := new(Config)
	var configBytes []byte
	err = dconf.Unmarshal(opts)
	if err != nil {
		fmt.Printf("[ERR] unmarshal error %v\n", err)
		return
	}
	configBytes, err = yaml.Marshal(opts)
	if err != nil {
		fmt.Printf("[ERR] yaml.Marshal error %v\n", err)
		return
	}
	config = string(configBytes)
	return
}
