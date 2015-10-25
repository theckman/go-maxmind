package minfraud

type EventType int

const (
	EventTypeUnknown EventType = iota
	EventTypeAccountCreation
	EventTypeAccountLogin
	EventTypePurchase
	EventTypeRecurringPurchase
	EventTypeReferral
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
func (e EventType) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		e = EventTypeUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "account_creation":
			e = EventTypeAccountCreation
		case "account_login":
			e = EventTypeAccountLogin
		case "purchase":
			e = EventTypePurchase
		case "recurring_purchase":
			e = EventTypeRecurringPurchase
		case "referral":
			e = EventTypeReferral
		case "survey":
			e = EventTypeSurvey
		default:
			e = EventTypeUnknown
		}
	}
	return
}
