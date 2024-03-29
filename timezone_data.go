// Copyright 2019 Atharva Pandey. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gountry

//Timezones contains data mapping for timezone country name and country capital and continental region
var Timezones = []TimezoneType{
	{
		Timezones: []string{"Europe/Andorra"},
		ISO2:      "AD",
		Continent: "Europe",
		Name:      "Andorra",
		Capital:   "Andorra la Vella",
	},
	{
		Timezones: []string{"Asia/Kabul"},
		ISO2:      "AF",
		Continent: "Asia",
		Name:      "Afghanistan",
		Capital:   "Kabul"},
	{
		Timezones: []string{"America/Antigua"},
		ISO2:      "AG",
		Continent: "North America",
		Name:      "Antigua and Barbuda",
		Capital:   "St. John's"},
	{
		Timezones: []string{"Europe/Tirane"},
		ISO2:      "AL",
		Continent: "Europe",
		Name:      "Albania",
		Capital:   "Tirana"},
	{
		Timezones: []string{"Asia/Yerevan"},
		ISO2:      "AM",
		Continent: "Asia",
		Name:      "Armenia",
		Capital:   "Yerevan"},
	{
		Timezones: []string{"Africa/Luanda"},
		ISO2:      "AO",
		Continent: "Africa",
		Name:      "Angola",
		Capital:   "Luanda"},
	{
		Timezones: []string{"America/Argentina/Buenos_Aires", "America/Argentina/Cordoba", "America/Argentina/Jujuy", "America/Argentina/Tucuman", "America/Argentina/Catamarca", "America/Argentina/La_Rioja", "America/Argentina/San_Juan", "America/Argentina/Mendoza", "America/Argentina/Rio_Gallegos", "America/Argentina/Ushuaia"},
		ISO2:      "AR",
		Continent: "South America",
		Name:      "Argentina",
		Capital:   "Buenos Aires"},
	{
		Timezones: []string{"Europe/Vienna"},
		ISO2:      "AT",
		Continent: "Europe",
		Name:      "Austria",
		Capital:   "Vienna"},
	{
		Timezones: []string{"Australia/Lord_Howe", "Australia/Hobart", "Australia/Currie", "Australia/Melbourne", "Australia/Sydney", "Australia/Broken_Hill", "Australia/Brisbane", "Australia/Lindeman", "Australia/Adelaide", "Australia/Darwin", "Australia/Perth"},
		ISO2:      "AU",
		Continent: "Oceania",
		Name:      "Australia",
		Capital:   "Canberra"},
	{
		Timezones: []string{"Asia/Baku"},
		ISO2:      "AZ",
		Continent: "Asia",
		Name:      "Azerbaijan",
		Capital:   "Baku"},
	{
		Timezones: []string{"America/Barbados"},
		ISO2:      "BB",
		Continent: "North America",
		Name:      "Barbados",
		Capital:   "Bridgetown"},
	{
		Timezones: []string{"Asia/Dhaka"},
		ISO2:      "BD",
		Continent: "Asia",
		Name:      "Bangladesh",
		Capital:   "Dhaka"},
	{
		Timezones: []string{"Europe/Brussels"},
		ISO2:      "BE",
		Continent: "Europe",
		Name:      "Belgium",
		Capital:   "Brussels"},
	{
		Timezones: []string{"Africa/Ouagadougou"},
		ISO2:      "BF",
		Continent: "Africa",
		Name:      "Burkina Faso",
		Capital:   "Ouagadougou"},
	{
		Timezones: []string{"Europe/Sofia"},
		ISO2:      "BG",
		Continent: "Europe",
		Name:      "Bulgaria",
		Capital:   "Sofia"},
	{
		Timezones: []string{"Asia/Bahrain"},
		ISO2:      "BH",
		Continent: "Asia",
		Name:      "Bahrain",
		Capital:   "Manama"},
	{
		Timezones: []string{"Africa/Bujumbura"},
		ISO2:      "BI",
		Continent: "Africa",
		Name:      "Burundi",
		Capital:   "Bujumbura"},
	{
		Timezones: []string{"Africa/Porto-Novo"},
		ISO2:      "BJ",
		Continent: "Africa",
		Name:      "Benin",
		Capital:   "Porto-Novo"},
	{
		Timezones: []string{"Asia/Brunei"},
		ISO2:      "BN",
		Continent: "Asia",
		Name:      "Brunei Darussalam",
		Capital:   "Bandar Seri Begawan"},
	{
		Timezones: []string{"America/La_Paz"},
		ISO2:      "BO",
		Continent: "South America",
		Name:      "Bolivia",
		Capital:   "Sucre"},
	{
		Timezones: []string{"America/Noronha", "America/Belem", "America/Fortaleza", "America/Recife", "America/Araguaina", "America/Maceio", "America/Bahia", "America/Sao_Paulo", "America/Campo_Grande", "America/Cuiaba", "America/Porto_Velho", "America/Boa_Vista", "America/Manaus", "America/Eirunepe", "America/Rio_Branco"},
		ISO2:      "BR",
		Continent: "South America",
		Name:      "Brazil",
		Capital:   "Brasília"},
	{
		Timezones: []string{"America/Nassau"},
		ISO2:      "BS",
		Continent: "North America",
		Name:      "Bahamas",
		Capital:   "Nassau"},
	{
		Timezones: []string{"Asia/Thimphu"},
		ISO2:      "BT",
		Continent: "Asia",
		Name:      "Bhutan",
		Capital:   "Thimphu"},
	{
		Timezones: []string{"Africa/Gaborone"},
		ISO2:      "BW",
		Continent: "Africa",
		Name:      "Botswana",
		Capital:   "Gaborone"},
	{
		Timezones: []string{"Europe/Minsk"},
		ISO2:      "BY",
		Continent: "Europe",
		Name:      "Belarus",
		Capital:   "Minsk"},
	{
		Timezones: []string{"America/Belize"},
		ISO2:      "BZ",
		Continent: "North America",
		Name:      "Belize",
		Capital:   "Belmopan"},
	{
		Timezones: []string{"America/St_Johns", "America/Halifax", "America/Glace_Bay", "America/Moncton", "America/Goose_Bay", "America/Blanc-Sablon", "America/Montreal", "America/Toronto", "America/Nipigon", "America/Thunder_Bay", "America/Pangnirtung", "America/Iqaluit", "America/Atikokan", "America/Rankin_Inlet", "America/Winnipeg", "America/Rainy_River", "America/Cambridge_Bay", "America/Regina", "America/Swift_Current", "America/Edmonton", "America/Yellowknife", "America/Inuvik", "America/Dawson_Creek", "America/Vancouver", "America/Whitehorse", "America/Dawson"},
		ISO2:      "CA",
		Continent: "North America",
		Name:      "Canada",
		Capital:   "Ottawa"},
	{
		Timezones: []string{"Africa/Kinshasa", "Africa/Lubumbashi"},
		ISO2:      "CD",
		Continent: "Africa",
		Name:      "Democratic Republic of the Congo",
		Capital:   "Kinshasa"},
	{
		Timezones: []string{"Africa/Brazzaville"},
		ISO2:      "CG",
		Continent: "Africa",
		Name:      "Republic of the Congo",
		Capital:   "Brazzaville"},
	{
		Timezones: []string{"Africa/Abidjan"},
		ISO2:      "CI",
		Continent: "Africa",
		Name:      "Ivory Coast",
		Capital:   "Yamoussoukro"},
	{
		Timezones: []string{"America/Santiago", "Pacific/Easter"},
		ISO2:      "CL",
		Continent: "South America",
		Name:      "Chile",
		Capital:   "Santiago"},
	{
		Timezones: []string{"Africa/Douala"},
		ISO2:      "CM",
		Continent: "Africa",
		Name:      "Cameroon",
		Capital:   "Yaoundé"},
	{
		Timezones: []string{"Asia/Shanghai", "Asia/Harbin", "Asia/Chongqing", "Asia/Urumqi", "Asia/Kashgar"},
		ISO2:      "CN",
		Continent: "Asia",
		Name:      "People's Republic of China",
		Capital:   "Beijing"},
	{
		Timezones: []string{"America/Bogota"},
		ISO2:      "CO",
		Continent: "South America",
		Name:      "Colombia",
		Capital:   "Bogotá"},
	{
		Timezones: []string{"America/Costa_Rica"},
		ISO2:      "CR",
		Continent: "North America",
		Name:      "Costa Rica",
		Capital:   "San José"},
	{
		Timezones: []string{"America/Havana"},
		ISO2:      "CU",
		Continent: "North America",
		Name:      "Cuba",
		Capital:   "Havana"},
	{
		Timezones: []string{"Atlantic/Cape_Verde"},
		ISO2:      "CV",
		Continent: "Africa",
		Name:      "Cape Verde",
		Capital:   "Praia"},
	{
		Timezones: []string{"Asia/Nicosia"},
		ISO2:      "CY",
		Continent: "Asia",
		Name:      "Cyprus",
		Capital:   "Nicosia"},
	{
		Timezones: []string{"Europe/Prague"},
		ISO2:      "CZ",
		Continent: "Europe",
		Name:      "Czech Republic",
		Capital:   "Prague"},
	{
		Timezones: []string{"Europe/Berlin"},
		ISO2:      "DE",
		Continent: "Europe",
		Name:      "Germany",
		Capital:   "Berlin"},
	{
		Timezones: []string{"Africa/Djibouti"},
		ISO2:      "DJ",
		Continent: "Africa",
		Name:      "Djibouti",
		Capital:   "Djibouti City"},
	{
		Timezones: []string{"Europe/Copenhagen"},
		ISO2:      "DK",
		Continent: "Europe",
		Name:      "Denmark",
		Capital:   "Copenhagen"},
	{
		Timezones: []string{"America/Dominica"},
		ISO2:      "DM",
		Continent: "North America",
		Name:      "Dominica",
		Capital:   "Roseau"},
	{
		Timezones: []string{"America/Santo_Domingo"},
		ISO2:      "DO",
		Continent: "North America",
		Name:      "Dominican Republic",
		Capital:   "Santo Domingo"},
	{
		Timezones: []string{"America/Guayaquil", "Pacific/Galapagos"},
		ISO2:      "EC",
		Continent: "South America",
		Name:      "Ecuador",
		Capital:   "Quito"},
	{
		Timezones: []string{"Europe/Tallinn"},
		ISO2:      "EE",
		Continent: "Europe",
		Name:      "Estonia",
		Capital:   "Tallinn"},
	{
		Timezones: []string{"Africa/Cairo"},
		ISO2:      "EG",
		Continent: "Africa",
		Name:      "Egypt",
		Capital:   "Cairo"},
	{
		Timezones: []string{"Africa/Asmera"},
		ISO2:      "ER",
		Continent: "Africa",
		Name:      "Eritrea",
		Capital:   "Asmara"},
	{
		Timezones: []string{"Africa/Addis_Ababa"},
		ISO2:      "ET",
		Continent: "Africa",
		Name:      "Ethiopia",
		Capital:   "Addis Ababa"},
	{
		Timezones: []string{"Europe/Helsinki"},
		ISO2:      "FI",
		Continent: "Europe",
		Name:      "Finland",
		Capital:   "Helsinki"},
	{
		Timezones: []string{"Pacific/Fiji"},
		ISO2:      "FJ",
		Continent: "Oceania",
		Name:      "Fiji",
		Capital:   "Suva"},
	{
		Timezones: []string{"Europe/Paris"},
		ISO2:      "FR",
		Continent: "Europe",
		Name:      "France",
		Capital:   "Paris"},
	{
		Timezones: []string{"Africa/Libreville"},
		ISO2:      "GA",
		Continent: "Africa",
		Name:      "Gabon",
		Capital:   "Libreville"},
	{
		Timezones: []string{"Asia/Tbilisi"},
		ISO2:      "GE",
		Continent: "Asia",
		Name:      "Georgia",
		Capital:   "Tbilisi"},
	{
		Timezones: []string{"Africa/Accra"},
		ISO2:      "GH",
		Continent: "Africa",
		Name:      "Ghana",
		Capital:   "Accra"},
	{
		Timezones: []string{"Africa/Banjul"},
		ISO2:      "GM",
		Continent: "Africa",
		Name:      "The Gambia",
		Capital:   "Banjul"},
	{
		Timezones: []string{"Africa/Conakry"},
		ISO2:      "GN",
		Continent: "Africa",
		Name:      "Guinea",
		Capital:   "Conakry"},
	{
		Timezones: []string{"Europe/Athens"},
		ISO2:      "GR",
		Continent: "Europe",
		Name:      "Greece",
		Capital:   "Athens"},
	{
		Timezones: []string{"America/Guatemala"},
		ISO2:      "GT",
		Continent: "North America",
		Name:      "Guatemala",
		Capital:   "Guatemala City"},
	{
		Timezones: []string{"America/Guatemala"},
		ISO2:      "HT",
		Continent: "North America",
		Name:      "Haiti",
		Capital:   "Port-au-Prince"},
	{
		Timezones: []string{"Africa/Bissau"},
		ISO2:      "GW",
		Continent: "Africa",
		Name:      "Guinea-Bissau",
		Capital:   "Bissau"},
	{
		Timezones: []string{"America/Guyana"},
		ISO2:      "GY",
		Continent: "South America",
		Name:      "Guyana",
		Capital:   "Georgetown"},
	{
		Timezones: []string{"America/Tegucigalpa"},
		ISO2:      "HN",
		Continent: "North America",
		Name:      "Honduras",
		Capital:   "Tegucigalpa"},
	{
		Timezones: []string{"Europe/Budapest"},
		ISO2:      "HU",
		Continent: "Europe",
		Name:      "Hungary",
		Capital:   "Budapest"},
	{
		Timezones: []string{"Asia/Jakarta", "Asia/Pontianak", "Asia/Makassar", "Asia/Jayapura"},
		ISO2:      "ID",
		Continent: "Asia",
		Name:      "Indonesia",
		Capital:   "Jakarta"},
	{
		Timezones: []string{"Europe/Dublin"},
		ISO2:      "IE",
		Continent: "Europe",
		Name:      "Republic of Ireland",
		Capital:   "Dublin"},
	{
		Timezones: []string{"Asia/Jerusalem"},
		ISO2:      "IL",
		Continent: "Asia",
		Name:      "Israel",
		Capital:   "Jerusalem"},
	{
		Timezones: []string{"Asia/Calcutta"},
		ISO2:      "IN",
		Continent: "Asia",
		Name:      "India",
		Capital:   "New Delhi"},
	{
		Timezones: []string{"Asia/Baghdad"},
		ISO2:      "IQ",
		Continent: "Asia",
		Name:      "Iraq",
		Capital:   "Baghdad"},
	{
		Timezones: []string{"Asia/Tehran"},
		ISO2:      "IR",
		Continent: "Asia",
		Name:      "Iran",
		Capital:   "Tehran"},
	{
		Timezones: []string{"Atlantic/Reykjavik"},
		ISO2:      "IS",
		Continent: "Europe",
		Name:      "Iceland",
		Capital:   "Reykjavik"},
	{
		Timezones: []string{"Europe/Rome"},
		ISO2:      "IT",
		Continent: "Europe",
		Name:      "Italy",
		Capital:   "Rome"},
	{
		Timezones: []string{"America/Jamaica"},
		ISO2:      "JM",
		Continent: "North America",
		Name:      "Jamaica",
		Capital:   "Kingston"},
	{
		Timezones: []string{"Asia/Amman"},
		ISO2:      "JO",
		Continent: "Asia",
		Name:      "Jordan",
		Capital:   "Amman"},
	{
		Timezones: []string{"Asia/Tokyo"},
		ISO2:      "JP",
		Continent: "Asia",
		Name:      "Japan",
		Capital:   "Tokyo"},
	{
		Timezones: []string{"Africa/Nairobi"},
		ISO2:      "KE",
		Continent: "Africa",
		Name:      "Kenya",
		Capital:   "Nairobi"},
	{
		Timezones: []string{"Asia/Bishkek"},
		ISO2:      "KG",
		Continent: "Asia",
		Name:      "Kyrgyzstan",
		Capital:   "Bishkek"},
	{
		Timezones: []string{"Pacific/Tarawa", "Pacific/Enderbury", "Pacific/Kiritimati"},
		ISO2:      "KI",
		Continent: "Oceania",
		Name:      "Kiribati",
		Capital:   "Tarawa"},
	{
		Timezones: []string{"Asia/Pyongyang"},
		ISO2:      "KP",
		Continent: "Asia",
		Name:      "North Korea",
		Capital:   "Pyongyang"},
	{
		Timezones: []string{"Asia/Seoul"},
		ISO2:      "KR",
		Continent: "Asia",
		Name:      "South Korea",
		Capital:   "Seoul"},
	{
		Timezones: []string{"Asia/Kuwait"},
		ISO2:      "KW",
		Continent: "Asia",
		Name:      "Kuwait",
		Capital:   "Kuwait City"},
	{
		Timezones: []string{"Asia/Beirut"},
		ISO2:      "LB",
		Continent: "Asia",
		Name:      "Lebanon",
		Capital:   "Beirut"},
	{
		Timezones: []string{"Europe/Vaduz"},
		ISO2:      "LI",
		Continent: "Europe",
		Name:      "Liechtenstein",
		Capital:   "Vaduz"},
	{
		Timezones: []string{"Africa/Monrovia"},
		ISO2:      "LR",
		Continent: "Africa",
		Name:      "Liberia",
		Capital:   "Monrovia"},
	{
		Timezones: []string{"Africa/Maseru"},
		ISO2:      "LS",
		Continent: "Africa",
		Name:      "Lesotho",
		Capital:   "Maseru"},
	{
		Timezones: []string{"Europe/Vilnius"},
		ISO2:      "LT",
		Continent: "Europe",
		Name:      "Lithuania",
		Capital:   "Vilnius"},
	{
		Timezones: []string{"Europe/Luxembourg"},
		ISO2:      "LU",
		Continent: "Europe",
		Name:      "Luxembourg",
		Capital:   "Luxembourg City"},
	{
		Timezones: []string{"Europe/Riga"},
		ISO2:      "LV",
		Continent: "Europe",
		Name:      "Latvia",
		Capital:   "Riga"},
	{
		Timezones: []string{"Africa/Tripoli"},
		ISO2:      "LY",
		Continent: "Africa",
		Name:      "Libya",
		Capital:   "Tripoli"},
	{
		Timezones: []string{"Indian/Antananarivo"},
		ISO2:      "MG",
		Continent: "Africa",
		Name:      "Madagascar",
		Capital:   "Antananarivo"},
	{
		Timezones: []string{"Pacific/Majuro", "Pacific/Kwajalein"},
		ISO2:      "MH",
		Continent: "Oceania",
		Name:      "Marshall Islands",
		Capital:   "Majuro"},
	{
		Timezones: []string{"Europe/Skopje"},
		ISO2:      "MK",
		Continent: "Europe",
		Name:      "Macedonia",
		Capital:   "Skopje"},
	{
		Timezones: []string{"Africa/Bamako"},
		ISO2:      "ML",
		Continent: "Africa",
		Name:      "Mali",
		Capital:   "Bamako"},
	{
		Timezones: []string{"Asia/Rangoon"},
		ISO2:      "MM",
		Continent: "Asia",
		Name:      "Myanmar",
		Capital:   "Naypyidaw"},
	{
		Timezones: []string{"Asia/Ulaanbaatar", "Asia/Hovd", "Asia/Choibalsan"},
		ISO2:      "MN",
		Continent: "Asia",
		Name:      "Mongolia",
		Capital:   "Ulaanbaatar"},
	{
		Timezones: []string{"Africa/Nouakchott"},
		ISO2:      "MR",
		Continent: "Africa",
		Name:      "Mauritania",
		Capital:   "Nouakchott"},
	{
		Timezones: []string{"Europe/Malta"},
		ISO2:      "MT",
		Continent: "Europe",
		Name:      "Malta",
		Capital:   "Valletta"},
	{
		Timezones: []string{"Indian/Mauritius"},
		ISO2:      "MU",
		Continent: "Africa",
		Name:      "Mauritius",
		Capital:   "Port Louis"},
	{
		Timezones: []string{"Indian/Maldives"},
		ISO2:      "MV",
		Continent: "Asia",
		Name:      "Maldives",
		Capital:   "Malé"},
	{
		Timezones: []string{"Africa/Blantyre"},
		ISO2:      "MW",
		Continent: "Africa",
		Name:      "Malawi",
		Capital:   "Lilongwe"},
	{
		Timezones: []string{"America/Mexico_City", "America/Cancun", "America/Merida", "America/Monterrey", "America/Mazatlan", "America/Chihuahua", "America/Hermosillo", "America/Tijuana"},
		ISO2:      "MX",
		Continent: "North America",
		Name:      "Mexico",
		Capital:   "Mexico City"},
	{
		Timezones: []string{"Asia/Kuala_Lumpur", "Asia/Kuching"},
		ISO2:      "MY",
		Continent: "Asia",
		Name:      "Malaysia",
		Capital:   "Kuala Lumpur"},
	{
		Timezones: []string{"Africa/Maputo"},
		ISO2:      "MZ",
		Continent: "Africa",
		Name:      "Mozambique",
		Capital:   "Maputo"},
	{
		Timezones: []string{"Africa/Windhoek"},
		ISO2:      "NA",
		Continent: "Africa",
		Name:      "Namibia",
		Capital:   "Windhoek"},
	{
		Timezones: []string{"Africa/Niamey"},
		ISO2:      "NE",
		Continent: "Africa",
		Name:      "Niger",
		Capital:   "Niamey"},
	{
		Timezones: []string{"Africa/Lagos"},
		ISO2:      "NG",
		Continent: "Africa",
		Name:      "Nigeria",
		Capital:   "Abuja"},
	{
		Timezones: []string{"America/Managua"},
		ISO2:      "NI",
		Continent: "North America",
		Name:      "Nicaragua",
		Capital:   "Managua"},
	{
		Timezones: []string{"Europe/Amsterdam"},
		ISO2:      "NL",
		Continent: "Europe",
		Name:      "Kingdom of the Netherlands",
		Capital:   "Amsterdam"},
	{
		Timezones: []string{"Europe/Oslo"},
		ISO2:      "NO",
		Continent: "Europe",
		Name:      "Norway",
		Capital:   "Oslo"},
	{
		Timezones: []string{"Asia/Katmandu"},
		ISO2:      "NP",
		Continent: "Asia",
		Name:      "Nepal",
		Capital:   "Kathmandu"},
	{
		Timezones: []string{"Pacific/Nauru"},
		ISO2:      "NR",
		Continent: "Oceania",
		Name:      "Nauru",
		Capital:   "Yaren"},
	{
		Timezones: []string{"Pacific/Auckland", "Pacific/Chatham"},
		ISO2:      "NZ",
		Continent: "Oceania",
		Name:      "New Zealand",
		Capital:   "Wellington"},
	{
		Timezones: []string{"Asia/Muscat"},
		ISO2:      "OM",
		Continent: "Asia",
		Name:      "Oman",
		Capital:   "Muscat"},
	{
		Timezones: []string{"America/Panama"},
		ISO2:      "PA",
		Continent: "North America",
		Name:      "Panama",
		Capital:   "Panama City"},
	{
		Timezones: []string{"America/Lima"},
		ISO2:      "PE",
		Continent: "South America",
		Name:      "Peru",
		Capital:   "Lima"},
	{
		Timezones: []string{"Pacific/Port_Moresby"},
		ISO2:      "PG",
		Continent: "Oceania",
		Name:      "Papua New Guinea",
		Capital:   "Port Moresby"},
	{
		Timezones: []string{"Asia/Manila"},
		ISO2:      "PH",
		Continent: "Asia",
		Name:      "Philippines",
		Capital:   "Manila"},
	{
		Timezones: []string{"Asia/Karachi"},
		ISO2:      "PK",
		Continent: "Asia",
		Name:      "Pakistan",
		Capital:   "Islamabad"},
	{
		Timezones: []string{"Europe/Warsaw"},
		ISO2:      "PL",
		Continent: "Europe",
		Name:      "Poland",
		Capital:   "Warsaw"},
	{
		Timezones: []string{"Europe/Lisbon", "Atlantic/Madeira", "Atlantic/Azores"},
		ISO2:      "PT",
		Continent: "Europe",
		Name:      "Portugal",
		Capital:   "Lisbon"},
	{
		Timezones: []string{"Pacific/Palau"},
		ISO2:      "PW",
		Continent: "Oceania",
		Name:      "Palau",
		Capital:   "Ngerulmud"},
	{
		Timezones: []string{"America/Asuncion"},
		ISO2:      "PY",
		Continent: "South America",
		Name:      "Paraguay",
		Capital:   "Asunción"},
	{
		Timezones: []string{"Asia/Qatar"},
		ISO2:      "QA",
		Continent: "Asia",
		Name:      "Qatar",
		Capital:   "Doha"},
	{
		Timezones: []string{"Europe/Bucharest"},
		ISO2:      "RO",
		Continent: "Europe",
		Name:      "Romania",
		Capital:   "Bucharest"},
	{
		Timezones: []string{"Europe/Kaliningrad", "Europe/Moscow", "Europe/Volgograd", "Europe/Samara", "Asia/Yekaterinburg", "Asia/Omsk", "Asia/Novosibirsk", "Asia/Krasnoyarsk", "Asia/Irkutsk", "Asia/Yakutsk", "Asia/Vladivostok", "Asia/Sakhalin", "Asia/Magadan", "Asia/Kamchatka", "Asia/Anadyr"},
		ISO2:      "RU",
		Continent: "Europe",
		Name:      "Russia",
		Capital:   "Moscow"},
	{
		Timezones: []string{"Africa/Kigali"},
		ISO2:      "RW",
		Continent: "Africa",
		Name:      "Rwanda",
		Capital:   "Kigali"},
	{
		Timezones: []string{"Asia/Riyadh"},
		ISO2:      "SA",
		Continent: "Asia",
		Name:      "Saudi Arabia",
		Capital:   "Riyadh"},
	{
		Timezones: []string{"Pacific/Guadalcanal"},
		ISO2:      "SB",
		Continent: "Oceania",
		Name:      "Solomon Islands",
		Capital:   "Honiara"},
	{
		Timezones: []string{"Indian/Mahe"},
		ISO2:      "SC",
		Continent: "Africa",
		Name:      "Seychelles",
		Capital:   "Victoria"},
	{
		Timezones: []string{"Africa/Khartoum"},
		ISO2:      "SD",
		Continent: "Africa",
		Name:      "Sudan",
		Capital:   "Khartoum"},
	{
		Timezones: []string{"Europe/Stockholm"},
		ISO2:      "SE",
		Continent: "Europe",
		Name:      "Sweden",
		Capital:   "Stockholm"},
	{
		Timezones: []string{"Asia/Singapore"},
		ISO2:      "SG",
		Continent: "Asia",
		Name:      "Singapore",
		Capital:   "Singapore"},
	{
		Timezones: []string{"Europe/Ljubljana"},
		ISO2:      "SI",
		Continent: "Europe",
		Name:      "Slovenia",
		Capital:   "Ljubljana"},
	{
		Timezones: []string{"Europe/Bratislava"},
		ISO2:      "SK",
		Continent: "Europe",
		Name:      "Slovakia",
		Capital:   "Bratislava"},
	{
		Timezones: []string{"Africa/Freetown"},
		ISO2:      "SL",
		Continent: "Africa",
		Name:      "Sierra Leone",
		Capital:   "Freetown"},
	{
		Timezones: []string{"Europe/San_Marino"},
		ISO2:      "SM",
		Continent: "Europe",
		Name:      "San Marino",
		Capital:   "San Marino"},
	{
		Timezones: []string{"Africa/Dakar"},
		ISO2:      "SN",
		Continent: "Africa",
		Name:      "Senegal",
		Capital:   "Dakar"},
	{
		Timezones: []string{"Africa/Mogadishu"},
		ISO2:      "SO",
		Continent: "Africa",
		Name:      "Somalia",
		Capital:   "Mogadishu"},
	{
		Timezones: []string{"America/Paramaribo"},
		ISO2:      "SR",
		Continent: "South America",
		Name:      "Suriname",
		Capital:   "Paramaribo"},
	{
		Timezones: []string{"Africa/Sao_Tome"},
		ISO2:      "ST",
		Continent: "Africa",
		Name:      "São Tomé and Príncipe",
		Capital:   "São Tomé"},
	{
		Timezones: []string{"Asia/Damascus"},
		ISO2:      "SY",
		Continent: "Asia",
		Name:      "Syria",
		Capital:   "Damascus"},
	{
		Timezones: []string{"Africa/Lome"},
		ISO2:      "TG",
		Continent: "Africa",
		Name:      "Togo",
		Capital:   "Lomé"},
	{
		Timezones: []string{"Asia/Bangkok"},
		ISO2:      "TH",
		Continent: "Asia",
		Name:      "Thailand",
		Capital:   "Bangkok"},
	{
		Timezones: []string{"Asia/Dushanbe"},
		ISO2:      "TJ",
		Continent: "Asia",
		Name:      "Tajikistan",
		Capital:   "Dushanbe"},
	{
		Timezones: []string{"Asia/Ashgabat"},
		ISO2:      "TM",
		Continent: "Asia",
		Name:      "Turkmenistan",
		Capital:   "Ashgabat"},
	{
		Timezones: []string{"Africa/Tunis"},
		ISO2:      "TN",
		Continent: "Africa",
		Name:      "Tunisia",
		Capital:   "Tunis"},
	{
		Timezones: []string{"Pacific/Tongatapu"},
		ISO2:      "TO",
		Continent: "Oceania",
		Name:      "Tonga",
		Capital:   "Nuku'alofa"},
	{
		Timezones: []string{"Europe/Istanbul"},
		ISO2:      "TR",
		Continent: "Asia",
		Name:      "Turkey",
		Capital:   "Ankara"},
	{
		Timezones: []string{"America/Port_of_Spain"},
		ISO2:      "TT",
		Continent: "North America",
		Name:      "Trinidad and Tobago",
		Capital:   "Port of Spain"},
	{
		Timezones: []string{"Pacific/Funafuti"},
		ISO2:      "TV",
		Continent: "Oceania",
		Name:      "Tuvalu",
		Capital:   "Funafuti"},
	{
		Timezones: []string{"Africa/Dar_es_Salaam"},
		ISO2:      "TZ",
		Continent: "Africa",
		Name:      "Tanzania",
		Capital:   "Dodoma"},
	{
		Timezones: []string{"Europe/Kiev", "Europe/Uzhgorod", "Europe/Zaporozhye", "Europe/Simferopol"},
		ISO2:      "UA",
		Continent: "Europe",
		Name:      "Ukraine",
		Capital:   "Kiev"},
	{
		Timezones: []string{"Africa/Kampala"},
		ISO2:      "UG",
		Continent: "Africa",
		Name:      "Uganda",
		Capital:   "Kampala"},
	{
		Timezones: []string{"America/New_York", "America/Detroit", "America/Kentucky/Louisville", "America/Kentucky/Monticello", "America/Indiana/Indianapolis", "America/Indiana/Marengo", "America/Indiana/Knox", "America/Indiana/Vevay", "America/Chicago", "America/Indiana/Vincennes", "America/Indiana/Petersburg", "America/Menominee", "America/North_Dakota/Center", "America/North_Dakota/New_Salem", "America/Denver", "America/Boise", "America/Shiprock", "America/Phoenix", "America/Los_Angeles", "America/Anchorage", "America/Juneau", "America/Yakutat", "America/Nome", "America/Adak", "Pacific/Honolulu"},
		ISO2:      "US",
		Continent: "North America",
		Name:      "United States",
		Capital:   "Washington, D.C."},
	{
		Timezones: []string{"America/Montevideo"},
		ISO2:      "UY",
		Continent: "South America",
		Name:      "Uruguay",
		Capital:   "Montevideo"},
	{
		Timezones: []string{"Asia/Samarkand", "Asia/Tashkent"},
		ISO2:      "UZ",
		Continent: "Asia",
		Name:      "Uzbekistan",
		Capital:   "Tashkent"},
	{
		Timezones: []string{"Europe/Vatican"},
		ISO2:      "VA",
		Continent: "Europe",
		Name:      "Vatican City",
		Capital:   "Vatican City"},
	{
		Timezones: []string{"America/Caracas"},
		ISO2:      "VE",
		Continent: "South America",
		Name:      "Venezuela",
		Capital:   "Caracas"},
	{
		Timezones: []string{"Asia/Saigon"},
		ISO2:      "VN",
		Continent: "Asia",
		Name:      "Vietnam",
		Capital:   "Hanoi"},
	{
		Timezones: []string{"Pacific/Efate"},
		ISO2:      "VU",
		Continent: "Oceania",
		Name:      "Vanuatu",
		Capital:   "Port Vila"},
	{
		Timezones: []string{"Asia/Aden"},
		ISO2:      "YE",
		Continent: "Asia",
		Name:      "Yemen",
		Capital:   "Sana'a"},
	{
		Timezones: []string{"Africa/Lusaka"},
		ISO2:      "ZM",
		Continent: "Africa",
		Name:      "Zambia",
		Capital:   "Lusaka"},
	{
		Timezones: []string{"Africa/Harare"},
		ISO2:      "ZW",
		Continent: "Africa",
		Name:      "Zimbabwe",
		Capital:   "Harare"},
	{
		Timezones: []string{"Africa/Algiers"},
		ISO2:      "DZ",
		Continent: "Africa",
		Name:      "Algeria",
		Capital:   "Algiers"},
	{
		Timezones: []string{"Europe/Sarajevo"},
		ISO2:      "BA",
		Continent: "Europe",
		Name:      "Bosnia and Herzegovina",
		Capital:   "Sarajevo"},
	{
		Timezones: []string{"Asia/Phnom_Penh"},
		ISO2:      "KH",
		Continent: "Asia",
		Name:      "Cambodia",
		Capital:   "Phnom Penh"},
	{
		Timezones: []string{"Africa/Bangui"},
		ISO2:      "CF",
		Continent: "Africa",
		Name:      "Central African Republic",
		Capital:   "Bangui"},
	{
		Timezones: []string{"Africa/Ndjamena"},
		ISO2:      "TD",
		Continent: "Africa",
		Name:      "Chad",
		Capital:   "N'Djamena"},
	{
		Timezones: []string{"Indian/Comoro"},
		ISO2:      "KM",
		Continent: "Africa",
		Name:      "Comoros",
		Capital:   "Moroni"},
	{
		Timezones: []string{"Europe/Zagreb"},
		ISO2:      "HR",
		Continent: "Europe",
		Name:      "Croatia",
		Capital:   "Zagreb"},
	{
		Timezones: []string{"Asia/Dili"},
		ISO2:      "TL",
		Continent: "Asia",
		Name:      "East Timor",
		Capital:   "Dili"},
	{
		Timezones: []string{"America/El_Salvador"},
		ISO2:      "SV",
		Continent: "North America",
		Name:      "El Salvador",
		Capital:   "San Salvador"},
	{
		Timezones: []string{"Africa/Malabo"},
		ISO2:      "GQ",
		Continent: "Africa",
		Name:      "Equatorial Guinea",
		Capital:   "Malabo"},
	{
		Timezones: []string{"America/Grenada"},
		ISO2:      "GD",
		Continent: "North America",
		Name:      "Grenada",
		Capital:   "St. George's"},
	{
		Timezones: []string{"Asia/Almaty", "Asia/Qyzylorda", "Asia/Aqtobe", "Asia/Aqtau", "Asia/Oral"},
		ISO2:      "KZ",
		Continent: "Asia",
		Name:      "Kazakhstan",
		Capital:   "Astana"},
	{
		Timezones: []string{"Asia/Vientiane"},
		ISO2:      "LA",
		Continent: "Asia",
		Name:      "Laos",
		Capital:   "Vientiane"},
	{
		Timezones: []string{"Pacific/Truk", "Pacific/Ponape", "Pacific/Kosrae"},
		ISO2:      "FM",
		Continent: "Oceania",
		Name:      "Federated States of Micronesia",
		Capital:   "Palikir"},
	{
		Timezones: []string{"Europe/Chisinau"},
		ISO2:      "MD",
		Continent: "Europe",
		Name:      "Moldova",
		Capital:   "Chișinău"},
	{
		Timezones: []string{"Europe/Monaco"},
		ISO2:      "MC",
		Continent: "Europe",
		Name:      "Monaco",
		Capital:   "Monaco"},
	{
		Timezones: []string{"Europe/Podgorica"},
		ISO2:      "ME",
		Continent: "Europe",
		Name:      "Montenegro",
		Capital:   "Podgorica"},
	{
		Timezones: []string{"Africa/Casablanca"},
		ISO2:      "MA",
		Continent: "Africa",
		Name:      "Morocco",
		Capital:   "Rabat"},
	{
		Timezones: []string{"America/St_Kitts"},
		ISO2:      "KN",
		Continent: "North America",
		Name:      "Saint Kitts and Nevis",
		Capital:   "Basseterre"},
	{
		Timezones: []string{"America/St_Lucia"},
		ISO2:      "LC",
		Continent: "North America",
		Name:      "Saint Lucia",
		Capital:   "Castries"},
	{
		Timezones: []string{"America/St_Vincent"},
		ISO2:      "VC",
		Continent: "North America",
		Name:      "Saint Vincent and the Grenadines",
		Capital:   "Kingstown"},
	{
		Timezones: []string{"Pacific/Apia"},
		ISO2:      "WS",
		Continent: "Oceania",
		Name:      "Samoa",
		Capital:   "Apia"},
	{
		Timezones: []string{"Europe/Belgrade"},
		ISO2:      "RS",
		Continent: "Europe",
		Name:      "Serbia",
		Capital:   "Belgrade"},
	{
		Timezones: []string{"Africa/Johannesburg"},
		ISO2:      "ZA",
		Continent: "Africa",
		Name:      "South Africa",
		Capital:   "Pretoria"},
	{
		Timezones: []string{"Europe/Madrid", "Africa/Ceuta", "Atlantic/Canary"},
		ISO2:      "ES",
		Continent: "Europe",
		Name:      "Spain",
		Capital:   "Madrid"},
	{
		Timezones: []string{"Asia/Colombo"},
		ISO2:      "LK",
		Continent: "Asia",
		Name:      "Sri Lanka",
		Capital:   "Sri Jayewardenepura Kotte"},
	{
		Timezones: []string{"Africa/Mbabane"},
		ISO2:      "SZ",
		Continent: "Africa",
		Name:      "Swaziland",
		Capital:   "Mbabane"},
	{
		Timezones: []string{"Europe/Zurich"},
		ISO2:      "CH",
		Continent: "Europe",
		Name:      "Switzerland",
		Capital:   "Bern"},
	{
		Timezones: []string{"Asia/Dubai"},
		ISO2:      "AE",
		Continent: "Asia",
		Name:      "United Arab Emirates",
		Capital:   "Abu Dhabi"},
	{
		Timezones: []string{"Europe/London"},
		ISO2:      "GB",
		Continent: "Europe",
		Name:      "United Kingdom",
		Capital:   "London"},
	{
		Timezones: []string{"Asia/Taipei"},
		ISO2:      "TW",
		Continent: "Asia",
		Name:      "Taiwan",
		Capital:   "Taipei"},
	{
		Timezones: []string{"Asia/Hong_Kong"},
		ISO2:      "HK",
		Continent: "Asia",
		Name:      "Hong Kong",
		Capital:   "Hong Kong"},
	{
		Timezones: []string{"CET"},
		ISO2:      "XK",
		Continent: "Europe",
		Name:      "Kosovo",
		Capital:   "Pristina"},
	{
		Timezones: []string{"America/Puerto_Rico"},
		ISO2:      "PR",
		Continent: "America",
		Name:      "Puerto Rico",
		Capital:   "San Juan"},
	{
		Timezones: []string{"Europe/Mariehamn"},
		ISO2:      "AX",
		Name:      "Aland Islands",
		Capital:   "Mariehamn",
		Continent: "Europe"},
	{
		Timezones: []string{"Pacific/Pago_Pago"},
		ISO2:      "AS",
		Name:      "American Samoa",
		Capital:   "Pago Pago",
		Continent: "America"},
	{
		Timezones: []string{"America/Anguilla"},
		ISO2:      "AI",
		Name:      "Anguilla",
		Capital:   "The Valley",
		Continent: "America"},
	{
		Timezones: []string{"Antarctica/Casey", "Antarctica/Davis", "Antarctica/DumontDUrville", "Antarctica/Mawson", "Antarctica/McMurdo", "Antarctica/Palmer", "Antarctica/Rothera", "Antarctica/Syowa", "Antarctica/Troll", "Antarctica/Vostok"},
		ISO2:      "AQ",
		Name:      "Antarctica",
		Capital:   "Antarctica",
		Continent: "Antarctica"},
	{
		Timezones: []string{"America/Aruba"},
		ISO2:      "AW",
		Name:      "Aruba",
		Capital:   "Oranjestad",
		Continent: "America"},
	{
		Timezones: []string{"Atlantic/Bermuda"},
		ISO2:      "BM",
		Name:      "Bermuda",
		Capital:   "Hamilton",
		Continent: "America"},
	{
		Timezones: []string{"America/Kralendijk"},
		ISO2:      "BQ",
		Name:      "Bonaire, Saint Eustatius and Saba",
		Capital:   "Kralendijk",
		Continent: "America"},
	{
		Timezones: []string{"Indian/Chagos"},
		ISO2:      "IO",
		Name:      "British Indian Ocean Territory",
		Capital:   "Diego Garcia",
		Continent: "Africa"},
	{
		Timezones: []string{"America/Tortola"},
		ISO2:      "VG",
		Name:      "British Virgin Islands",
		Capital:   "Road Town",
		Continent: "America"},
	{
		Timezones: []string{"America/Cayman"},
		ISO2:      "KY",
		Name:      "Cayman Islands",
		Capital:   "George Town",
		Continent: "America"},
	{
		Timezones: []string{"Indian/Christmas"},
		ISO2:      "CX",
		Name:      "Christmas Island",
		Capital:   "Flying Fish Cove",
		Continent: "Oceania"},
	{
		Timezones: []string{"Indian/Cocos"},
		ISO2:      "CC",
		Name:      "Cocos Islands",
		Capital:   "West Island",
		Continent: "Oceania"},
	{
		Timezones: []string{"Pacific/Rarotonga"},
		ISO2:      "CK",
		Name:      "Cook Islands",
		Capital:   "Avarua",
		Continent: "Oceania"},
	{
		Timezones: []string{"America/Curacao"},
		ISO2:      "CW",
		Name:      "Curacao",
		Capital:   "Willemstad",
		Continent: "America"},
	{
		Timezones: []string{"Atlantic/Stanley"},
		ISO2:      "FK",
		Name:      "Falkland Islands",
		Capital:   "Stanley",
		Continent: "America"},
	{
		Timezones: []string{"Atlantic/Faroe"},
		ISO2:      "FO",
		Name:      "Faroe Islands",
		Capital:   "Tórshavn",
		Continent: "Europe"},
	{
		Timezones: []string{"America/Cayenne"},
		ISO2:      "GF",
		Name:      "French Guiana",
		Capital:   "Cayenne",
		Continent: "America"},
	{
		Timezones: []string{"Pacific/Gambier", "Pacific/Marquesas", "Pacific/Tahiti"},
		ISO2:      "PF",
		Name:      "French Polynesia",
		Capital:   "Papeetē",
		Continent: "Oceania"},
	{
		Timezones: []string{"Indian/Kerguelen"},
		ISO2:      "TF",
		Name:      "French Southern and Antarctic Lands",
		Capital:   "Port-aux-Français",
		Continent: "Oceania"},
	{
		Timezones: []string{"Europe/Gibraltar"},
		ISO2:      "GI",
		Name:      "Gibraltar",
		Capital:   "Gibraltar",
		Continent: "Europe"},
	{
		Timezones: []string{"America/Danmarkshavn", "America/Godthab", "America/Scoresbysund", "America/Thule"},
		ISO2:      "GL",
		Name:      "Greenland",
		Capital:   "Nuuk",
		Continent: "America"},
	{
		Timezones: []string{"America/Guadeloupe"},
		ISO2:      "GP",
		Name:      "Guadeloupe",
		Capital:   "Basse-Terre",
		Continent: "America"},
	{
		Timezones: []string{"Pacific/Guam"},
		ISO2:      "GU",
		Name:      "Guam",
		Capital:   "Hagåtña",
		Continent: "Oceania"},
	{
		Timezones: []string{"Europe/Guernsey"},
		ISO2:      "GG",
		Name:      "Guernsey",
		Capital:   "St. Peter Port",
		Continent: "Europe"},
	{
		Timezones: []string{"Europe/Isle_of_Man"},
		ISO2:      "IM",
		Name:      "Isle of Man",
		Capital:   "Douglas",
		Continent: "Europe"},
	{
		Timezones: []string{"Europe/Jersey"},
		ISO2:      "JE",
		Name:      "Jersey",
		Capital:   "Saint Helier",
		Continent: "Europe"},
	{
		Timezones: []string{"Asia/Macau"},
		ISO2:      "MO",
		Name:      "Macao",
		Capital:   "Macao",
		Continent: "Asia"},
	{
		Timezones: []string{"America/Martinique"},
		ISO2:      "MQ",
		Name:      "Martinique",
		Capital:   "Fort-de-France",
		Continent: "America"},
	{
		Timezones: []string{"Indian/Mayotte"},
		ISO2:      "YT",
		Name:      "Mayotte",
		Capital:   "Mamoudzou",
		Continent: "Africa"},
	{
		Timezones: []string{"America/Montserrat"},
		ISO2:      "MS",
		Name:      "Montserrat",
		Capital:   "Plymouth",
		Continent: "America"},
	{
		Timezones: []string{"Pacific/Noumea"},
		ISO2:      "NC",
		Name:      "New Caledonia",
		Capital:   "Nouméa",
		Continent: "Oceania"},
	{
		Timezones: []string{"Pacific/Niue"},
		ISO2:      "NU",
		Name:      "Niue",
		Capital:   "Alofi",
		Continent: "Oceania"},
	{
		Timezones: []string{"Pacific/Norfolk"},
		ISO2:      "NF",
		Name:      "Norfolk Island",
		Capital:   "Kingston",
		Continent: "Oceania"},
	{
		Timezones: []string{"Pacific/Saipan"},
		ISO2:      "MP",
		Name:      "Northern Mariana Islands",
		Capital:   "Saipan",
		Continent: "Oceania"},
	{
		Timezones: []string{"Asia/Gaza", "Asia/Hebron"},
		ISO2:      "PS",
		Name:      "Palestinian Territory",
		Capital:   "Ramallah",
		Continent: "Asia"},
	{
		Timezones: []string{"Pacific/Pitcairn"},
		ISO2:      "PN",
		Name:      "Pitcairn",
		Capital:   "Adamstown",
		Continent: "Oceania"},
	{
		Timezones: []string{"Indian/Reunion"},
		ISO2:      "RE",
		Name:      "Reunion",
		Capital:   "Saint-Denis",
		Continent: "Africa"},
	{
		Timezones: []string{"America/St_Barthelemy"},
		ISO2:      "BL",
		Name:      "Saint Barthelemy",
		Capital:   "Gustavia",
		Continent: "America"},
	{
		Timezones: []string{"Atlantic/St_Helena"},
		ISO2:      "SH",
		Name:      "Saint Helena",
		Capital:   "Jamestown",
		Continent: "Africa"},
	{
		Timezones: []string{"America/Marigot"},
		ISO2:      "MF",
		Name:      "Saint Martin",
		Capital:   "Marigot",
		Continent: "America"},
	{
		Timezones: []string{"America/Miquelon"},
		ISO2:      "PM",
		Name:      "Saint Pierre and Miquelon",
		Capital:   "Saint-Pierre",
		Continent: "America"},
	{
		Timezones: []string{"America/Lower_Princes"},
		ISO2:      "SX",
		Name:      "Sint Maarten",
		Capital:   "Philipsburg",
		Continent: "America"},
	{
		Timezones: []string{"Atlantic/South_Georgia"},
		ISO2:      "GS",
		Name:      "South Georgia and the South Sandwich Islands",
		Capital:   "King Edward Point",
		Continent: "America"},
	{
		Timezones: []string{"Africa/Juba"},
		ISO2:      "SS",
		Name:      "South Sudan",
		Capital:   "Juba",
		Continent: "Africa"},
	{
		Timezones: []string{"Arctic/Longyearbyen"},
		ISO2:      "SJ",
		Name:      "Svalbard and Jan Mayen",
		Capital:   "Longyearbyen",
		Continent: "Europe"},
	{
		Timezones: []string{"Pacific/Fakaofo"},
		ISO2:      "TK",
		Name:      "Tokelau",
		Capital:   "Fakaofo",
		Continent: "Oceania"},
	{
		Timezones: []string{"America/Grand_Turk"},
		ISO2:      "TC",
		Name:      "Turks and Caicos Islands",
		Capital:   "Cockburn Town",
		Continent: "America"},
	{
		Timezones: []string{"America/St_Thomas"},
		ISO2:      "VI",
		Name:      "United States Virgin Islands",
		Capital:   "Charlotte Amalie",
		Continent: "America"},
	{
		Timezones: []string{"Pacific/Midway", "Pacific/Wake"},
		ISO2:      "UM",
		Name:      "United States Minor Outlying Islands",
		Capital:   "United States Minor Outlying Islands",
		Continent: "America"},
	{
		Timezones: []string{"Pacific/Wallis"},
		ISO2:      "WF",
		Name:      "Wallis and Futuna",
		Capital:   "Mata-Utu",
		Continent: "Oceania"},
	{
		Timezones: []string{"Africa/El_Aaiun"},
		ISO2:      "EH",
		Name:      "Western Sahara",
		Capital:   "El Aaiún",
		Continent: "Africa",
	},
}
