package minfraud

import "time"

// Device contains information about the device used in the transaction.
type Device struct {
	IPAddress      string `json:"ip_address"`
	UserAgent      string `json:"user_agent,omitempty"`
	AcceptLanguage string `json:"accept_lagnuage,omitempty"`
}

// Event contains general information related to the event being scored.
type Event struct {
	TransactionID string    `json:"transaction_id,omitempty"`
	ShopID        string    `json:"shop_id,omitempty"`
	Time          time.Time `json:"time,omitempty"`
	Type          EventType `json:"type,omitempty"`
}

// Account contains account information for the end-user on the site where the event took place.
type Account struct {
	UserID      string      `json:"user_id,omitempty"`
	UsernamdMD5 MD5Stringer `json:"username_md5,omitempty"`
}

// Email contains information about an email address to include within the lookup.
type Email struct {
	AddressMD5 MD5Stringer `json:"address,omitempty"`
	Domain     string      `json:"domain,omitempty"`
}

// ContactDetails is a struct used for compositing. Both the Shipping and Billing types
// need these same fields.
type ContactDetails struct {
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	Company          string `json:"company,omitempty"`
	Address          string `json:"address,omitempty"`
	Address2         string `json:"address_2,omitempty"`
	City             string `json:"city,omitempty"`
	Region           string `json:"region,omitempty"`  // ISO 3166-2
	Country          string `json:"country,omitempty"` // ISO 3166-1 (alpha-2)
	Postal           string `json:"postal,omitempty"`
	PhoneNumber      string `json:"phone_number,omitempty"`
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
}

// Billing contains information about the billing contact.
// It's composed of the ContactDetails struct.
type Billing struct {
	ContactDetails
}

// Shipping contains information about the shipping contact.
// It's composed of the ContactDetails struct as well as the delivery
// speed.
type Shipping struct {
	ContactDetails
	DeliverySpeed DeliverySpeed `json:"delivery_speed,omitempty"`
}

// Payment is something that is fuuuucked
type Payment struct {
	Processor     string `json:"processor,omitempty"`
	WasAuthorized bool   `json:"was_authorized,omitempty"`
	DeclineCode   string `json:"decline_code,omitempty"`
}

// CreditCardRequest contains information about the credit card used.
type CreditCardRequest struct {
	IssuerIDNumber       string `json:"issuer_id_number,omitempty"`
	Last4Digits          string `json:"last_4_digits,omitempty"`
	BankName             string `json:"bank_name,omitempty"`
	BankPhoneCountryCode string `json:"bank_phone_country_code,omitempty"`
	BankPhoneNumber      string `json:"bank_phone_number,omitempty"`
	AVSResult            string `json:"avs_result,omitempty"`
	CVVResult            string `json:"cvv_result,omitempty"`
}

// NewCreditCardRequest takes a full credit card number, and then returns
// a *CreditCardRequest with both the IssuerIDNumber and Last4Digits fields
// populated.
func NewCreditCardRequest(cc string) *CreditCardRequest {
	return &CreditCardRequest{
		IssuerIDNumber: cc[:6],
		Last4Digits:    cc[len(cc)-4:],
	}
}

// Order contains information about the order made.
type Order struct {
	Amount         float64 `json:"amonut,omitempty"`
	Currency       string  `json:"currency,omitempty"` // ISO 4217
	DiscountCode   string  `json:"discount_code,omitempty"`
	AffiliateID    string  `json:"affiliate_id,omitempty"`
	SubaffiliateID string  `json:"subaffiliate_id,omitempty"`
	ReferrerID     string  `json:"referrer_uri,omitempty"`
	IsGift         bool    `json:"is_gift,omitempty"`
	HasGiftMessage bool    `json:"has_gift_message,omitempty"`
}

// ShoppingCartItem contains information about items in the shopping cart.
type ShoppingCartItem struct {
	Category string  `json:"category,omitempty"`
	ItemID   string  `json:"item_id,omitempty"`
	Quantity float64 `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

// Request contains the information needed to build a request body for the
// minFraud Score and Insights APIs.
type Request struct {
	Device       *Device             `json:"device,omitempty"`
	Event        *Event              `json:"event,omitempty"`
	Account      *Account            `json:"account,omitempty"`
	Email        *Email              `json:"email,omitempty"`
	Billing      *Billing            `json:"billing,omitempty"`
	Shipping     *Shipping           `json:"shipping,omitempty"`
	Payment      *Payment            `json:"payment,omitempty"`
	Order        *Order              `json:"order,omitempty"`
	ShoppingCart []*ShoppingCartItem `json:"shopping_cart,omitempty"`
}
