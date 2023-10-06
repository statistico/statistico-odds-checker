package statistico

func marketIsSupported(market string) bool {
	markets := []string{OverUnder05, OverUnder15, OverUnder25, OverUnder35}

	for _, m := range markets {
		if m == market {
			return true
		}
	}

	return false
}
