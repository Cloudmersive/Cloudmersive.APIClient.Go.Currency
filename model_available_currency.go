/*
 * currencyapi
 *
 * The currency APIs help you retrieve exchange rates and convert prices between currencies easily.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveCurrencyApiClient

// Details of a specific currency
type AvailableCurrency struct {
	// ISO 4217 currency three-letter code associated with the country
	ISOCurrencyCode string `json:"ISOCurrencyCode,omitempty"`
	// Symbol associated with the currency
	CurrencySymbol string `json:"CurrencySymbol,omitempty"`
	// Full name of the currency
	CurrencyEnglishName string `json:"CurrencyEnglishName,omitempty"`
	// Name of the country
	CountryName string `json:"CountryName,omitempty"`
	// Three-letter ISO 3166-1 country code
	CountryThreeLetterCode string `json:"CountryThreeLetterCode,omitempty"`
	// Two-letter ISO 3166-1 country code
	CountryISOTwoLetterCode string `json:"CountryISOTwoLetterCode,omitempty"`
	// True if this country is currently a member of the European Union (EU), false otherwise
	IsEuropeanUnionMember bool `json:"IsEuropeanUnionMember,omitempty"`
}