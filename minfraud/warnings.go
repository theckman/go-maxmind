// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud

// WarningCode is a value identifying the warning as returning
// in the response of the API.
type WarningCode int

const (
	// WarnUnknown is for when the warning is an unknown type.
	WarnUnknown WarningCode = iota

	// WarnBillingCityNotFound is for when the billing city could not be found
	// in the MaxMind database.
	WarnBillingCityNotFound

	// WarnBillingCountryNotFound is for when the billing country could not be
	// found in the MaxMind database.
	WarnBillingCountryNotFound

	// WarnBillingPostalNotFound is for when the billing postal code could not
	// be found in the MaxMind database.
	WarnBillingPostalNotFound

	// WarnInputInvalid is for when the value associated with the key does
	// not meet the required constraints. An example of this is entering
	// "United States" in a field that requires a two-lette country code.
	WarnInputInvalid

	// WarnInputUnknown is for when an unknown key is encountered in ther
	// request body.
	WarnInputUnknown

	// WarnIPAddressNotFound is for when the IP address could not be
	// geolocated.
	WarnIPAddressNotFound

	// WarnShippingCityNotFound is for when the shipping city could not be
	// found in the MaxMind database.
	WarnShippingCityNotFound

	// WarnShippingCountryNotFound is for when the shipping country could not
	// be found in the MaxMind database.
	WarnShippingCountryNotFound

	// WarnShippingPostalNotFound is for when the shipping postal code could
	// not be found in the MaxMind database.
	WarnShippingPostalNotFound
)

// MarshalJSON implements the json.Marshaler interface.
func (w WarningCode) MarshalJSON() ([]byte, error) {
	switch w {
	case WarnBillingCityNotFound:
		return []byte(`"BILLING_CITY_NOT_FOUND"`), nil
	case WarnBillingCountryNotFound:
		return []byte(`"BILLING_COUNTRY_NOT_FOUND"`), nil
	case WarnBillingPostalNotFound:
		return []byte(`"BILLING_POSTAL_NOT_FOUND"`), nil
	case WarnInputInvalid:
		return []byte(`"INPUT_INVALID"`), nil
	case WarnInputUnknown:
		return []byte(`"INPUT_UNKNOWN"`), nil
	case WarnIPAddressNotFound:
		return []byte(`"IP_ADDRESS_NOT_FOUND"`), nil
	case WarnShippingCityNotFound:
		return []byte(`"SHIPPING_CITY_NOT_FOUND"`), nil
	case WarnShippingCountryNotFound:
		return []byte(`"SHIPPING_COUNTRY_NOT_FOUND"`), nil
	case WarnShippingPostalNotFound:
		return []byte(`"SHIPPING_POSTAL_NOT_FOUND"`), nil
	default:
		return []byte(`"UNKNOWN_WARNING"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (w *WarningCode) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		*w = WarnUnknown
		return
	}

	switch string(data[1 : len(data)-1]) {
	case "BILLING_CITY_NOT_FOUND":
		*w = WarnBillingCityNotFound
	case "BILLING_COUNTRY_NOT_FOUND":
		*w = WarnBillingCountryNotFound
	case "BILLING_POSTAL_NOT_FOUND":
		*w = WarnBillingPostalNotFound
	case "INPUT_INVALID":
		*w = WarnInputInvalid
	case "INPUT_UNKNOWN":
		*w = WarnInputUnknown
	case "IP_ADDRESS_NOT_FOUND":
		*w = WarnIPAddressNotFound
	case "SHIPPING_CITY_NOT_FOUND":
		*w = WarnShippingCityNotFound
	case "SHIPPING_COUNTRY_NOT_FOUND":
		*w = WarnShippingCountryNotFound
	case "SHIPPING_POSTAL_NOT_FOUND":
		*w = WarnShippingPostalNotFound
	default:
		*w = WarnUnknown
	}

	return
}
