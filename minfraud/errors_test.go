// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud_test

import (
	"github.com/theckman/go-maxmind/minfraud"
	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestErrCode_MarshalJSON(c *C) {
	var ec minfraud.ErrorCode
	var j []byte
	var err error

	ec = minfraud.ErrUnknown
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UNKNOWN_ERROR"`)

	ec = minfraud.ErrIPAddressInvalid
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"IP_ADDRESS_INVALID"`)

	ec = minfraud.ErrIPAddressRequired
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"IP_ADDRESS_REQUIRED"`)

	ec = minfraud.ErrIPAddressReserved
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"IP_ADDRESS_RESERVED"`)

	ec = minfraud.ErrJSONInvalid
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"JSON_INVALID"`)

	ec = minfraud.ErrAuthorizationInvalid
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AUTHORIZATION_INVALID"`)

	ec = minfraud.ErrLicenseKeyRequired
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"LICENSE_KEY_REQUIRED"`)

	ec = minfraud.ErrUserIDRequired
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"USER_ID_REQUIRED"`)

	ec = minfraud.ErrInsufficientFunds
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"INSUFFICIENT_FUNDS"`)

	ec = minfraud.ErrHTTPError
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `""`)

	ec++
	j, err = ec.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UNKNOWN_ERROR"`)
}

func (t *TestSuite) TestErrCode_UnmarshalJSON(c *C) {
	var ec minfraud.ErrorCode
	var err error

	data := []byte(`""`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrUnknown)

	data = []byte(`"IP_ADDRESS_INVALID"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrIPAddressInvalid)

	data = []byte(`"IP_ADDRESS_REQUIRED"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrIPAddressRequired)

	data = []byte(`"IP_ADDRESS_RESERVED"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrIPAddressReserved)

	data = []byte(`"JSON_INVALID"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrJSONInvalid)

	data = []byte(`"AUTHORIZATION_INVALID"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrAuthorizationInvalid)

	data = []byte(`"LICENSE_KEY_REQUIRED"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrLicenseKeyRequired)

	data = []byte(`"USER_ID_REQUIRED"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrUserIDRequired)

	data = []byte(`"INSUFFICIENT_FUNDS"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrInsufficientFunds)

	data = []byte(`"UNKNOWN_ERROR"`)
	err = ec.UnmarshalJSON(data)
	c.Assert(err, IsNil)
	c.Check(ec, Equals, minfraud.ErrUnknown)
}
