package utils

var RubAliases = []string{
	"rub", "rur", "rub.", "ruble", "rubles", "rouble", "roubles",
	"руб", "руб.", "рубль", "рубля", "рублей",
	"р", "р.", "₽",
	"ру",
	"py", "p", "p.",
}

var UsdAliases = []string{
	"usd", "us$", "us$.", "us dollars", "us dollar", "us-dollars",
	"unitedstatesdollar", "united states dollar",
	"$", "＄",
	"usd$", "$usd",
	"dollar", "dollars", "us dollar", "us dollars",
	"американскийдоллар", "американские-доллары",
	"доллар", "доллары", "долл", "долл.", "долларов",
	"u$s", "u$s.", "us d", "usdl",
}

var AedAliases = []string{
	"aed",
	"uae dirham", "uaedirham",
	"united arab emirates dirham", "unitedarabemiratesdirham",
	"emirati dirham", "emiratidirham",
	"dirham", "dirhams",
	"dh", "dhs", "dhm", "dhms",
	"درهم", "د.إ", "دإ", "دإ.", "د.إ", "د. إ",
	"aed دإ", "دإ aed",
	"дирхам", "дирхамы", "дирхамов",
	"дирхамоаэ", "дирхам оаэ",
	"дирхам-оаэ", "оаэдирхам", "оаэ дирхам",
}

var TryAliases = []string{
	"try", "trl",
	"₺",
	"turkish lira", "turkishlira",
	"turkey lira", "turkeylira",
	"tl",
	"try₺", "₺try",
	"турецкаялира", "турецкая лира",
	"лира", "лиры", "лир",
	"турецкиелиры", "турецкие лиры",
	"турецкихлир", "турецких лир",
	"турлира", "тур лира",
	"tl.", "tl₺", "₺tl",
}

var EurAliases = []string{
	"eur", "euro", "euros",
	"€", "eur€", "€eur",
	"eu", "eur.", "euro.", "€.",
	"евро",
	"eurо",
	"евр",
	"евров",
	"european union euro", "europeanunionuro",
	"европейская валюта", "европейский евро",
}
