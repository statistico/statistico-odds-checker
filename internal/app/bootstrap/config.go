package bootstrap

import "os"

type Config struct {
	AwsConfig
	BetFair
	FootballConfig
	Publisher string
	Sentry
	StatisticoFootballDataService
}

type AwsConfig struct {
	Key      string
	Secret   string
	Region   string
	TopicArn string
}

type BetFair struct {
	Username string
	Password string
	Key      string
}

type FootballConfig struct {
	SupportedSeasons []uint64
	Markets          []string
}

type Sentry struct {
	DSN string
}

type StatisticoFootballDataService struct {
	Host string
	Port string
}

func BuildConfig() *Config {
	config := Config{}

	config.Publisher = os.Getenv("PUBLISHER")

	config.FootballConfig = FootballConfig{
		SupportedSeasons: []uint64{
			19734, // England - Premier League
			19793, // England - Championship
			19794, // England - League One
			19795, // England - League Two
			19917, // England - National League,
		},
		Markets: []string{
			"BOTH_TEAMS_TO_SCORE",
			"MATCH_ODDS",
			"MATCH_SHOTS",
			"MATCH_SHOTS_TARGET",
			"OVER_UNDER_05",
			"OVER_UNDER_15",
			"OVER_UNDER_25",
			"OVER_UNDER_35",
			"OVER_UNDER_45",
			"OVER_UNDER_25_CARDS",
			"OVER_UNDER_35_CARDS",
			"OVER_UNDER_45_CARDS",
			"OVER_UNDER_65_CARDS",
			"OVER_UNDER_55_CORNR",
			"OVER_UNDER_85_CORNR",
			"OVER_UNDER_105_CORNR",
			"OVER_UNDER_135_CORNR",
		},
	}

	config.AwsConfig = AwsConfig{
		Key:      os.Getenv("AWS_KEY"),
		Secret:   os.Getenv("AWS_SECRET"),
		Region:   os.Getenv("AWS_REGION"),
		TopicArn: os.Getenv("AWS_TOPIC_ARN"),
	}

	config.BetFair = BetFair{
		Username: os.Getenv("BETFAIR_USERNAME"),
		Password: os.Getenv("BETFAIR_PASSWORD"),
		Key:      os.Getenv("BETFAIR_KEY"),
	}

	config.Sentry = Sentry{DSN: os.Getenv("SENTRY_DSN")}

	config.StatisticoFootballDataService = StatisticoFootballDataService{
		Host: os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_HOST"),
		Port: os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_PORT"),
	}

	return &config
}
