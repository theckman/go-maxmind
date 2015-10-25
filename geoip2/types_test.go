package geoip2

import (
	"testing"

	. "gopkg.in/check.v1"
)

type TestSuite struct{}

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (t *TestSuite) TestContinentCode_MarshalJSON(c *C) {
	var cc ContinentCode
	var j []byte
	var err error

	cc = ContinentUnknown
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UN"`)

	cc = ContinentAfrica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AF"`)

	cc = ContinentAntarctica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AN"`)

	cc = ContinentAsia
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"AS"`)

	cc = ContinentEurope
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"EU"`)

	cc = ContinentNorthAmerica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"NA"`)

	cc = ContinentOceania
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"OC"`)

	cc = ContinentSouthAmerica
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"SA"`)

	cc++
	j, err = cc.MarshalJSON()
	c.Assert(err, IsNil)
	c.Check(string(j), Equals, `"UN"`)
}
