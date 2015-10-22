package minfraud

type WarningCode int

const (
	WarnUnknown WarningCode = iota
	WarnBillingCityNotFound
	WarnBillingCountryNotFound
	WarnBillingPostalNotFound
	WarnInputInvalid
	WarnInputUnknown
	WarnIPAddressNotFound
	WarnShippingCityNotFound
	WarnShippingCountryNotFound
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
func (w WarningCode) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		w = WarnUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "BILLING_CITY_NOT_FOUND":
			w = WarnBillingCityNotFound
		case "BILLING_COUNTRY_NOT_FOUND":
			w = WarnBillingCountryNotFound
		case "BILLING_POSTAL_NOT_FOUND":
			w = WarnBillingPostalNotFound
		case "INPUT_INVALID":
			w = WarnInputInvalid
		case "INPUT_UNKNOWN":
			w = WarnInputUnknown
		case "IP_ADDRESS_NOT_FOUND":
			w = WarnIPAddressNotFound
		case "SHIPPING_CITY_NOT_FOUND":
			w = WarnShippingCityNotFound
		case "SHIPPING_COUNTRY_NOT_FOUND":
			w = WarnShippingCountryNotFound
		case "SHIPPING_POSTAL_NOT_FOUND":
			w = WarnShippingPostalNotFound
		default:
			w = WarnUnknown
		}
	}
	return
}
