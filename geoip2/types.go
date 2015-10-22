package geoip2

type Names struct {
	GeonameID float64           `json:"geoname_id,omitempty"`
	Names     map[string]string `json:"names,omitempty"`
}

type City struct {
	Names
	Confidence float64 `json:"confidence,omitempty"`
}

type Continent struct {
	Names
	Code GeoIP2ContinentCode `json:"code,omitempty"`
}

type Country struct {
	Names
	Confidence float64 `json:"confidence,omitempty"`
	ISOCode    string  `json:"iso_code,omitempty"` // ISO 3166-1
}

type Location struct {
	AccuracyRadius    float64 `json:"accuracy_radius,omitempty"`
	AverageIncome     float64 `json:"average_income,omitempty"`
	Latitude          float64 `json:"latitude,omitempty"`
	Longitude         float64 `json:"longitude,omitempty"`
	MetroCode         float64 `json:"metro_code,omitempty"`
	PopulationDensity float64 `json:"population_density,omitempty"`
	TimeZone          string  `json:"time_zone,omitempty"`
}

type Postal struct {
	Code       string  `json:"code,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

type RegisteredCountry struct {
	Names
	ISOCode string `json:"iso_code,omitempty"` // ISO 3166-1
}

type RepresentedCountry struct {
	Names
	ISOCode string `json:"iso_code,omitempty"` // ISO 3166-1
	Type    string `json:"type,omitempty"`
}

type Subdivisions struct {
	Names
	Confidence float64 `json:"confidence,omitempty"`
	ISOCode    string  `json:"iso_code,omitempty"` // ISO 3166-1
}

type Traits struct {
	ASN                          float64 `json:"autonomous_system_number,omitempty"`
	AutonomousSystemOrganization string  `json:"autonomous_system_organization,omitempty"`
	Domain                       string  `json:"domain,omitempty"`
	IPAddress                    string  `json:"ip_address,omitempty"`
	ISP                          string  `json:"isp,omitempty"`
	Organization                 string  `json:"organization,omitempty"`
	UserType                     string  `json:"user_type,omitempty"`
}

type MaxMind struct {
	QueriesRemaining float64 `json:"queries_remaining,omitempty"`
}

type GeoIP2 struct {
	City               *City               `json:"city,omitempty"`
	Continent          *Continent          `json:"continent,omitempty"`
	Country            *Country            `json:"country,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Postal             *Postal             `json:"postal,omitempty"`
	RegisteredCountry  *RegisteredCountry  `json:"registered_country,omitempty"`
	RepresentedCountry *RepresentedCountry `json:"represented_country,omitempty"`
	Subdivisions       *Subdivisions       `json:"subdivisions,omitempty"`
	Traits             *Traits             `json:"traits,omitempty"`
	MaxMind            *MaxMind            `json:"maxmind,omitempty"`
}

type ContinentCode int

const (
	ContinentUnknown ContinentCode = iota
	ContinentAfrica
	ContinentAntarctica
	ContinentAsia
	ContinentEurope
	ContinentNorthAmerica
	ContinentOceania
	ContinentSouthAmerica
)

// MarshalJSON implements the json.Marshaler interface.
func (g ContinentCode) MarshalJSON() ([]byte, error) {
	switch g {
	case ContinentAfrica:
		return []byte("AF"), nil
	case ContinentAntarctica:
		return []byte("AN"), nil
	case ContinentAsia:
		return []byte("AS"), nil
	case ContinentEurope:
		return []byte("EU"), nil
	case ContinentNorthAmerica:
		return []byte("NA"), nil
	case ContinentOceania:
		return []byte("OC"), nil
	case ContinentSouthAmerica:
		return []byte("SA"), nil
	default:
		return []byte("UN"), nil
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
}
