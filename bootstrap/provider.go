package bootstrap

import (
	"go.danale.net/be/be/biz/demo/dartman-demo-client/client"
	"go.danale.net/be/be/biz/demo/dartman-demo-client/internal/config"
	"go.danale.net/be/be/biz/demo/dartman-demo-client/provider"
	"go.danale.net/be/be/biz/demo/dartman-demo-client/service"
	ddaskit "go.danale.net/be/be/go/ddas-kit"
	"go.danale.net/be/be/go/logx"
)

func setupProvider() (*provider.Provider, error) {
	options, err := config.NewConfig()
	if err != nil {
		logx.X.Error(err)
		return nil, err
	}
	ddaskit.SetLogger(logx.X)
	logx.X.Info("start program")
	clientClient, err := client.NewClient(options)
	if err != nil {
		logx.X.Error(err)
		return nil, err
	}
	serviceService, err := service.NewService(options, clientClient)
	if err != nil {
		logx.X.Error(err)
		return nil, err
	}
	providerProvider, err := provider.NewProvider(options, serviceService, clientClient)
	if err != nil {
		logx.X.Error(err)
		return nil, err
	}
	return providerProvider, nil
}
