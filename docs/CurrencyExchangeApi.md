# \CurrencyExchangeApi

All URIs are relative to *https://api.cloudmersive.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CurrencyExchangeConvertCurrency**](CurrencyExchangeApi.md#CurrencyExchangeConvertCurrency) | **Post** /currency/exchange-rates/convert/{source}/to/{destination} | Converts a price from the source currency into the destination currency
[**CurrencyExchangeGetAvailableCurrencies**](CurrencyExchangeApi.md#CurrencyExchangeGetAvailableCurrencies) | **Post** /currency/exchange-rates/list-available | Get a list of available currencies and corresponding countries
[**CurrencyExchangeGetExchangeRate**](CurrencyExchangeApi.md#CurrencyExchangeGetExchangeRate) | **Post** /currency/exchange-rates/get/{source}/to/{destination} | Gets the exchange rate from the source currency into the destination currency


# **CurrencyExchangeConvertCurrency**
> ConvertedCurrencyResult CurrencyExchangeConvertCurrency(ctx, source, destination, sourcePrice)
Converts a price from the source currency into the destination currency

Automatically converts the price in the source currency into the destination currency using the latest available currency exchange rate data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**| Source currency three-digit code (ISO 4217), e.g. USD, EUR, etc. | 
  **destination** | **string**| Destination currency three-digit code (ISO 4217), e.g. USD, EUR, etc. | 
  **sourcePrice** | **float64**| Input price, such as 19.99 in source currency | 

### Return type

[**ConvertedCurrencyResult**](ConvertedCurrencyResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: application/json, text/json
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyExchangeGetAvailableCurrencies**
> AvailableCurrencyResponse CurrencyExchangeGetAvailableCurrencies(ctx, )
Get a list of available currencies and corresponding countries

Enumerates available currencies and the countries that correspond to these currencies.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AvailableCurrencyResponse**](AvailableCurrencyResponse.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CurrencyExchangeGetExchangeRate**
> ExchangeRateResult CurrencyExchangeGetExchangeRate(ctx, source, destination)
Gets the exchange rate from the source currency into the destination currency

Automatically gets the exchange rate from the source currency into the destination currency using the latest available currency exchange rate data.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **source** | **string**| Source currency three-digit code (ISO 4217), e.g. USD, EUR, etc. | 
  **destination** | **string**| Destination currency three-digit code (ISO 4217), e.g. USD, EUR, etc. | 

### Return type

[**ExchangeRateResult**](ExchangeRateResult.md)

### Authorization

[Apikey](../README.md#Apikey)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/json, application/xml, text/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

