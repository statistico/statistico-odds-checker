package bootstrap

import "os"

type Config struct {
	AwsConfig
	BetFair
	FootballConfig
	Publisher string
	Sentry
	SportsMonks
	StatisticoFootballDataService
	StatisticoOddsCompilerService
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
	Markets []string
}

type Sentry struct {
	DSN string
}

type SportsMonks struct {
	ApiKey string
}

type StatisticoFootballDataService struct {
	Host string
	Port string
}

type StatisticoOddsCompilerService struct {
	Host string
	Port string
}

func BuildConfig() *Config {
	config := Config{}

	config.Publisher = os.Getenv("PUBLISHER")

	config.FootballConfig = FootballConfig{
		// Betfair market terminology is used as our blueprint and standard. Internally we need to parse and handle
		// other supported exchanges markets using Betfair as a base i.e. OVER_UNDER_25 refers to Over/Under 2.5 goals.
		// Some markets may refer to this market a different way, so this needs to be handled that internally.
		Markets: []string{
			"BOTH_TEAMS_TO_SCORE",
			"MATCH_ODDS",
			"OVER_UNDER_05",
			"OVER_UNDER_15",
			"OVER_UNDER_25",
			"OVER_UNDER_35",
			"OVER_UNDER_45",
			"OVER_UNDER_55_CORNR",
			"OVER_UNDER_85_CORNR",
			"OVER_UNDER_95_CORNR",
			"OVER_UNDER_105_CORNR",
			"OVER_UNDER_115_CORNR",
			"OVER_UNDER_125_CORNR",
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

	config.SportsMonks = SportsMonks{
		ApiKey: os.Getenv("SPORTMONKS_API_KEY"),
	}

	config.StatisticoFootballDataService = StatisticoFootballDataService{
		Host: os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_HOST"),
		Port: os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_PORT"),
	}

	config.StatisticoOddsCompilerService = StatisticoOddsCompilerService{
		Host: os.Getenv("STATISTICO_ODDS_COMPILER_SERVICE_HOST"),
		Port: os.Getenv("STATISTICO_ODDS_COMPILER_SERVICE_PORT"),
	}

	return &config
}
