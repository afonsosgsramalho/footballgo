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
		"atm": "Atletico Madrid",
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
		//Serie A
		"juv": "Juventus",
		"int": "Inter",
		"mil": "Milan",
		"nap": "Napoli",
		"ata": "Atalanta",
		"fio": "Fiorentina",
		"rom": "Roma",
		"bol": "Bologna",
		"mon": "Monza",
		"laz": "Lazio",
		"tor": "Torino",
		"fro": "Frosinone",
		"gen": "Genoa",
		"lec": "Lecce",
		"sas": "Sassuolo",
		"udi": "Udinese",
		"emp": "Empoli",
		"cag": "Cagliari",
		"ver": "Verona",
		"sal": "Salernitana",
		//Bundesliga
		"lev": "Leverkusen",
		"bay": "Bayern Munich",
		"stu": "Stuttgart",
		"lei": "Leipzig",
		"bvb": "Borussia Dortmund",
		"hof": "Hoffenheim",
		"fra": "Frankfurt",
		"fre": "Freiburg",
		"mgl": "Mongladbach",
		"aug": "Augsburg",
		"wof": "Wolfsburg",
		"wbr": "Bremen",
		"hei": "Heidenheim",
		"boc": "Bochum",
		"dar": "Darmstadt",
		"mai": "Mainz",
		"col": "FC Cologne",
		"uni": "Union Berlin",
		//Primeira liga
		"spo": "Sporting",
		"ben": "Benfica",
		"por": "Porto",
		"bra": "Braga",
		//Ligue One
		"par": "PSG",
		"nic": "Nice",
		"mno": "Monaco",
		"lil": "Lille",
		"rei": "Reims",
		"len": "RC Lens",
		"leh": "Le Havre",
		"brt": "Brest",
		"nan": "Nantes",
		"mar": "Marseille",
		"met": "Metz",
		"mot": "Montpellier",
		"ren": "Rennes",
		"tou": "Toulouse",
		"str": "Strasbourg",
		"lor": "Lorient",
		"cle": "Clermont",
		"lyo": "Lyon",
	}

	homeTeam := teamNameMapping[game[0:3]]
	awayTeam := teamNameMapping[game[3:]]

	// Build similar string to the one of the website
	gameString := homeTeam + "-vs-" + awayTeam

	return gameString
}
