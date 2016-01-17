// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud

// ErrorCode is a type that allows us to specify a const in our structs
// that can then be translated to the API representationw when we marshal
// the value to JSON.
type ErrorCode int

const (
	// ErrUnknown is for unknown errors
	ErrUnknown ErrorCode = iota

	// ErrIPAddressInvalid is for when you have not supplied a valid
	// IPv4 or IPv6 address.
	ErrIPAddressInvalid

	// ErrIPAddressRequired is for when you have not supplied an IP
	// address, which is a required field.
	ErrIPAddressRequired

	// ErrIPAddressReserved is for when you have supplied an IP address
	// which is reserved.
	ErrIPAddressReserved

	// ErrJSONInvalid is for when MaxMind cannot decode the body as a
	// JSON object.
	ErrJSONInvalid

	// ErrAuthorizationInvalid is for when ou have supplied an invalid
	// MaxMind user ID and/or license key in the Authorization header.
	ErrAuthorizationInvalid

	// ErrLicenseKeyRequired is for when you have not supplied a MaxMind
	// license key in the Authorization header.
	ErrLicenseKeyRequired

	// ErrUserIDRequired is for when you have not supplied a MaxMind
	// user ID in the Authorization header.
	ErrUserIDRequired

	// ErrInsufficientFunds is for when the license key you have provided
	// does not have sufficient funds to use this service. Please purchase
	// more service credits.
	ErrInsufficientFunds

	// ErrHTTPError is for when there is an HTTP related error. This includes
	// 5XX level errors sent by MaxMind as well as an HTTP 415:
	//
	// Your request included a Content-Type header that is not supported. For
	// GET requests, this means the web service cannot return content of that
	// type. For PUT and POST queries, this means the web service cannot parse
	// a request body of that type.
	ErrHTTPError
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
	case ErrHTTPError:
		return []byte(`""`), nil
	default:
		return []byte(`"UNKNOWN_ERROR"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *ErrorCode) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		*e = ErrUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "IP_ADDRESS_INVALID":
			*e = ErrIPAddressInvalid
		case "IP_ADDRESS_REQUIRED":
			*e = ErrIPAddressRequired
		case "IP_ADDRESS_RESERVED":
			*e = ErrIPAddressReserved
		case "JSON_INVALID":
			*e = ErrJSONInvalid
		case "AUTHORIZATION_INVALID":
			*e = ErrAuthorizationInvalid
		case "LICENSE_KEY_REQUIRED":
			*e = ErrLicenseKeyRequired
		case "USER_ID_REQUIRED":
			*e = ErrUserIDRequired
		case "INSUFFICIENT_FUNDS":
			*e = ErrInsufficientFunds
		default:
			*e = ErrUnknown
		}
	}
	return
}
