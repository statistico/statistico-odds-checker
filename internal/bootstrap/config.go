package bootstrap

type Config struct {
	FootballConfig
}

type FootballConfig struct {
	SupportedSeasons []uint64
	Markets          []string
}

func BuildConfig() *Config {
	config := Config{}

	config.FootballConfig = FootballConfig{
		SupportedSeasons: []uint64{
			17361,
			17420,
			17488,
		},
		Markets:          []string{
			"OVER_UNDER_25",
		},
	}

	return &config
}
