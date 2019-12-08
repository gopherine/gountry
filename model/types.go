package model

//SubdivisionsType : Type definition for country subdivision which includes terit
type SubdivisionsType struct {
	Code       string `json:"Code"`
	ISO2       string `json:"ISO2"`
	Name       string `json:"Name"`
	ParentCode string `json:"ParentCode"`
	Type       string `json:"Type"`
}

//CountryType : Type definition for country
type CountryType struct {
	Country      string             `json:"Country"`
	ISO2         string             `json:"ISO2"`
	ISO3         string             `json:"ISO3"`
	Numeric      string             `json:"CountryCode"`
	Currency     string             `json:"Currency"`
	Code         string             `json:"Code"`
	Subdivisions []SubdivisionsType `json:"Subdivisions"`
}
