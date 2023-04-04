package bootstrap

import (
	"github.com/statistico/statistico-proto/go"
	"google.golang.org/grpc"
)

func (c Container) GrpcFixtureClient() statistico.FixtureServiceClient {
	config := c.Config

	address := config.StatisticoFootballDataService.Host + ":" + config.StatisticoFootballDataService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		c.Logger.Warnf("Error initializing statistico data service grpc client %s", err.Error())
	}

	return statistico.NewFixtureServiceClient(conn)
}

func (c Container) GrpcOddsCompilerClient() statistico.OddsCompilerServiceClient {
	config := c.Config

	address := config.StatisticoOddsCompilerService.Host + ":" + config.StatisticoOddsCompilerService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		c.Logger.Warnf("Error initializing statistico odd compiler service grpc client %s", err.Error())
	}

	return statistico.NewOddsCompilerServiceClient(conn)
}
