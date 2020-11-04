package bootstrap

import (
	sg "github.com/statistico/statistico-odds-checker/internal/grpc"
	"github.com/statistico/statistico-odds-checker/internal/grpc/proto"
	"google.golang.org/grpc"
)

func (c Container) GrpcFixtureClient() sg.FixtureClient {
	config := c.Config

	address := config.StatisticoDataService.Host + ":" + config.StatisticoDataService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		c.Logger.Warnf("Error initializing statistico data service grpc client %s", err.Error())
	}

	client := proto.NewFixtureServiceClient(conn)

	return sg.NewFixtureClient(client)
}

func (c Container) GrpcOddsCompilerClient() sg.OddsCompilerClient {
	config := c.Config

	address := config.StatisticoOddsCompilerService.Host + ":" + config.StatisticoOddsCompilerService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		c.Logger.Warnf("Error initializing statistico data service grpc client %s", err.Error())
	}

	client := proto.NewOddsCompilerServiceClient(conn)

	return sg.NewOddsCompilerClient(client)
}
