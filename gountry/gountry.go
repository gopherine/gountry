package gountry

import (
	"errors"
	"strings"

	"github.com/gopherine/Gountry/model"
)

//GetCountries returns list of all the countries
func GetCountries() []model.CountryType {
	return model.Country
}

//GetCountryISO2 return country data based on the ISO(2) code passed
func GetCountryISO2(ISO2 string) (*model.CountryType, error) {
	countries := model.Country
	for _, country := range countries {
		if country.ISO2 == ISO2 {
			return &country, nil
		}
	}
	return nil, errors.New(ISO2 + ",is a invalid ISO3 or not mapped - please check again.")
}

//GetCountryISO3 return country data based on the ISO(2) code passed
func GetCountryISO3(ISO3 string) (*model.CountryType, error) {
	countries := model.Country
	for _, country := range countries {
		if country.ISO3 == ISO3 {
			return &country, nil
		}
	}
	return nil, errors.New(ISO3 + ",is a invalid ISO3 or not mapped - please check again.")
}

//GetCountryNumeric return country data based on the ISO(2) code passed
func GetCountryNumeric(Numeric string) (*model.CountryType, error) {
	countries := model.Country
	for _, country := range countries {
		if country.Numeric == Numeric {
			return &country, nil
		}
	}
	return nil, errors.New(Numeric + ",is a invalid Numeric Code or not mapped - please check again.")
}

//GetSubdivision return country data based on the ISO(2) code passed
func GetSubdivision(Code string) (*model.SubdivisionsType, error) {
	countries := model.Country
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
