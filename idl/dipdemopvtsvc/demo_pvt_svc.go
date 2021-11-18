package dipdemopvtsvc

import (
	"context"

	pvt "go.danale.net/be/be/biz/demo/dartman-demo-client/idl"
	"go.danale.net/be/be/go/srpc"
)

// 服务发布方 ---------------------------------------------------------------------

const DIP_DEMO_PVT_SVC_NAME = "demo.DipDemoPvtSvc"

type DipDemoPvtSvc interface {
	SayHello(ctx context.Context, input *SayHelloInput, output *SayHelloOutput) error
}

type SayHelloInput struct {
	pvt.InputCommon
	Value string
}

type SayHelloOutput struct {
	pvt.OutputCommon
	Value string
}

// 服务使用方 ---------------------------------------------------------------------

type DipDemoPvtSvcClient interface {
	SayHello(ctx context.Context, input *SayHelloInput, opt ...srpc.CallOption) (*SayHelloOutput, error)
}

func NewDipDemoPvtSvcClient(srpcCli srpc.SrpcClient) DipDemoPvtSvcClient {
	return &dipDemoPvtSvcClientImpl{srpcCli: srpcCli}
}

type dipDemoPvtSvcClientImpl struct {
	srpcCli srpc.SrpcClient
}

func (d *dipDemoPvtSvcClientImpl) SayHello(ctx context.Context, input *SayHelloInput, opts ...srpc.CallOption) (*SayHelloOutput, error) {
	output := &SayHelloOutput{}
	err := d.srpcCli.CallSRPC(ctx, 0, DIP_DEMO_PVT_SVC_NAME, "SayHello", input, output, opts...)
	return output, err
}
