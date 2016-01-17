// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud_test

import (
	"github.com/theckman/go-maxmind/minfraud"
	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestEventType_MarshalJSON(c *C) {
	var et minfraud.EventType
	var j []byte
	var err error

	et = minfraud.EventTypeUnknown
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"unknown_type"`)

	et = minfraud.EventTypeAccountCreation
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"account_creation"`)

	et = minfraud.EventTypeAccountLogin
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"account_login"`)

	et = minfraud.EventTypePurchase
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"purchase"`)

	et = minfraud.EventTypeRecurringPurchase
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"recurring_purchase"`)

	et = minfraud.EventTypeReferral
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"referral"`)

	et = minfraud.EventTypeSurvey
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"survey"`)

	et++
	j, err = et.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"unknown_type"`)

}

func (t *TestSuite) TestEventType_UnmarshalJSON(c *C) {
	var et minfraud.EventType
	var data []byte
	var err error

	data = []byte(`""`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeUnknown)

	data = []byte(`"account_creation"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeAccountCreation)

	data = []byte(`"account_login"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeAccountLogin)

	data = []byte(`"purchase"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypePurchase)

	data = []byte(`"recurring_purchase"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeRecurringPurchase)

	data = []byte(`"referral"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeReferral)

	data = []byte(`"survey"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeSurvey)

	data = []byte(`"random_garbage"`)
	err = et.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(et, Equals, minfraud.EventTypeUnknown)
}
