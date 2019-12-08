// Copyright 2019 Atharva Pandey. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package gountry

import "testing"

func TestGetCountries(t *testing.T) {
	if len(GetCountries()) != 249 {
		t.Errorf("Failed")
	}
}

func TestGetCountryISO2(t *testing.T) {
	countries := Country
	for _, country := range countries {
		_, err := GetCountryISO2(country.ISO2)
		if err != nil {
			t.Errorf("Failed")
		}
	}
}

func TestGetCountryISO3(t *testing.T) {
	countries := Country
	for _, country := range countries {
		_, err := GetCountryISO3(country.ISO3)
		if err != nil {
			t.Errorf("Failed")
		}
	}
}

func TestGetCountryNumeric(t *testing.T) {
	countries := Country
	for _, country := range countries {
		_, err := GetCountryNumeric(country.Numeric)
		if err != nil {
			t.Errorf("Failed")
		}
	}
}

func TestGetSubdivision(t *testing.T) {
	countries := Country
	for _, country := range countries {
		subdivisions := country.Subdivisions
		for _, subdivision := range subdivisions {
			_, err := GetSubdivision(subdivision.Code)
			if err != nil {
				t.Errorf("Failed")
			}
		}
	}
}
