package client

import (
	"go.danale.net/be/be/biz/demo/dartman-demo-client/internal/config"
	"go.danale.net/be/be/go/srpc"
)

type Client struct {
	cfg        *config.Config
	SrpcClient srpc.SrpcClient
}

func NewClient(cfg *config.Config) (*Client, error) {
	srpcClient, err := srpc.NewClientByConf(&cfg.SrpcClient)
	if err != nil {
		return nil, err
	}
	c := Client{
		cfg:        cfg,
		SrpcClient: srpcClient,
	}
	return &c, nil
}
