package minfraud

type ErrorCode int

const (
	ErrUnknown ErrorCode = iota
	ErrIPAddressInvalid
	ErrIPAddressRequired
	ErrIPAddressReserved
	ErrJSONInvalid
	ErrAuthorizationInvalid
	ErrLicenseKeyRequired
	ErrUserIDRequired
	ErrInsufficientFunds
)

// MarshalJSON implements the json.Marshaler interface.
func (e ErrorCode) MarshalJSON() ([]byte, error) {
	switch e {
	case ErrIPAddressInvalid:
		return []byte(`"IP_ADDRESS_INVALID"`), nil
	case ErrIPAddressRequired:
		return []byte(`"IP_ADDRESS_REQUIRED"`), nil
	case ErrIPAddressReserved:
		return []byte(`"IP_ADDRESS_RESERVED"`), nil
	case ErrJSONInvalid:
		return []byte(`"JSON_INVALID"`), nil
	case ErrAuthorizationInvalid:
		return []byte(`"AUTHORIZATION_INVALID"`), nil
	case ErrLicenseKeyRequired:
		return []byte(`"LICENSE_KEY_REQUIRED"`), nil
	case ErrUserIDRequired:
		return []byte(`"USER_ID_REQUIRED"`), nil
	case ErrInsufficientFunds:
		return []byte(`"INSUFFICIENT_FUNDS"`), nil
	default:
		return []byte(`"UNKNOWN_ERROR"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e ErrorCode) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		e = ErrUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "IP_ADDRESS_INVALID":
			e = ErrIPAddressInvalid
		case "IP_ADDRESS_REQUIRED":
			e = ErrIPAddressRequired
		case "IP_ADDRESS_RESERVED":
			e = ErrIPAddressReserved
		case "JSON_INVALID":
			e = ErrJSONInvalid
		case "AUTHORIZATION_INVALID":
			e = ErrAuthorizationInvalid
		case "LICENSE_KEY_REQUIRED":
			e = ErrLicenseKeyRequired
		case "USER_ID_REQUIRED":
			e = ErrUserIDRequired
		case "INSUFFICIENT_FUNDS":
			e = ErrInsufficientFunds
		default:
			e = ErrUnknown
		}
	}
	return
}
