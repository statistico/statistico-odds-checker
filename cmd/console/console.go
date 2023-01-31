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
					from, err := time.Parse(time.RFC3339, c.String("from"))

					if err != nil {
						return err
					}

					to, err := time.Parse(time.RFC3339, c.String("to"))

					if err != nil {
						return err
					}

					ctx := context.Background()

					if err = processor.Process(ctx, from, to); err != nil {
						return err
					}

					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "from",
						Usage:    "Find events with start date/time after value provided",
						Value:    clock.Now().Format(time.RFC3339),
						Required: false,
					},
					&cli.StringFlag{
						Name:     "to",
						Usage:    "Find events with start date/time before value provided",
						Value:    clock.Now().Add(time.Hour * 12).Format(time.RFC3339),
						Required: false,
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
