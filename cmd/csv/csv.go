package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/statistico/statistico-odds-checker/internal/app/bootstrap"
	"github.com/statistico/statistico-odds-checker/internal/app/exchange"
	"github.com/statistico/statistico-odds-checker/internal/app/sport"
	"os"
	"strconv"
	"time"
)

const (
	MatchOddsPrefix = "1"
	OverUnder25Prefix = "2"
)

var path = flag.String("path", "", "The filepath of the csv file")
var market = flag.String("market", "", "The market to publish")

func main() {
	flag.Parse()

	app := bootstrap.BuildContainer(bootstrap.BuildConfig())
	logger := app.Logger
	publisher := app.Publisher()

	lines, err := readCsv(*path)

	if err != nil {
		logger.Warnf("Error reading csv: %s", err.Error())
		os.Exit(1)
	}

	for _, line := range lines {
		var m *sport.EventMarket

		switch *market {
		case "MATCH_ODDS":
			m = buildMatchOdds(line)
		case "OVER_UNDER_25":
			m = buildOverUnder(line)
		default:
			fmt.Println("Market provided is not supported")
			os.Exit(1)
		}

		if err := publisher.PublishMarket(m); err != nil {
			logger.Warnf("Error publishing market %s: %s", m.ID, err.Error())
			continue
		}

		fmt.Printf("Processed market %s\n", m.ID)
	}

	fmt.Println("Processing complete. Goodbye")
	os.Exit(0)
}

func readCsv(filename string) ([][]string, error) {
	f, err := os.Open(filename)

	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()

	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func buildEventMarket(row []string) *sport.EventMarket {
	eventID, err := strconv.Atoi(row[0])

	if err != nil {
		panic(err)
	}

	seasonID, err := strconv.Atoi(row[1])

	if err != nil {
		panic(err)
	}

	date, err := strconv.Atoi(row[2])

	if err != nil {
		panic(err)
	}

	compID, err := strconv.Atoi(row[6])

	if err != nil {
		panic(err)
	}

	return &sport.EventMarket{
		EventID:       uint64(eventID),
		CompetitionID: uint64(compID),
		SeasonID:      uint64(seasonID),
		Sport:         "football",
		EventDate:     time.Unix(int64(date), 0).Format(time.RFC3339),
		Side:          "BACK",
		Exchange:      "n/a",
		Timestamp:     int64(date),
	}
}

func buildMatchOdds(row []string) *sport.EventMarket {
 	m := buildEventMarket(row)

 	home, err := strconv.ParseFloat(row[7], 32)

	if err != nil {
		panic(err)
	}

	away, err := strconv.ParseFloat(row[9], 32)

	if err != nil {
		panic(err)
	}

	draw, err := strconv.ParseFloat(row[8], 32)

	if err != nil {
		panic(err)
	}

	m.ID = fmt.Sprintf("%s-%s", MatchOddsPrefix, row[0])
	m.MarketName = "MATCH_ODDS"
 	m.Runners = []*exchange.Runner{
 		{
 			ID: 0,
 			Name: "Home",
 			Sort: 1,
 			Prices: []exchange.PriceSize{
 				{
 					Price: float32(home),
				},
			},
		},
		{
			ID: 0,
			Name: "Away",
			Sort: 2,
			Prices: []exchange.PriceSize{
				{
					Price: float32(away),
				},
			},
		},
		{
			ID: 0,
			Name: "Draw",
			Sort: 3,
			Prices: []exchange.PriceSize{
				{
					Price: float32(draw),
				},
			},
		},
	}

	return m
}

func buildOverUnder(row []string) *sport.EventMarket {
	m := buildEventMarket(row)

	over, err := strconv.ParseFloat(row[10], 32)

	if err != nil {
		panic(err)
	}

	under, err := strconv.ParseFloat(row[11], 32)

	if err != nil {
		panic(err)
	}

	m.ID = fmt.Sprintf("%s-%s", OverUnder25Prefix, row[0])
	m.MarketName = "OVER_UNDER_25"
	m.Runners = []*exchange.Runner{
		{
			ID: 0,
			Name: "Over 2.5 Goals",
			Sort: 1,
			Prices: []exchange.PriceSize{
				{
					Price: float32(over),
				},
			},
		},
		{
			ID: 0,
			Name: "Under 2.5 Goals",
			Sort: 2,
			Prices: []exchange.PriceSize{
				{
					Price: float32(under),
				},
			},
		},
	}

	return m
}
