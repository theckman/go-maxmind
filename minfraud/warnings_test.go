// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud_test

import (
	"github.com/theckman/go-maxmind/minfraud"
	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestWarnings_MarshalJSON(c *C) {
	var wc minfraud.WarningCode
	var j []byte
	var err error

	wc = minfraud.WarnUnknown
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UNKNOWN_WARNING"`)

	wc = minfraud.WarnBillingCountryNotFound
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"BILLING_COUNTRY_NOT_FOUND"`)

	wc = minfraud.WarnBillingPostalNotFound
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"BILLING_POSTAL_NOT_FOUND"`)

	wc = minfraud.WarnInputInvalid
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"INPUT_INVALID"`)

	wc = minfraud.WarnInputUnknown
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"INPUT_UNKNOWN"`)

	wc = minfraud.WarnIPAddressNotFound
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"IP_ADDRESS_NOT_FOUND"`)

	wc = minfraud.WarnShippingCityNotFound
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"SHIPPING_CITY_NOT_FOUND"`)

	wc = minfraud.WarnShippingCountryNotFound
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"SHIPPING_COUNTRY_NOT_FOUND"`)

	wc = minfraud.WarnShippingPostalNotFound
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"SHIPPING_POSTAL_NOT_FOUND"`)

	wc++
	j, err = wc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UNKNOWN_WARNING"`)
}

func (t *TestSuite) TestWarnings_UnmarshalJSON(c *C) {
	var wc minfraud.WarningCode
	var data []byte
	var err error

	data = []byte(`""`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnUnknown)

	data = []byte(`"BILLING_COUNTRY_NOT_FOUND"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnBillingCountryNotFound)

	data = []byte(`"BILLING_POSTAL_NOT_FOUND"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnBillingPostalNotFound)

	data = []byte(`"INPUT_INVALID"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnInputInvalid)

	data = []byte(`"INPUT_UNKNOWN"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnInputUnknown)

	data = []byte(`"IP_ADDRESS_NOT_FOUND"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnIPAddressNotFound)

	data = []byte(`"SHIPPING_CITY_NOT_FOUND"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnShippingCityNotFound)

	data = []byte(`"SHIPPING_COUNTRY_NOT_FOUND"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnShippingCountryNotFound)

	data = []byte(`"SHIPPING_POSTAL_NOT_FOUND"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnShippingPostalNotFound)

	data = []byte(`"RANDOM_DATA"`)
	err = wc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(wc, Equals, minfraud.WarnUnknown)
}
