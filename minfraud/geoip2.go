// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud

import (
	"time"

	"github.com/theckman/go-maxmind/geoip2"
)

// GeoIP2Country contains details about the country where MaxMind believes the
// end user is located.
type GeoIP2Country struct {
	geoip2.Country
	IsHighRisk bool `json:"is_high_risk,omitempty"`
}

// GeoIP2Location contains specific details about the location associated with
// the IP address.
type GeoIP2Location struct {
	geoip2.Location
	LocalTime time.Time `json:"local_time,omitempty"`
}

// GeoIP2 is the minFraud-specific GeoIP2 struct. For some reason, MaxMind
// thought it was a good idea to add additional keys to some parts of the
// GeoIP2 response only for the MinFraud requests. This tries to reuse
// as many structures from the neighboring geoip2 package as possible.
type GeoIP2 struct {
	geoip2.GeoIP2
	Risk     float64         `json:"risk,omitempty"`
	Country  *GeoIP2Country  `json:"country,omitempty"`
	Location *GeoIP2Location `json:"location,omitempty"`
}
