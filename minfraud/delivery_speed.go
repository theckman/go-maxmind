package minfraud

type DeliverySpeed int

const (
	ShippingSpeedUnknown DeliverySpeed = iota
	ShippingSameDay
	ShippingOvernight
	ShippingExpedited
	ShippingStandard
)

// MarshalJSON implements the json.Marshaler interface.
func (d DeliverySpeed) MarshalJSON() ([]byte, error) {
	switch d {
	case ShippingSameDay:
		return []byte(`"same_day"`), nil
	case ShippingOvernight:
		return []byte(`"overnight"`), nil
	case ShippingExpedited:
		return []byte(`"expedited"`), nil
	case ShippingStandard:
		return []byte(`"standard"`), nil
	default:
		return []byte(`"unknown_delivery_speed"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d DeliverySpeed) UnmarshalJSON(data []byte) (err error) {
	if len(data) < 3 {
		d = ShippingSpeedUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "same_day":
			d = ShippingSameDay
		case "overnight":
			d = ShippingOvernight
		case "expedited":
			d = ShippingExpedited
		case "standard":
			d = ShippingStandard
		default:
			d = ShippingSpeedUnknown
		}
	}
	return
}
