package main

import (
	"context"
	"flag"
	"fmt"
	"runtime/debug"

	"go.danale.net/be/be/biz/demo/dartman-demo-client/client"
	demo "go.danale.net/be/be/biz/demo/dartman-demo-client/idl/dipdemopvtsvc"
	"go.danale.net/be/be/biz/demo/dartman-demo-client/internal/config"
	_ "go.danale.net/be/be/go/ddas-kit"
	"go.danale.net/be/be/go/logx"
	"go.danale.net/be/be/go/version"
)

func main() {
	c := flag.Bool("c", false, "show config")
	v := flag.Bool("v", false, "show version")
	flag.Parse()
	if *c {
		cfg, err := config.GetConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(cfg)
		return
	}
	if *v {
		fmt.Println(version.FullVersion())
		return
	}

	defer func() {
		if r := recover(); r != nil {
			logx.X.With("err", r, "stack", string(debug.Stack())).Error("panic")
		}
		logx.X.Flush()
	}()

	options, err := config.NewConfig()
	if err != nil {
		logx.X.Error(err)
		return
	}
	clientClient, err := client.NewClient(options)
	if err != nil {
		logx.X.Error(err)
		return
	}
	err = CallSayHello(clientClient, options, "demo.DipDemoPvtSvc")

	return
}

func CallSayHello(cli *client.Client, options *config.Config, serviceName string) (err error) {
	in := demo.SayHelloInput{
		Value: "hello demo",
	}
	out := demo.SayHelloOutput{}

	err = cli.SrpcClient.CallSRPC(context.TODO(), options.SrpcClient.NodeID, serviceName, "SayHello", &in, &out)
	if err != nil {
		err = fmt.Errorf("CallSRPC SayHello failed,err is %v,service is %v,args is %v", err, serviceName, in)
		return
	}
	logx.X.Debugf("service %v SayHello is %v", serviceName, out.Value)
	return
}
