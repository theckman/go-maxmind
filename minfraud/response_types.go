package minfraud

type Warning struct {
	Code    WarningCode `json:"code,omitempty"`
	Warning string      `json:"warning,omitempty"`
	Input   []string    `json:"input,omitempty"`
}

type Score struct {
	ID               string     `json:"id,omitempty"`
	RiskScore        float64    `json:"risk_score,omitempty"`
	CreditsRemaining float64    `json:"credits_remaining,omitempty"`
	Warnings         []*Warning `json:"warnings,omitempty"`
}

type Issuer struct {
	Name                       string `json:"name"`
	MatchesProvidedName        bool   `json:"matches_provided_name"`
	PhoneNumber                string `json:"phone_number"`
	MatchesProvidedPhoneNumber bool   `json:"matches_provided_phone_number"`
}

type CreditCardResponse struct {
	Issuer                          *Issuer `json:"issuer"`
	Country                         string  `json:"country"`
	IsIssuedInBillingAddressCountry bool    `json:"is_issued_in_billing_address_country"` // wowza
	IsPrepaid                       bool    `json"is_prepaid"`
}

type ShippingAddress struct {
	IsHighRisk              bool    `json:"is_high_risk"`
	IsPostalInCity          bool    `json:"is_postal_in_city"`
	Latitude                float64 `json:"latitude"`
	Longitude               float64 `json:"longitude"`
	DistanceToIPLocation    float64 `json:"distance_to_ip_location"`
	DistanceToBillinAddress float64 `json:"distance_to_billing_address"`
	IsInIPCountry           bool    `json:"is_in_ip_country"`
}

type BillingAddress struct {
	IsPostalInCity       bool    `json:"is_postal_in_city"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	DistanceToIPLocation float64 `json:"distance_to_ip_location"`
	IsInIPCountry        bool    `json:"is_in_ip_country"`
}

type Insights struct {
	MinFraudScore
	IPAddress       *GeoIP2             `json:"ip_address,omitempty"`
	CreditCard      *CreditCardResponse `json:"credit_card,omitempty"`
	ShippingAddress *ShippingAddress    `json:"shipping_address,omitempty"`
	BillingAddress  *BillingAddress     `json:"billing_address,omitempty"`
}

type ErrorResponse struct {
	Code  ErrorCode `json:"code"`
	Error string    `json:"error"`
}
