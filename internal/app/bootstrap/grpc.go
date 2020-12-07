package bootstrap

import (
	sg "github.com/statistico/statistico-odds-checker/internal/app/grpc"
	statisticoproto "github.com/statistico/statistico-proto/statistico-data/go"
	"google.golang.org/grpc"
)

func (c Container) GrpcFixtureClient() sg.FixtureClient {
	config := c.Config

	address := config.StatisticoDataService.Host + ":" + config.StatisticoDataService.Port

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		c.Logger.Warnf("Error initializing statistico data service grpc client %s", err.Error())
	}

	client := statisticoproto.NewFixtureServiceClient(conn)

	return sg.NewFixtureClient(client)
}
