// Package geoip2 provides types and functions to work with the MaxMind
// GeoIP2 Web Services API.
//
// As of writing only the data structures have been implemented as they are
// neded for the github.com/theckman/go-maxmind/minfraud package.
package geoip2

// Names is a struct to compose reused fields in to other structs.
type Names struct {
	// GeonameID is a unique identifier for the city as specified
	// by GeoNames: http://www.geonames.org/
	GeonameID float64 `json:"geoname_id,omitempty"`

	// Names is a map from locale codes, such as "en", to localized
	// names for the item.
	Names map[string]string `json:"names,omitempty"`
}

// City contains details about the city associated with the IP address.
type City struct {
	Names

	// Confidence is a value from 0-100 representing MaxMind's confidence
	// that the city is correct.
	Confidence float64 `json:"confidence,omitempty"`
}

// Continent contains information about the continent associated with the
// IP address.
type Continent struct {
	Names

	// Code is the ContinentCode associated with the IP address. The
	// ContinentCode type automatically converts the value to and from the
	// two-letter code used by the MaxMind API.
	Code ContinentCode `json:"code,omitempty"`
}

// Country contains details about the country where MaxMind believes the
// end user is located.
type Country struct {
	Names

	// Confidence is a value from 0-100 representing MaxMind's confidence
	// that the country is correct.
	Confidence float64 `json:"confidence,omitempty"`

	// ISOCode is a two-character ISO 3166-1 country code for the country
	// associated with the IP address.
	ISOCode string `json:"iso_code,omitempty"`
}

// Location contains specific details about the location associated with
// the IP address.
type Location struct {
	// AccuracyRadius is the radius in kilometers around the specified
	// location where the IP address is likely to be.
	AccuracyRadius float64 `json:"accuracy_radius,omitempty"`

	// AverageIncome is the average income associated with the IP address.
	AverageIncome float64 `json:"average_income,omitempty"`

	// Latitude is the latitude of the location associated with the IP address.
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude is the longitude of the location associated with the IP address.
	Longitude float64 `json:"longitude,omitempty"`

	// MetroCode is the metro code associated with the IP address. These are
	// only available for IP addresses in the US. MaxMind returns the same
	// metro codes as the Google AdWords API.
	MetroCode float64 `json:"metro_code,omitempty"`

	// PopulationDensity is the estimated number of people per square kilometer.
	PopulationDensity float64 `json:"population_density,omitempty"`

	// TimeZone is the time zone associated with location, as specified by the,
	// IANA Time Zone Database, e.g. "America/New_York".
	TimeZone string `json:"time_zone,omitempty"`
}

// Postal contains details about the postal code associated with the IP address.
type Postal struct {
	// Code is the postal code associated with the IP address. These are
	// available for some IP addresses in Australia, Canada, France, Germany,
	// Italy, Spain, Switzerland, United Kingdom, and the US. MaxMind returns
	// the first 3 characters for Canadian postal codes. We return the the
	// first 2-4 characters (outward code) for postal codes in the United Kingdom.
	Code string `json:"code,omitempty"`

	// Confidence is the value from 0-100 representing MaxMind's confidence that
	// the postal code is correct.
	Confidence float64 `json:"confidence,omitempty"`
}

// RegisteredCountry contains details about the country in which the ISP has
// registered the IP address.
type RegisteredCountry struct {
	Names

	// ISOCode is a two-character ISO 3166-1 country code for the country
	// associated with the IP address.
	ISOCode string `json:"iso_code,omitempty"`
}

// RepresentedCountry contains details about the country which is represented
// by users of the IP address. For instance, the country represented by an
// overseas military base.
type RepresentedCountry struct {
	Names

	// ISOCode is a two-character ISO 3166-1 country code for the country
	// associated with the IP address.
	ISOCode string `json:"iso_code,omitempty"`

	// Type is the type of represented country. Currently limited to military
	// but may include other types in the future.
	Type string `json:"type,omitempty"`
}

// Subdivision contains details about a subdivision of the country in which the
// IP address resides.
type Subdivision struct {
	Names
	Confidence float64 `json:"confidence,omitempty"`

	// ISOCode is a two-character ISO 3166-1 country code for the country
	// associated with the IP address.
	ISOCode string `json:"iso_code,omitempty"`
}

// Traits contains general traits associated with the IP address.
type Traits struct {
	// ASN is the Autonomous System Number associated with the IP address.
	ASN float64 `json:"autonomous_system_number,omitempty"`
	// AutonomousSystemOrganization is the organization associated with the
	// registered ASN for the IP address.
	AutonomousSystemOrganization string `json:"autonomous_system_organization,omitempty"`

	// Domain is the second-level domain associated with the IP address,
	// e.g. example.com
	Domain string `json:"domain,omitempty"`

	IPAddress    string `json:"ip_address,omitempty"`
	ISP          string `json:"isp,omitempty"`
	Organization string `json:"organization,omitempty"`

	// UserType is the user type associated with the IP address. See the
	// MaxMind API documentation for more information on possible types.
	UserType string `json:"user_type,omitempty"`
}

// MaxMind contains information related to your MaxMind account.
type MaxMind struct {
	// QueriesRemaining is the approximate number of remaining queries
	// available for the end point which is being called.
	QueriesRemaining float64 `json:"queries_remaining,omitempty"`
}

// GeoIP2 is a struct representing the full response body of the GeoIP2 Web
// Services API.
type GeoIP2 struct {
	City               *City               `json:"city,omitempty"`
	Continent          *Continent          `json:"continent,omitempty"`
	Country            *Country            `json:"country,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Postal             *Postal             `json:"postal,omitempty"`
	RegisteredCountry  *RegisteredCountry  `json:"registered_country,omitempty"`
	RepresentedCountry *RepresentedCountry `json:"represented_country,omitempty"`

	// Subdivisions is a slice of *Subdivision. Each contains details about a
	// subdivision of the country in which the IP address resides. Subdivisions
	// are arranged from largest to smallest.
	//
	// For instance, the response for Oxford in the United Kingdom would have an
	// object for England as the first element in subdivisions array and an
	// object for Oxfordshire as the second element. The subdivisions array for
	// Minneapolis in the United States will have a single object for Minnesota.
	Subdivisions []*Subdivision `json:"subdivisions,omitempty"`

	Traits  *Traits  `json:"traits,omitempty"`
	MaxMind *MaxMind `json:"maxmind,omitempty"`
}

// ContinentCode is a type to easily work with the different continent codes
// used by the MaxMind GeoIP2 Web Services API. Thie type has a Marshal and
// Unmarshal JSON function which will automatically convert it to and from the
// values used by the API.
//
// For example, if the MaxMind API responds with "EU" we convert that to
// geoip2.ContinentEurope.
type ContinentCode int

const (
	// ContinentUnknown is for unknown values
	ContinentUnknown ContinentCode = iota

	// ContinentAfrica is for Africa.
	ContinentAfrica

	// ContinentAntarctica is for Antarctica.
	ContinentAntarctica

	// ContinentAsia is for Asia.
	ContinentAsia

	// ContinentEurope is for Europe.
	ContinentEurope

	// ContinentNorthAmerica is for North America (and Moose).
	ContinentNorthAmerica

	// ContinentOceania is for Oceania.
	ContinentOceania

	// ContinentSouthAmerica is for South America.
	ContinentSouthAmerica
)

// MarshalJSON implements the json.Marshaler interface.
func (g ContinentCode) MarshalJSON() ([]byte, error) {
	switch g {
	case ContinentAfrica:
		return []byte(`"AF"`), nil
	case ContinentAntarctica:
		return []byte(`"AN"`), nil
	case ContinentAsia:
		return []byte(`"AS"`), nil
	case ContinentEurope:
		return []byte(`"EU"`), nil
	case ContinentNorthAmerica:
		return []byte(`"NA"`), nil
	case ContinentOceania:
		return []byte(`"OC"`), nil
	case ContinentSouthAmerica:
		return []byte(`"SA"`), nil
	default:
		return []byte(`"UN"`), nil
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (g ContinentCode) UnmarshalJSON(data []byte) error {
	if len(data) < 3 {
		g = ContinentUnknown
	} else {
		switch string(data[1 : len(data)-1]) {
		case "AF":
			g = ContinentAfrica
		case "AN":
			g = ContinentAntarctica
		case "AS":
			g = ContinentAsia
		case "EU":
			g = ContinentEurope
		case "NA":
			g = ContinentNorthAmerica
		case "OC":
			g = ContinentOceania
		case "SA":
			g = ContinentSouthAmerica
		default:
			g = ContinentUnknown
		}
	}

	return nil
}
