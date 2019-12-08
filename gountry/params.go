package gountry

import "github.com/gopherine/Gountry/model"

//Country - Optional Parameter Types
func Country(c *model.CountryType) error {
	c.Country = "India"
	return nil
}
