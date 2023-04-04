package bootstrap

import "github.com/statistico/statistico-odds-compiler-go-grpc-client"

func (c Container) OddsCompilerClient() statisticooddscompiler.OddCompilerClient {
	return statisticooddscompiler.NewOddsCompilerClient(c.GrpcOddsCompilerClient())
}
