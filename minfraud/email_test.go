// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud_test

import (
	"github.com/theckman/go-maxmind/minfraud"
	. "gopkg.in/check.v1"
)

func (t *TestSuite) TestMD5Stringer_String(c *C) {
	var e minfraud.MD5Stringer = "TestWord"
	c.Check(e.String(), Equals, "TestWord")
}

func (t *TestSuite) TestMD5Stringer_MD5(c *C) {
	var e minfraud.MD5Stringer = "These pretzels are making me thirsty."
	c.Check(e.MD5(), Equals, "b0804ec967f48520697662a204f5fe72")
}

func (t *TestSuite) TestMD5Stringer_MarshalJSON(c *C) {
	var b []byte
	var err error
	var e minfraud.MD5Stringer = "These pretzels are making me thirsty."

	b, err = e.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(b), Equals, `"b0804ec967f48520697662a204f5fe72"`)
}
