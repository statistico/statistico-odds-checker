package main

import (
	"context"
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/bootstrap"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig())

	processor := app.Processor()
	clock := app.Clock

	console := &cli.App{
		Name: "Statistico Odds Checker - Command Line Application",
		Commands: []cli.Command{
			{
				Name:      "market:fetch",
				Usage:     "Fetch and parse markets for supported competitions",
				UsageText: "Fetch and parse markets for supported competitions",
				Before: func(c *cli.Context) error {
					fmt.Println("[INFO] Building and publishing markets")
					return nil
				},
				After: func(c *cli.Context) error {
					fmt.Println("[INFO] Complete")
					return nil
				},
				Action: func(c *cli.Context) error {
					from := clock.Now()
					to := clock.Now().Add(time.Hour * 12)

					ctx := context.Background()

					if err := processor.Process(ctx, from, to, c.String("exchange")); err != nil {
						return err
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "exchange",
						Usage:    "Find odds for upcoming events published by provided exchange",
						Required: true,
					},
				},
			},
		},
	}

	err := console.Run(os.Args)

	if err != nil {
		fmt.Printf("Error in executing command: %s\n", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
