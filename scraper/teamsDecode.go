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
		//La Liga
		"bar": "Barcelona",
		"ala": "Alaves",
		"rma": "Real Madrid",
		"val": "Valencia",
		"atb": "Athletic Bilbao",
		"gir": "Girona",
		"rso": "Real Sociedad",
		"bet": "Real Betis",
		"lpa": "Las Palmas",
		"rav": "Rayo Vallecano",
		"get": "Getafe",
		"osa": "Osasuna",
		"sev": "Sevilla",
		"vil": "Villareal",
		"cad": "Cadiz",
		"mal": "Mallorca",
		"cev": "Celta Vigo",
		"gra": "Granada",
		"alm": "Almeria",
	}

	homeTeam := teamNameMapping[game[0:3]]
	awayTeam := teamNameMapping[game[3:]]

	// Build similar string to the one of the website
	gameString := homeTeam + "-vs-" + awayTeam

	return gameString
}
