# Gountry [![GoDoc](https://godoc.org/github.com/gopherine/gountry?status.svg)](http://godoc.org/github.com/gopherine/gountry)  [![Build Status](https://travis-ci.org/gopherine/gountry.svg?branch=master)](https://travis-ci.org/gopherine/gountry)  [![Coverage Status](https://coveralls.io/repos/github/gopherine/gountry/badge.svg?branch=master)](https://coveralls.io/github/gopherine/gountry?branch=master)

<img src="/logo.png" alt="Gountry"
	title="A cute kitten" width="300" height="300" />

Gountry is a library to get direct mappings for country ,ISO-3166 Countries , ISO-3166-2 Subdivisions,ISO-4217 Currency.

This package maps ISO data directly to Go Structs resulting in high performance extraction. Changes to data and addition to data will be accepted as long as it is standardized.

```go
Package gountry is a data mapping to extract country details , currency and subdivisions The data is extracted and mapped on to go using python package pycountry but is open to change as per community suggestion

Example :

package main

import (

"fmt"
"github.com/gopherine/gountry"
)

func main() {

// returns list of all country data
fmt.Println(gountry.GetCountries())

//ISO2 returns one country data
fmt.Println(gountry.GetCountryISO2("NZ"))

//ISO3 returns one country data
fmt.Println(gountry.GetCountryISO3("IND"))

//Numeric returns one country data
fmt.Println(gountry.GetCountryNumeric("398"))

//Returns specific subdivision data
fmt.Println(gountry.GetSubdivision("RS-KM"))

//Refer godoc to view type definition to further get specific data for example
//gountry.GetCountryISO2("NZ").Currency
//currency
country, _ := gountry.GetCountryISO2("NZ")
fmt.Println(country.Currency)
}
```

Data is stored in data.go this will renamed as more data is added

Country : Type definition for getting countries ISO-3166 and currency ISO-4217

```go
type CountryType
type CountryType struct {
    Country      string             `json:"Country"`
    ISO2         string             `json:"ISO2"`
    ISO3         string             `json:"ISO3"`
    Numeric      string             `json:"CountryCode"`
    Currency     string             `json:"Currency"`
    Code         string             `json:"Code"`
    Subdivisions []SubdivisionsType `json:"Subdivisions"`
}
```


SubdivisionsType : Type definition for country subdivision ISO-3166-2
```go
type SubdivisionsType
type SubdivisionsType struct {
    Code       string `json:"Code"`
    ISO2       string `json:"ISO2"`
    Name       string `json:"Name"`
    ParentCode string `json:"ParentCode"`
    Type       string `json:"Type"`
}
```

Contribution:

- [X] Create an issue on github repo.
- [X] Fork the project and create a pull request.
- [X] Keep the name of your branch similar to the issue you are trying to solve.
- [X] Remember to run tests and gofmt before pull as it will fail while building via travis otherwise

TODO:

- [X] Create a local go data mapping for countries and subdivisions.
- [ ] Create a local go data mapping for cities in every country.
- [ ] Create a local go data mapping for language script.
- [ ] Improve performance.

Note :: Every data should be searchable via the main countries package i.e should work more like query tables and foreign keys.
