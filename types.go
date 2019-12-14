// Copyright 2019 Atharva Pandey. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gountry

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

//TimezoneType : Type definition for timezone,capital and continent
type TimezoneType struct {
	Timezones []string `json:"Timezones"`
	ISO2      string   `json:"Code"`
	Continent string   `json:"Continent"`
	Name      string   `json:"Name"`
	Capital   string   `json:"Capital"`
}
