// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud_test

import (
	"github.com/theckman/go-maxmind/minfraud"
	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestDeliverySpeed_MarshalJSON(c *C) {
	var ds minfraud.DeliverySpeed
	var j []byte
	var err error

	ds = minfraud.DeliverySpeedUnknown
	j, err = ds.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"unknown_delivery_speed"`)

	ds = minfraud.DeliverySpeedSameDay
	j, err = ds.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"same_day"`)

	ds = minfraud.DeliverySpeedOvernight
	j, err = ds.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"overnight"`)

	ds = minfraud.DeliverySpeedExpedited
	j, err = ds.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"expedited"`)

	ds = minfraud.DeliverySpeedStandard
	j, err = ds.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"standard"`)

	ds++
	j, err = ds.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"unknown_delivery_speed"`)
}

func (t *TestSuite) TestDeliverySpeed_UnmarshalJSON(c *C) {
	var ds minfraud.DeliverySpeed
	var err error

	data := []byte(`""`)
	err = ds.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ds, Equals, minfraud.DeliverySpeedUnknown)

	data = []byte(`"same_day"`)
	err = ds.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ds, Equals, minfraud.DeliverySpeedSameDay)

	data = []byte(`"overnight"`)
	err = ds.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ds, Equals, minfraud.DeliverySpeedOvernight)

	data = []byte(`"expedited"`)
	err = ds.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ds, Equals, minfraud.DeliverySpeedExpedited)

	data = []byte(`"standard"`)
	err = ds.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ds, Equals, minfraud.DeliverySpeedStandard)

	data = []byte(`"unknown_delivery_soeed"`)
	err = ds.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ds, Equals, minfraud.DeliverySpeedUnknown)
}
