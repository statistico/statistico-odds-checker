package statistico

func marketIsSupported(market string) bool {
	markets := []string{OverUnder25}

	for _, m := range markets {
		if m == market {
			return true
		}
	}

	return false
}
