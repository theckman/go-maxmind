// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package geoip2_test

import (
	"github.com/theckman/go-maxmind/geoip2"

	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestContinentCode_MarshalJSON(c *C) {
	var cc geoip2.ContinentCode
	var j []byte
	var err error

	cc = geoip2.ContinentUnknown
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UN"`)

	cc = geoip2.ContinentAfrica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AF"`)

	cc = geoip2.ContinentAntarctica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AN"`)

	cc = geoip2.ContinentAsia
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AS"`)

	cc = geoip2.ContinentEurope
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"EU"`)

	cc = geoip2.ContinentNorthAmerica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"NA"`)

	cc = geoip2.ContinentOceania
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"OC"`)

	cc = geoip2.ContinentSouthAmerica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"SA"`)

	cc++
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UN"`)
}

func (t *TestSuite) TestContinentCode_UnmarshalJSON(c *C) {
	var cc geoip2.ContinentCode

	data := []byte(`""`)
	err := cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentUnknown)

	data = []byte(`"AF"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentAfrica)

	data = []byte(`"AN"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentAntarctica)

	data = []byte(`"AS"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentAsia)

	data = []byte(`"EU"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentEurope)

	data = []byte(`"NA"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentNorthAmerica)

	data = []byte(`"OC"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentOceania)

	data = []byte(`"SA"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentSouthAmerica)

	data = []byte(`"UN"`)
	err = cc.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(cc, Equals, geoip2.ContinentUnknown)
}
