package bootstrap

import "github.com/statistico/statistico-data-go-grpc-client"

func (c Container) DataServiceResultClient() statisticodata.FixtureClient {
	return statisticodata.NewFixtureClient(c.GrpcFixtureClient())
}
