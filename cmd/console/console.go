package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/jonboulle/clockwork"
	"github.com/statistico/statistico-odds-checker/internal/app/bootstrap"
	"os"
	"time"
)

var sport = flag.String("sport", "", "The sport to check odds for")
var dateFrom = flag.String("dateFrom", "", "Date range to begin from")
var dateTo = flag.String("dateTo", "", "Date range to end at")

func main() {
	app := bootstrap.BuildContainer(bootstrap.BuildConfig())
	processor := app.Processor()
	clock := app.Clock

	flag.Parse()

	ctx := context.Background()
	from := parseDateFrom(clock)
	to := parseDateTo(clock)

	fmt.Println("[INFO] Building and publishing markets...")

	err := processor.Process(ctx, *sport, from, to)

	if err != nil {
		fmt.Printf("Error parsing markets: %s", err.Error())
		os.Exit(1)
	}

	fmt.Println("[INFO] Complete")
	os.Exit(0)
}

func parseDateFrom(clock clockwork.Clock) time.Time {
	if *dateFrom == "" {
		return clock.Now()
	}

	from, err := time.Parse(time.RFC3339, *dateFrom)

	if err != nil {
		fmt.Printf("Error parsing date from %q", err)
		os.Exit(1)
	}

	return from
}

func parseDateTo(clock clockwork.Clock) time.Time {
	if *dateTo == "" {
		return clock.Now().Add(time.Hour * 12)
	}

	to, err := time.Parse(time.RFC3339, *dateTo)

	if err != nil {
		fmt.Printf("Error parsing date to %q", err)
		os.Exit(1)
	}

	return to
}
