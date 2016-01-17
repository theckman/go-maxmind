// Copyright 2015-2016 Tim Heckman. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license that can be found in the LICENSE file.

package minfraud

import (
	"crypto/md5"
	"fmt"
)

// MD5Stringer is a string that contains a function to convert its value
// to an MD5 hash of the contents. This is used by the MaxMind MinFraud
// API for the Username and Email Address fields. This allows you to give
// MaxMind the Username or Email in an obfuscated way. In the end it is
// only MD5, but slightly better than plain text.
type MD5Stringer string

func (e MD5Stringer) String() string {
	return string(e)
}

// MD5 is a function that returns the MD5 hexadecimal hash of the value.
func (e MD5Stringer) MD5() string {
	return fmt.Sprintf("%x", md5.Sum([]byte(string(e))))
}

// MarshalJSON is a function that implements the json.Marshaler interface.
func (e MD5Stringer) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf(`"%s"`, e.MD5())
	return []byte(str), nil
}
