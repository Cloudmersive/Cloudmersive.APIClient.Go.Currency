/*
 * currencyapi
 *
 * The currency APIs help you retrieve exchange rates and convert prices between currencies easily.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package Go-CloudmersiveCurrencyApiClient

// Result of performing an enumerate currencies operation
type AvailableCurrencyResponse struct {
	// List of available currencies from the API
	Currencies []AvailableCurrency `json:"Currencies,omitempty"`
}
