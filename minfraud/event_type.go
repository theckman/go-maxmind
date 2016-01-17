// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud

// EventType is the type of event being scored. The constants
// with the EventType prefix are the individual types. They
// should be self-explanatory.
type EventType int

const (
	// EventTypeUnknown for when the type is unknown.
	EventTypeUnknown EventType = iota

	// EventTypeAccountCreation for when the type is account creation.
	EventTypeAccountCreation

	// EventTypeAccountLogin for when the type is an account login.
	EventTypeAccountLogin

	// EventTypePurchase for when the type is a purchase.
	EventTypePurchase

	// EventTypeRecurringPurchase for when the type is a recurring purchase.
	EventTypeRecurringPurchase

	// EventTypeReferral for when the type is a referral.
	EventTypeReferral

	// EventTypeSurvey is for when the type is a survey.
	EventTypeSurvey
)

// MarshalJSON implements the json.Marshaler interface.
func (e EventType) MarshalJSON() ([]byte, error) {
	switch e {
	case EventTypeAccountCreation:
		return []byte(`"account_creation"`), nil
	case EventTypeAccountLogin:
		return []byte(`"account_login"`), nil
	case EventTypePurchase:
		return []byte(`"purchase"`), nil
	case EventTypeRecurringPurchase:
		return []byte(`"recurring_purchase"`), nil
	case EventTypeReferral:
		return []byte(`"referral"`), nil
	case EventTypeSurvey:
		return []byte(`"survey"`), nil
	default:
		return []byte(`"unknown_type"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (e *EventType) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		*e = EventTypeUnknown
		return
	}

	switch string(data[1 : len(data)-1]) {
	case "account_creation":
		*e = EventTypeAccountCreation
	case "account_login":
		*e = EventTypeAccountLogin
	case "purchase":
		*e = EventTypePurchase
	case "recurring_purchase":
		*e = EventTypeRecurringPurchase
	case "referral":
		*e = EventTypeReferral
	case "survey":
		*e = EventTypeSurvey
	default:
		*e = EventTypeUnknown
	}

	return
}
