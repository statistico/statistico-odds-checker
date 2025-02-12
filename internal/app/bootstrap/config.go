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
			"MATCH_CORNERS",
			"MATCH_GOALS",
			"MATCH_ODDS",
			"PLAYER_SHOTS_ON_TARGET",
			"PLAYER_TACKLES",
			"PLAYER_TO_SCORE_ANYTIME",
			"PLAYER_TOTAL_SHOTS",
			"TEAM_CARDS",
			"TEAM_CORNERS",
			"TEAM_SHOTS",
			"TEAM_SHOTS_ON_TARGET",
		},
	}

	config.AwsConfig = AwsConfig{
		Region:   os.Getenv("AWS_REGION"),
		TopicArn: os.Getenv("AWS_TOPIC_ARN"),
	}

	config.BetFair = BetFair{
		Username: os.Getenv("BETFAIR_USERNAME"),
		Password: os.Getenv("BETFAIR_PASSWORD"),
		Key:      os.Getenv("BETFAIR_KEY"),
	}

	config.RedisConfig = RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Port:     os.Getenv("REDIS_PORT"),
		Database: os.Getenv("REDIS_DATABASE"),
	}

	config.Sentry = Sentry{DSN: os.Getenv("SENTRY_DSN")}

	config.SportsMonks = SportsMonks{
		ApiKey: os.Getenv("SPORTMONKS_API_KEY"),
	}

	config.StatisticoFootballDataService = StatisticoFootballDataService{
		Host: os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_HOST"),
		Port: os.Getenv("STATISTICO_FOOTBALL_DATA_SERVICE_PORT"),
	}

	return &config
}
