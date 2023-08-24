package betfair

var teams = map[string]string{
	"Las Palmas":               "Las Palmas",
	"Hearts":                   "Hearts",
	"Dundee":                   "Dundee",
	"Catanzaro":                "Catanzaro",
	"Modena":                   "Modena",
	"FeralpiSalò":              "Feralpisalo",
	"Reggiana":                 "Reggiana",
	"Venezia":                  "Venezia",
	"Frosinone":                "Frosinone",
	"Le Havre":                 "Le Havre",
	"Heidenheim":               "FC Heidenheim",
	"Darmstadt 98":             "SV Darmstadt",
	"RP Leipzig":               "RP Leipzig",
	"Almere City":              "Almere City",
	"NEC":                      "NEC Nijmegen",
	"Bromley":                  "Bromley",
	"Solihull Moors":           "Solihull Moors",
	"Eastleigh":                "Eastleigh",
	"Halifax Town":             "FC Halifax Town",
	"Aldershot Town":           "Aldershot",
	"Wealdstone":               "Wealdstone",
	"Altrincham":               "Altrincham",
	"Chesterfield":             "Chesterfield",
	"Boreham Wood":             "Boreham Wood",
	"Dagenham & Redbridge":     "Dag and Red",
	"Barnet":                   "Barnet",
	"Dorking Wanderers":        "Dorking Wanderers",
	"York City":                "York City",
	"Ebbsfleet United":         "Ebbsfleet Utd",
	"Kidderminster Harriers":   "Kidderminster",
	"Gateshead":                "Gateshead",
	"Woking":                   "Woking",
	"Maidenhead United":        "Maidenhead",
	"Oxford City":              "Oxford City",
	"Hartlepool United":        "Hartlepool",
	"Fylde":                    "AFC Fylde",
	"Ascoli":                   "Ascoli",
	"Bari 1908":                "SSD Bari",
	"Brescia":                  "Brescia",
	"Cittadella":               "Cittadella",
	"Cosenza":                  "Cosenza",
	"Palermo":                  "Palermo",
	"Perugia":                  "Perugia",
	"Pisa":                     "Pisa",
	"SPAL":                     "Spal",
	"Südtirol":                 "Sudtirol",
	"Como":                     "Como",
	"Ternana":                  "Ternana",
	"FC Augsburg":              "Augsburg",
	"VfL Bochum 1848":          "Bochum",
	"Borussia Dortmund":        "Dortmund",
	"Eintracht Frankfurt":      "Eintracht Frankfurt",
	"FC Bayern München":        "Bayern Munich",
	"FC Köln":                  "FC Koln",
	"Hertha BSC":               "Hertha Berlin",
	"TSG Hoffenheim":           "Hoffenheim",
	"Bayer 04 Leverkusen":      "Leverkusen",
	"Borussia Mönchengladbach": "Mgladbach",
	"RB Leipzig":               "RB Leipzig",
	"Schalke 04":               "Schalke 04",
	"VfB Stuttgart":            "Stuttgart",
	"FC Union Berlin":          "Union Berlin",
	"Werder Bremen":            "Werder Bremen",
	"VfL Wolfsburg":            "Wolfsburg",
	"FSV Mainz 05":             "Mainz",
	"West Ham United":          "West Ham",
	"Tottenham Hotspur":        "Tottenham",
	"Liverpool":                "Liverpool",
	"Manchester City":          "Man City",
	"West Bromwich Albion":     "West Brom",
	"Fulham":                   "Fulham",
	"Everton":                  "Everton",
	"Manchester United":        "Man Utd",
	"Aston Villa":              "Aston Villa",
	"Chelsea":                  "Chelsea",
	"Arsenal":                  "Arsenal",
	"Newcastle United":         "Newcastle",
	"Sheffield United":         "Sheff Utd",
	"Burnley":                  "Burnley",
	"Wolverhampton Wanders":    "Wolves",
	"Roma":                     "Roma",
	"Leicester City":           "Leicester",
	"Lazio":                    "Lazio",
	"Crystal Palace":           "Crystal Palace",
	"Southampton":              "Southampton",
	"Leeds United":             "Leeds",
	"Brighton & Hove Albion":   "Brighton",
	"Crotone":                  "Crotone",
	"Augsburg":                 "Augsburg",
	"Genoa":                    "Genoa",
	"Fiorentina":               "Fiorentina",
	"Milan":                    "AC Milan",
	"Spezia":                   "Spezia",
	"Udinese":                  "Udinese",
	"Parma":                    "Parma",
	"Bayern München":           "Bayern Munich",
	"Wolfsburg":                "Wolfsburg",
	"Sampdoria":                "Sampdoria",
	"Benevento":                "Benevento",
	"Cagliari":                 "Cagliari",
	"Napoli":                   "Napoli",
	"Torino":                   "Torino",
	"Juventus":                 "Juventus",
	"Salernitana":              "Salernitana",
	"Borussia M'gladbach":      "Mgladbach",
	"Atalanta":                 "Atalanta",
	"Mainz 05":                 "Mainz",
	"Union Berlin":             "Union Berlin",
	"Hellas Verona":            "Verona",
	"Sassuolo":                 "Sassuolo",
	"Empoli":                   "Empoli",
	"Cremonese":                "US Cremonese",
	"Hoffenheim":               "Hoffenheim",
	"Arminia Bielefeld":        "Arminia Bielefeld",
	"Inter":                    "Inter",
	"Stuttgart":                "Stuttgart",
	"Bayer Leverkusen":         "Leverkusen",
	"SC Freiburg":              "Freiburg",
	"Bologna":                  "Bologna",
	"Lens":                     "Lens",
	"Nantes":                   "Nantes",
	"Nîmes":                    "Nimes",
	"Reims":                    "Reims",
	"Paris Saint Germain":      "Paris St-G",
	"Saint-Étienne":            "St Etienne",
	"Montpellier":              "Montpellier",
	"Lille":                    "Lille",
	"Dijon":                    "Dijon",
	"Nice":                     "Nice",
	"Strasbourg":               "Strasbourg",
	"Clermont":                 "Clermont",
	"Toulouse":                 "Toulouse",
	"Troyes":                   "ESTAC Troyes",
	"Monaco":                   "Monaco",
	"Brest":                    "Brest",
	"Rennes":                   "Rennes",
	"Lorient":                  "Lorient",
	"Olympique Lyonnais":       "Lyon",
	"Bordeaux":                 "Bordeaux",
	"Olympique Marseille":      "Marseille",
	"Angers SCO":               "Angers",
	"Ajaccio":                  "AC Ajaccio",
	"Metz":                     "Metz",
	"Famalicão":                "Famalicao",
	"Santa Clara":              "Santa Clara",
	"Benfica":                  "Benfica",
	"Vitória SC":               "Guimaraes",
	"Moreirense":               "Moreirense",
	"Gil Vicente":              "Gil Vicente",
	"Tondela":                  "Tondela",
	"Sporting CP":              "Sporting Lisbon",
	"Nacional":                 "Nacional",
	"Belenenses":               "Belenenses",
	"Marítimo":                 "Maritimo",
	"Sporting Braga":           "Braga",
	"Paços de Ferreira":        "Ferreira",
	"Porto":                    "Porto",
	"Rio Ave":                  "Rio Ave",
	"Portimonense":             "Portimonense",
	"Boavista":                 "Boavista",
	"Farense":                  "Farense",
	"Elche":                    "Elche",
	"Almería":                  "Almeria",
	"Real Betis":               "Betis",
	"Villarreal":               "Villarreal",
	"Granada":                  "Granada",
	"Cádiz":                    "Cadiz",
	"Osasuna":                  "Osasuna",
	"Getafe":                   "Getafe",
	"Real Madrid":              "Real Madrid",
	"Levante":                  "Levante",
	"Real Sociedad":            "Real Sociedad",
	"Rayo Vallecano":           "Rayo Vallecano",
	"SD Eibar":                 "Eibar",
	"Huesca":                   "Huesca",
	"Athletic Club":            "Athletic Bilbao",
	"Barcelona":                "Barcelona",
	"Sevilla":                  "Sevilla",
	"Celta de Vigo":            "Celta Vigo",
	"FC Barcelona":             "Barcelona",
	"Valencia":                 "Valencia",
	"Atlético Madrid":          "Atletico Madrid",
	"Deportivo Alavés":         "Alaves",
	"Mallorca":                 "Mallorca",
	"Girona":                   "Girona",
	"Real Valladolid":          "Valladolid",
	"Sparta Rotterdam":         "Sparta Rotterdam",
	"Heracles":                 "Heracles",
	"Willem II":                "Willem II",
	"Vitesse":                  "Vitesse Arnhem",
	"FC Utrecht":               "FC Utrecht",
	"FC Groningen":             "FC Groningen",
	"Ajax":                     "Ajax",
	"PEC Zwolle":               "PEC Zwolle",
	"FC Emmen":                 "Emmen",
	"FC Volendam":              "FC Volendam",
	"Excelsior":                "Excelsior",
	"Feyenoord":                "Feyenoord",
	"FC Twente":                "FC Twente",
	"SC Cambuur":               "Cambuur Leeuwarden",
	"SC Heerenveen":            "Heerenveen",
	"AZ Alkmaar":               "Az Alkmaar",
	"VVV-Venlo":                "VVV Venlo",
	"PSV":                      "PSV",
	"ADO Den Haag":             "ADO Den Haag",
	"Fortuna Sittard":          "Fortuna Sittard",
	"RKC Waalwijk":             "RKC Waalwijk",
	"Varberg BoIS":             "Varbergs BolS",
	"Örebro":                   "Orebro",
	"Mjällby":                  "Mjallby",
	"Djurgården":               "Djurgarden",
	"AIK":                      "AIK",
	"IFK Göteborg":             "IFK Goteborg",
	"Sirius":                   "IK Sirius",
	"Östersunds FK":            "Ostersunds FK",
	"Helsingborg":              "Helsingborgs",
	"Falkenberg":               "Falkenbergs FF",
	"Kalmar":                   "Kalmar FF",
	"Malmö FF":                 "Malmo",
	"Häcken":                   "Hacken",
	"Elfsborg":                 "Elfsborg",
	"Norrköping":               "Norrkoping",
	"Hammarby":                 "Hammarby",
	"Grimsby Town":             "Grimsby",
	"Cambridge United":         "Cambridge Utd",
	"Tranmere Rovers":          "Tranmere",
	"Crawley Town":             "Crawley Town",
	"Salford City":             "Salford City",
	"Stevenage":                "Stevenage",
	"Newport County":           "Newport County",
	"Barrow":                   "Barrow",
	"Cheltenham Town":          "Cheltenham",
	"Carlisle United":          "Carlisle",
	"Morecambe":                "Morecambe",
	"Exeter City":              "Exeter",
	"Colchester United":        "Colchester",
	"Mansfield Town":           "Mansfield",
	"Southend United":          "Southend",
	"Walsall":                  "Walsall",
	"Harrogate Town":           "Harrogate Town",
	"Forest Green Rovers":      "Forest Green",
	"Leyton Orient":            "Leyton Orient",
	"Bolton Wanderers":         "Bolton",
	"Bradford City":            "Bradford",
	"Oldham Athletic":          "Oldham",
	"Port Vale":                "Port Vale",
	"Scunthorpe United":        "Scunthorpe",
	"Shrewsbury Town":          "Shrewsbury",
	"Oxford United":            "Oxford Utd",
	"Charlton Athletic":        "Charlton",
	"Fleetwood Town":           "Fleetwood Town",
	"Blackpool":                "Blackpool",
	"Peterborough United":      "Peterborough",
	"Crewe Alexandra":          "Crewe",
	"Burton Albion":            "Burton",
	"Wigan Athletic":           "Wigan",
	"Accrington Stanley":       "Accrington",
	"AFC Wimbledon":            "AFC Wimbledon",
	"Hull City":                "Hull",
	"Bristol Rovers":           "Bristol Rovers",
	"Ipswich Town":             "Ipswich",
	"Sunderland":               "Sunderland",
	"Swindon Town":             "Swindon",
	"Portsmouth":               "Portsmouth",
	"Rochdale":                 "Rochdale",
	"Northampton Town":         "Northampton",
	"Gillingham":               "Gillingham",
	"Doncaster Rovers":         "Doncaster",
	"Plymouth Argyle":          "Plymouth",
	"Lincoln City":             "Lincoln",
	"Milton Keynes Dons":       "MK Dons",
	"Luton Town":               "Luton",
	"Derby County":             "Derby",
	"Reading":                  "Reading",
	"AFC Bournemouth":          "Bournemouth",
	"Blackburn Rovers":         "Blackburn",
	"Barnsley":                 "Barnsley",
	"Middlesbrough":            "Middlesbrough",
	"Brentford":                "Brentford",
	"Swansea City":             "Swansea",
	"Stoke City":               "Stoke",
	"Millwall":                 "Millwall",
	"Watford":                  "Watford",
	"Preston North End":        "Preston",
	"Wycombe Wanderers":        "Wycombe",
	"Norwich City":             "Norwich",
	"Bristol City":             "Bristol City",
	"Birmingham City":          "Birmingham",
	"Coventry City":            "Coventry",
	"Huddersfield Town":        "Huddersfield",
	"Nottingham Forest":        "Nottm Forest",
	"Queens Park Rangers":      "QPR",
	"Rotherham United":         "Rotherham",
	"Cardiff City":             "Cardiff",
	"Sheffield Wednesday":      "Sheff Wed",
	"OH Leuven":                "Oud-Heverlee Leuven",
	"Kortrijk":                 "Kortrijk",
	"Sporting Charleroi":       "Charleroi",
	"Club Brugge":              "Club Brugge",
	"Antwerp":                  "Antwerp",
	"Standard Liège":           "Standard",
	"KV Oostende":              "KV Oostende",
	"AS Eupen":                 "Eupen",
	"Waasland-Beveren":         "Waasland-Beveren",
	"Zulte-Waregem":            "Zulte-Waregem",
	"Sint-Truiden":             "Sint Truiden",
	"Anderlecht":               "Anderlecht",
	"Gent":                     "Gent",
	"Genk":                     "Genk",
	"Cercle Brugge":            "Cercle Brugge",
	"Mechelen":                 "Yellow-Red Mechelen",
	"Beerschot-Wilrijk":        "Kfco Beerschot Wilrijk",
	"Royal Excel Mouscron":     "Royal Mouscron-peruwelz",
	"Kilmarnock":               "Kilmarnock",
	"Ross County":              "Ross Co",
	"Hibernian":                "Hibernian",
	"Celtic":                   "Celtic",
	"St. Mirren":               "St Mirren",
	"Dundee United":            "Dundee Utd",
	"St. Johnstone":            "St Johnstone",
	"Rangers":                  "Rangers",
	"Livingston":               "Livingston",
	"Aberdeen":                 "Aberdeen",
	"Motherwell":               "Motherwell",
	"Hamilton Academical":      "Hamilton",
	"Lecce":                    "Lecce",
	"Monza":                    "AC Monza",
	"Adana Demirspor":          "Adana Demirspor",
	"Alanyaspor":               "Alanyaspor",
	"Ankaragücü":               "Ankaragucu",
	"Antalyaspor":              "Antalyaspor",
	"Beşiktaş":                 "Besiktas",
	"Fatih Karagümrük":         "Fatih Karagumruk Istanbul",
	"Fenerbahçe":               "Fenerbahce",
	"Galatasaray":              "Galatasaray",
	"Gaziantep F.K.":           "Gaziantep FK",
	"Giresunspor":              "Giresunspor",
	"Hatayspor":                "Hatayspor",
	"İstanbulspor":             "Istanbulspor",
	"Kasımpaşa":                "Kasimpasa",
	"Kayserispor":              "Kayserispor",
	"Konyaspor":                "Konyaspor",
	"Pendikspor":               "Pendikspor",
	"Rizespor":                 "Rizespor",
	"Samsunspor":               "Samsunspor",
	"Sivasspor":                "Sivasspor",
	"Trabzonspor":              "Trabzonspor",
	"Ümraniyespor":             "Umraniyespor",
}
