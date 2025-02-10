package bootstrap

import "os"

type Config struct {
	AwsConfig
	BetFair
	FootballConfig
	Publisher string
	RedisConfig
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

type RedisConfig struct {
	Host     string
	Port     string
	Database string
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
		Markets: []string{
			"BOTH_TEAMS_TO_SCORE",
			"MATCH_ODDS",
		},
	}

	config.AwsConfig = AwsConfig{
		Region:   os.Getenv("AWS_REGION"),
		TopicArn: getSsmParameter("statistico-odds-checker-AWS_TOPIC_ARN"),
	}

	config.BetFair = BetFair{
		Username: getSsmParameter("statistico-odds-checker-BETFAIR_USERNAME"),
		Password: getSsmParameter("statistico-odds-checker-BETFAIR_PASSWORD"),
		Key:      getSsmParameter("statistico-odds-checker-BETFAIR_KEY"),
	}

	config.RedisConfig = RedisConfig{
		Host:     getSsmParameter("statistico-odds-checker-REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Database: os.Getenv("REDIS_DATABASE"),
	}

	config.Sentry = Sentry{DSN: getSsmParameter("statistico-odds-checker-SENTRY_DSN")}

	config.SportsMonks = SportsMonks{
		ApiKey: getSsmParameter("statistico-odds-checker-SPORTMONKS_API_KEY"),
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
