package bootstrap

import (
	"github.com/statistico/statistico-odds-checker/internal/app/sportmonks"
	spClient "github.com/statistico/statistico-sportmonks-go-client"
	"net"
	"net/http"
	"time"
)

func (c Container) SportmonksOddsParser() sportmonks.OddsParser {
	s := c.Config.SportsMonks

	cl := spClient.NewDefaultHTTPClient(s.ApiKey)

	trans := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		ResponseHeaderTimeout: 120 * time.Second,
		TLSHandshakeTimeout:   15 * time.Second,
	}

	client := &http.Client{
		Timeout:   time.Second * 60,
		Transport: trans,
	}

	cl.SetHTTPClient(client)

	return sportmonks.NewOddsParser(cl)
}
