package models

type CountryMedal struct {
	Position     int64  `json Position`
	Country      string `json Country`
	GoldMedals   int64  `json GoldMedals`
	SilverMedals int64  `json SilverMedals`
	BronzeMedals int64  `json BronzeMedals`
	//AllMedals    int64  `json: AllMedals`
}

type AllCountry []CountryMedal

var AllCountrys = AllCountry{
	{
		Position:     1,
		Country:      "China",
		GoldMedals:   19,
		SilverMedals: 10,
		BronzeMedals: 11,
		//AllMedals:    31,
	},
	{
		Position:     2,
		Country:      "Japon",
		GoldMedals:   17,
		SilverMedals: 4,
		BronzeMedals: 7,
		//AllMedals:    25,
	},
	{
		Position:     3,
		Country:      "Estados Unidos",
		GoldMedals:   14,
		SilverMedals: 16,
		BronzeMedals: 11,
		//AllMedals:    38,
	},
	{
		Position:     4,
		Country:      "COR",
		GoldMedals:   10,
		SilverMedals: 14,
		BronzeMedals: 10,
		//AllMedals:    28,
	},
	{
		Position:     5,
		Country:      "Australia",
		GoldMedals:   9,
		SilverMedals: 2,
		BronzeMedals: 1,
		//AllMedals:    20,
	},
	{
		Position:     6,
		Country:      "Gran Bretaña",
		GoldMedals:   6,
		SilverMedals: 9,
		BronzeMedals: 9,
		//AllMedals:    20,
	},
}
