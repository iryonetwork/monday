package monday

// Locale identifies locales supported by 'monday' package.
// Monday uses ICU locale identifiers. See http://userguide.icu-project.org/locale
type Locale string

const (
	LocaleEnUS = "en_US" // English (United States)
	LocaleEnGB = "en_GB" // English (United Kingdom)
	LocaleDaDK = "da_DK" // Danish (Denmark)
	LocaleNlBE = "nl_BE" // Dutch (Belgium)
	LocaleNlNL = "nl_NL" // Dutch (Netherlands)
	LocaleFiFI = "fi_FI" // Finnish (Finland)
	LocaleFrFR = "fr_FR" // French (France)
	LocaleFrCA = "fr_CA" // French (Canada)
	LocaleDeDE = "de_DE" // German (Germany)
	LocaleHuHU = "hu_HU" // Hungarian (Hungary)
	LocaleItIT = "it_IT" // Italian (Italy)
	LocaleNnNO = "nn_NO" // Norwegian Nynorsk (Norway)
	LocaleNbNO = "nb_NO" // Norwegian Bokmål (Norway)
	LocalePlPL = "pl_PL" // Polish (Poland)
	LocalePtPT = "pt_PT" // Portuguese (Portugal)
	LocalePtBR = "pt_BR" // Portuguese (Brazil)
	LocaleRoRO = "ro_RO" // Romanian (Romania)
	LocaleRuRU = "ru_RU" // Russian (Russia)
	LocaleEsES = "es_ES" // Spanish (Spain)
	LocaleCaES = "ca_ES" // Catalan (Spain)
	LocaleSvSE = "sv_SE" // Swedish (Sweden)
	LocaleTrTR = "tr_TR" // Turkish (Turkey)
	LocaleBgBG = "bg_BG" // Bulgarian (Bulgaria)
	LocaleZhCN = "zh_CN" // Chinese (Mainland)
	LocaleZhTW = "zh_TW" // Chinese (Taiwan)
	LocaleZhHK = "zh_HK" // Chinese (Hong Kong)
	LocaleJaJP = "ja_JP" // Japanese (Japan)
	LocaleElGR = "el_GR" // Greek (Greece)
	LocaleIdID = "id_ID" // Indonesian (Indonesia)
	LocaleFrGP = "fr_GP" // Guadeloupe
)

// ListLocales returns all locales supported by the package.
func ListLocales() []Locale {
	return []Locale{
		LocaleEnUS,
		LocaleEnGB,
		LocaleDaDK,
		LocaleNlBE,
		LocaleNlNL,
		LocaleFiFI,
		LocaleFrFR,
		LocaleFrCA,
		LocaleDeDE,
		LocaleHuHU,
		LocaleItIT,
		LocaleNnNO,
		LocaleNbNO,
		LocalePlPL,
		LocalePtPT,
		LocalePtBR,
		LocaleRoRO,
		LocaleRuRU,
		LocaleEsES,
		LocaleCaES,
		LocaleSvSE,
		LocaleTrTR,
		LocaleBgBG,
		LocaleZhCN,
		LocaleZhTW,
		LocaleZhHK,
		LocaleJaJP,
		LocaleElGR,
		LocaleFrGP,
	}
}
