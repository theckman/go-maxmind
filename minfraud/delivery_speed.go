// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud

// DeliverySpeed is a helper type for the shipping delivery speed enum
// field as part of the MinFraud Request Shipping type.
type DeliverySpeed int

const (
	// DeliverySpeedUnknown is for when the delivery speed is not known
	DeliverySpeedUnknown DeliverySpeed = iota

	// DeliverySpeedSameDay is for same day deliveries.
	DeliverySpeedSameDay

	// DeliverySpeedOvernight is for overnight deliveries.
	DeliverySpeedOvernight

	// DeliverySpeedExpedited is for expedited deliveries.
	DeliverySpeedExpedited

	// DeliverySpeedStandard is for standard deliveries.
	DeliverySpeedStandard
)

// MarshalJSON implements the json.Marshaler interface.
func (d DeliverySpeed) MarshalJSON() ([]byte, error) {
	switch d {
	case DeliverySpeedSameDay:
		return []byte(`"same_day"`), nil
	case DeliverySpeedOvernight:
		return []byte(`"overnight"`), nil
	case DeliverySpeedExpedited:
		return []byte(`"expedited"`), nil
	case DeliverySpeedStandard:
		return []byte(`"standard"`), nil
	default:
		return []byte(`"unknown_delivery_speed"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *DeliverySpeed) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		*d = DeliverySpeedUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "same_day":
			*d = DeliverySpeedSameDay
		case "overnight":
			*d = DeliverySpeedOvernight
		case "expedited":
			*d = DeliverySpeedExpedited
		case "standard":
			*d = DeliverySpeedStandard
		default:
			*d = DeliverySpeedUnknown
		}
	}
	return nil
}
