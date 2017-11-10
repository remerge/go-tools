package normalize

import "strings"

// CountryAlpha2 returns the ISO-3166-1 alpha 2 code for a given
// alpha 2, 3 or longer name. If not found "xx" is returned
func CountryAlpha2(country string) string {
	if country == "" {
		return "xx"
	}

	c := strings.ToLower(country)

	if len(c) == 2 {
		return c
	}

	if len(c) == 3 {
		if match, ok := countries[c]; ok {
			return match
		}
	}

	if match, ok := countriesNonStandard[c]; ok {
		return match
	}

	return "xx"
}

var countries = map[string]string{
	"and": "ad",
	"are": "ae",
	"afg": "af",
	"atg": "ag",
	"aia": "ai",
	"alb": "al",
	"arm": "am",
	"ago": "ao",
	"ata": "aq",
	"arg": "ar",
	"asm": "as",
	"aut": "at",
	"aus": "au",
	"abw": "aw",
	"ala": "ax",
	"aze": "az",
	"bih": "ba",
	"brb": "bb",
	"bgd": "bd",
	"bel": "be",
	"bfa": "bf",
	"bgr": "bg",
	"bhr": "bh",
	"bdi": "bi",
	"ben": "bj",
	"blm": "bl",
	"bmu": "bm",
	"brn": "bn",
	"bol": "bo",
	"bes": "bq",
	"bra": "br",
	"bhs": "bs",
	"btn": "bt",
	"bvt": "bv",
	"bwa": "bw",
	"blr": "by",
	"blz": "bz",
	"can": "ca",
	"cck": "cc",
	"cod": "cd",
	"caf": "cf",
	"cog": "cg",
	"che": "ch",
	"civ": "ci",
	"cok": "ck",
	"chl": "cl",
	"cmr": "cm",
	"chn": "cn",
	"col": "co",
	"cri": "cr",
	"cub": "cu",
	"cpv": "cv",
	"cuw": "cw",
	"cxr": "cx",
	"cyp": "cy",
	"cze": "cz",
	"deu": "de",
	"dji": "dj",
	"dnk": "dk",
	"dma": "dm",
	"dom": "do",
	"dza": "dz",
	"ecu": "ec",
	"est": "ee",
	"egy": "eg",
	"esh": "eh",
	"eri": "er",
	"esp": "es",
	"eth": "et",
	"fin": "fi",
	"fji": "fj",
	"flk": "fk",
	"fsm": "fm",
	"fro": "fo",
	"fra": "fr",
	"gab": "ga",
	"gbr": "gb",
	"grd": "gd",
	"geo": "ge",
	"guf": "gf",
	"ggy": "gg",
	"gha": "gh",
	"gib": "gi",
	"grl": "gl",
	"gmb": "gm",
	"gin": "gn",
	"glp": "gp",
	"gnq": "gq",
	"grc": "gr",
	"sgs": "gs",
	"gtm": "gt",
	"gum": "gu",
	"gnb": "gw",
	"guy": "gy",
	"hkg": "hk",
	"hmd": "hm",
	"hnd": "hn",
	"hrv": "hr",
	"hti": "ht",
	"hun": "hu",
	"idn": "id",
	"irl": "ie",
	"isr": "il",
	"imn": "im",
	"ind": "in",
	"iot": "io",
	"irq": "iq",
	"irn": "ir",
	"isl": "is",
	"ita": "it",
	"jey": "je",
	"jam": "jm",
	"jor": "jo",
	"jpn": "jp",
	"ken": "ke",
	"kgz": "kg",
	"khm": "kh",
	"kir": "ki",
	"com": "km",
	"kna": "kn",
	"prk": "kp",
	"kor": "kr",
	"xkx": "xk",
	"kwt": "kw",
	"cym": "ky",
	"kaz": "kz",
	"lao": "la",
	"lbn": "lb",
	"lca": "lc",
	"lie": "li",
	"lka": "lk",
	"lbr": "lr",
	"lso": "ls",
	"ltu": "lt",
	"lux": "lu",
	"lva": "lv",
	"lby": "ly",
	"mar": "ma",
	"mco": "mc",
	"mda": "md",
	"mne": "me",
	"maf": "mf",
	"mdg": "mg",
	"mhl": "mh",
	"mkd": "mk",
	"mli": "ml",
	"mmr": "mm",
	"mng": "mn",
	"mac": "mo",
	"mnp": "mp",
	"mtq": "mq",
	"mrt": "mr",
	"msr": "ms",
	"mlt": "mt",
	"mus": "mu",
	"mdv": "mv",
	"mwi": "mw",
	"mex": "mx",
	"mys": "my",
	"moz": "mz",
	"nam": "na",
	"ncl": "nc",
	"ner": "ne",
	"nfk": "nf",
	"nga": "ng",
	"nic": "ni",
	"nld": "nl",
	"nor": "no",
	"npl": "np",
	"nru": "nr",
	"niu": "nu",
	"nzl": "nz",
	"omn": "om",
	"pan": "pa",
	"per": "pe",
	"pyf": "pf",
	"png": "pg",
	"phl": "ph",
	"pak": "pk",
	"pol": "pl",
	"spm": "pm",
	"pcn": "pn",
	"pri": "pr",
	"pse": "ps",
	"prt": "pt",
	"plw": "pw",
	"pry": "py",
	"qat": "qa",
	"reu": "re",
	"rou": "ro",
	"srb": "rs",
	"rus": "ru",
	"rwa": "rw",
	"sau": "sa",
	"slb": "sb",
	"syc": "sc",
	"sdn": "sd",
	"ssd": "ss",
	"swe": "se",
	"sgp": "sg",
	"shn": "sh",
	"svn": "si",
	"sjm": "sj",
	"svk": "sk",
	"sle": "sl",
	"smr": "sm",
	"sen": "sn",
	"som": "so",
	"sur": "sr",
	"stp": "st",
	"slv": "sv",
	"sxm": "sx",
	"syr": "sy",
	"swz": "sz",
	"tca": "tc",
	"tcd": "td",
	"atf": "tf",
	"tgo": "tg",
	"tha": "th",
	"tjk": "tj",
	"tkl": "tk",
	"tls": "tl",
	"tkm": "tm",
	"tun": "tn",
	"ton": "to",
	"tur": "tr",
	"tto": "tt",
	"tuv": "tv",
	"twn": "tw",
	"tza": "tz",
	"ukr": "ua",
	"uga": "ug",
	"umi": "um",
	"usa": "us",
	"ury": "uy",
	"uzb": "uz",
	"vat": "va",
	"vct": "vc",
	"ven": "ve",
	"vgb": "vg",
	"vir": "vi",
	"vnm": "vn",
	"vut": "vu",
	"wlf": "wf",
	"wsm": "ws",
	"yem": "ye",
	"myt": "yt",
	"zaf": "za",
	"zmb": "zm",
	"zwe": "zw",
	"scg": "cs",
	"ant": "an",
}

// non-standardised names used by ea, some with misspellings
var countriesNonStandard = map[string]string{
	"afghanistan":                      "af",
	"aland islands":                    "ax",
	"albania":                          "al",
	"algeria":                          "dz",
	"american samoa":                   "as",
	"andorra":                          "ad",
	"angola":                           "ao",
	"anguilla":                         "ai",
	"antigua and barbuda":              "ag",
	"argentina":                        "ar",
	"armenia":                          "am",
	"aruba":                            "aw",
	"australia":                        "au",
	"austria":                          "at",
	"azerbaijan":                       "az",
	"bahamas":                          "bs",
	"bahrain":                          "bh",
	"bangladesh":                       "bd",
	"barbados":                         "bb",
	"belarus":                          "by",
	"belgium":                          "be",
	"belize":                           "bz",
	"bermuda":                          "bm",
	"bhutan":                           "bt",
	"bolivia":                          "bo",
	"bonaire":                          "bq",
	"bosnia":                           "ba",
	"botswana":                         "bw",
	"brazil":                           "br",
	"brunei darussalam":                "bn",
	"bulgaria":                         "bg",
	"cambodia":                         "kh",
	"cameroon":                         "cm",
	"canada":                           "ca",
	"cape verde":                       "cv",
	"cayman islands":                   "ky",
	"central african republic":         "cf",
	"chile":                            "cl",
	"china":                            "cn",
	"colombia":                         "co",
	"congo drc":                        "cd",
	"costa rica":                       "cr",
	"cote divoire":                     "ci",
	"croatia":                          "hr",
	"curaçao":                          "cw",
	"cyprus":                           "cy",
	"czech republic":                   "cz",
	"denmark":                          "dk",
	"djibouti":                         "dj",
	"dominica":                         "dm",
	"dominican republic":               "do",
	"ecuador":                          "ec",
	"egypt":                            "eg",
	"el salvador":                      "sv",
	"estonia":                          "ee",
	"ethiopia":                         "et",
	"faroe islands":                    "fo",
	"fiji":                             "fj",
	"finland":                          "fi",
	"france":                           "fr",
	"french guiana":                    "gf",
	"french polynesia":                 "pf",
	"gabon":                            "ga",
	"georgia":                          "ge",
	"germany":                          "de",
	"ghana":                            "gh",
	"gibraltar":                        "gi",
	"greece":                           "gr",
	"greenland":                        "gl",
	"grenada":                          "gd",
	"guadeloupe":                       "gp",
	"guam":                             "gu",
	"guatemala":                        "gt",
	"guernsey":                         "gg",
	"guinea":                           "gn",
	"guyana":                           "gy",
	"haiti":                            "ht",
	"honduras":                         "hn",
	"hong kong":                        "hk",
	"hungary":                          "hu",
	"iceland":                          "is",
	"india":                            "in",
	"indonesia":                        "id",
	"iran":                             "ir",
	"iraq":                             "iq",
	"ireland":                          "ie",
	"isle of man":                      "im",
	"israel":                           "il",
	"italy":                            "it",
	"jamaica":                          "jm",
	"japan":                            "jp",
	"jersey":                           "je",
	"jordan":                           "jo",
	"kazakhstan":                       "kz",
	"kenya":                            "ke",
	"korea":                            "kr",
	"kuwait":                           "kw",
	"laos":                             "la",
	"latvia":                           "lv",
	"lebanon":                          "lb",
	"liberia":                          "lr",
	"liechtenstein":                    "li",
	"lithuania":                        "lt",
	"luxembourg":                       "lu",
	"macao":                            "mo",
	"macedonia":                        "mk",
	"malawi":                           "mw",
	"malaysia":                         "my",
	"maldives":                         "mv",
	"mali":                             "ml",
	"malta":                            "mt",
	"marshall islands":                 "mh",
	"martinique":                       "mq",
	"mauritania":                       "mr",
	"mauritius":                        "mu",
	"mexico":                           "mx",
	"moldova":                          "md",
	"monaco":                           "mc",
	"mongolia":                         "mn",
	"morocco":                          "ma",
	"mozambique":                       "mz",
	"myanmar":                          "mm",
	"namibia":                          "na",
	"nepal":                            "np",
	"netherlands":                      "nl",
	"new caledonia":                    "nc",
	"new zealand":                      "nz",
	"nicaragua":                        "ni",
	"niger":                            "ne",
	"nigeria":                          "ng",
	"northern mariana islands":         "mp",
	"norway":                           "no",
	"oman":                             "om",
	"pakistan":                         "pk",
	"palestine":                        "ps",
	"panama":                           "pa",
	"papua new guinea":                 "pg",
	"paraguay":                         "py",
	"peru":                             "pe",
	"philippines":                      "ph",
	"poland":                           "pl",
	"portugal":                         "pt",
	"puerto rico":                      "pr",
	"qatar":                            "qa",
	"reunion":                          "re",
	"romania":                          "ro",
	"russia":                           "ru",
	"rwanda":                           "rw",
	"saint kitts and nevis":            "kn",
	"saint lucia":                      "lc",
	"saint martin (french part)":       "mf",
	"saint vincent and the grenadines": "vc",
	"san marino":                       "sm",
	"saudi arabia":                     "sa",
	"senegal":                          "sn",
	"serbia & montenegro":              "rs",
	"serbia":                           "rs",
	"seychelles":                       "sc",
	"sierre leone":                     "sl",
	"singapore":                        "sg",
	"sint maarten (dutch part)":        "sx",
	"slovakia":                         "sk",
	"slovenia":                         "si",
	"south africa":                     "za",
	"south sudan":                      "ss",
	"spain":                            "es",
	"sri lanka":                        "lk",
	"suriname":                         "sr",
	"sweden":                           "se",
	"switzerland":                      "ch",
	"syria":                            "sy",
	"taiwan":                           "tw",
	"tajikistan":                       "tj",
	"tanzania":                         "tz",
	"thailand":                         "th",
	"tonga":                            "to",
	"trinidad and tobago":              "tt",
	"tunisia":                          "tn",
	"turkey":                           "tr",
	"uganda":                           "ug",
	"ukraine":                          "ua",
	"united arab emirates":             "ae",
	"united kingdom":                   "gb",
	"united states":                    "us",
	"uruguay":                          "uy",
	"us virgin islands":                "vi",
	"uzbekistan":                       "uz",
	"vanuatu":                          "vu",
	"venezuela":                        "ve",
	"vietnam":                          "vn",
	"virgin islands, british":          "vg",
	"yemen":                            "ye",
	"zambia":                           "zm",
	"zimbabwe":                         "zw",
}
