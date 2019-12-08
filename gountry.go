// Copyright 2019 Atharva Pandey. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

//Package gountry is a data mapping to extract country details , currency and subdivisions
//The data is extracted and mapped on to go using python package pycountry but is open to change
//as per community suggestion
//
//Example :
//
// package main
//
// import (
// 	"fmt"
// 	"github.com/gopherine/gountry"
// )
//
// func main() {
// 	// returns list of all country data
// 	fmt.Println(gountry.GetCountries())
//
// 	//ISO2 returns one country data
// 	fmt.Println(gountry.GetCountryISO2("NZ"))
//
// 	//ISO3 returns one country data
// 	fmt.Println(gountry.GetCountryISO3("IND"))
//
// 	//Numeric returns one country data
// 	fmt.Println(gountry.GetCountryNumeric("398"))
//
// 	//Returns specific subdivision data
// 	fmt.Println(gountry.GetSubdivision("RS-KM"))
//
// 	//Refer godoc to view type definition to further get specific data for example
// 	//gountry.GetCountryISO2("NZ").Currency
// 	//currency
// 	country, _ := gountry.GetCountryISO2("NZ")
// 	fmt.Println(country.Currency)
//
// }
package gountry

import (
	"errors"
	"strings"
)

//GetCountries returns list of all the countries
func GetCountries() []CountryType {
	return Country
}

//GetCountryISO2 return country data based on the ISO(2) code passed
func GetCountryISO2(ISO2 string) (*CountryType, error) {
	countries := Country
	for _, country := range countries {
		if country.ISO2 == ISO2 {
			return &country, nil
		}
	}
	return nil, errors.New(ISO2 + ",is a invalid ISO3 or not mapped - please check again.")
}

//GetCountryISO3 return country data based on the ISO(2) code passed
func GetCountryISO3(ISO3 string) (*CountryType, error) {
	countries := Country
	for _, country := range countries {
		if country.ISO3 == ISO3 {
			return &country, nil
		}
	}
	return nil, errors.New(ISO3 + ",is a invalid ISO3 or not mapped - please check again.")
}

//GetCountryNumeric return country data based on the ISO(2) code passed
func GetCountryNumeric(Numeric string) (*CountryType, error) {
	countries := Country
	for _, country := range countries {
		if country.Numeric == Numeric {
			return &country, nil
		}
	}
	return nil, errors.New(Numeric + ",is a invalid Numeric Code or not mapped - please check again.")
}

//GetSubdivision return country data based on the ISO(2) code passed
func GetSubdivision(Code string) (*SubdivisionsType, error) {
	countries := Country
	ISO2 := strings.Split(Code, "-")[0]
	for _, country := range countries {
		if country.ISO2 == ISO2 {
			subdivisions := country.Subdivisions
			for _, subdivision := range subdivisions {
				if subdivision.Code == Code {
					return &subdivision, nil
				}
			}
		}
	}
	return nil, errors.New(Code + ",is a invalid Subdivision Code or not mapped - please check again.")
}
