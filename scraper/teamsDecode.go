package scraper

func _parseGame(game string) string {
	teamNameMapping := map[string]string{
		//Premier League
		"ars": "Arsenal",
		"bur": "Burnley",
		"bha": "Brighton",
		"cry": "Crystal Palace",
		"mun": "Manchester United",
		"avl": "Aston Villa",
		"eve": "Everton",
		"bou": "Bournemouth",
		"liv": "Liverpool",
		"che": "Chelsea",
		"mci": "Manchester City",
		"ful": "Fulham",
		"bre": "Brentford",
		"lut": "Luton",
		"new": "Newcastle",
		"nfo": "Nottingham",
		"tot": "Tottenham",
		"whu": "West Ham",
		"wol": "Wolves",
		"shu": "Sheffield United",
	}

	homeTeam := teamNameMapping[game[0:3]]
	awayTeam := teamNameMapping[game[3:]]

	// Build similar string to the one of the website
	gameString := homeTeam + "-vs-" + awayTeam

	return gameString
}
