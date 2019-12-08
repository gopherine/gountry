package main

import "github.com/gopherine/Gountry/gountry"

import "log"

func main() {
	country, _ := gountry.GetSubdivision("IN-HR")
	log.Println(country)
}
