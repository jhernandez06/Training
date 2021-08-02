package models

type CountryMedal struct {
	Position     int
	Country      string
	GoldMedals   int64
	SilverMedals int64
	BronzeMedals int64
	AllMedals    int64
}
