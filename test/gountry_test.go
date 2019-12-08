package test

import "testing"

import "github.com/gopherine/Gountry/gountry"

import "github.com/gopherine/Gountry/model"

func TestGetCountries(t *testing.T) {
	if len(gountry.GetCountries()) != 249 {
		t.Errorf("Failed")
	}
}

func TestGetCountryISO2(t *testing.T) {
	countries := model.Country
	for _, country := range countries {
		_, err := gountry.GetCountryISO2(country.ISO2)
		if err != nil {
			t.Errorf("Failed")
		}
	}
}

func TestGetCountryISO3(t *testing.T) {
	countries := model.Country
	for _, country := range countries {
		_, err := gountry.GetCountryISO3(country.ISO3)
		if err != nil {
			t.Errorf("Failed")
		}
	}
}

func TestGetCountryNumeric(t *testing.T) {
	countries := model.Country
	for _, country := range countries {
		_, err := gountry.GetCountryNumeric(country.Numeric)
		if err != nil {
			t.Errorf("Failed")
		}
	}
}

func TestGetSubdivision(t *testing.T) {
	countries := model.Country
	for _, country := range countries {
		subdivisions := country.Subdivisions
		for _, subdivision := range subdivisions {
			_, err := gountry.GetSubdivision(subdivision.Code)
			if err != nil {
				t.Errorf("Failed")
			}
		}
	}
}
