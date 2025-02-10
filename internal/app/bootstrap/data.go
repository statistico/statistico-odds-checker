package bootstrap

import "github.com/statistico/statistico-football-data-go-grpc-client"

func (c Container) DataServiceResultClient() statisticofootballdata.FixtureClient {
	return statisticofootballdata.NewFixtureClient(c.GrpcFixtureClient())
}
