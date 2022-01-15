// Copyright (C) 2013-2018 by Maxim Bublis <b@codemonkey.ru>
// Copyright (C) 2022 by Andrew Thornton <art27@cantab.net>
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package uuid

import (
	"github.com/google/uuid"
)

var (
	global = newGoogleUUIDGenerator()
)

// NewV1 returns UUID based on current timestamp and MAC address.
func NewV1() (UUID, error) {
	return global.NewV1()
}

// NewV2 returns DCE Security UUID based on POSIX UID/GID.
func NewV2(domain byte) (UUID, error) {
	return global.NewV2(domain)
}

// NewV3 returns UUID based on MD5 hash of namespace UUID and name.
func NewV3(ns UUID, name string) UUID {
	return global.NewV3(ns, name)
}

// NewV4 returns random generated UUID.
func NewV4() (UUID, error) {
	return global.NewV4()
}

// NewV5 returns UUID based on SHA-1 hash of namespace UUID and name.
func NewV5(ns UUID, name string) UUID {
	return global.NewV5(ns, name)
}

// Generator provides interface for generating UUIDs.
type Generator interface {
	NewV1() (UUID, error)
	NewV2(domain byte) (UUID, error)
	NewV3(ns UUID, name string) UUID
	NewV4() (UUID, error)
	NewV5(ns UUID, name string) UUID
}

type GoogleUUIDGenerator struct{}

func (GoogleUUIDGenerator) NewV1() (UUID, error) {
	return uuid.NewUUID()
}

func (GoogleUUIDGenerator) NewV2(domain byte) (UUID, error) {
	switch domain {
	case DomainPerson:
		return uuid.NewDCEPerson()
	case DomainGroup:
		return uuid.NewDCEGroup()
	}
	return uuid.NewUUID()
}

func (GoogleUUIDGenerator) NewV3(ns UUID, name string) UUID {
	return uuid.NewMD5(ns, []byte(name))
}

func (GoogleUUIDGenerator) NewV4() (UUID, error) {
	return uuid.New(), nil
}

func (GoogleUUIDGenerator) NewV5(ns UUID, name string) UUID {
	return uuid.NewSHA1(ns, []byte(name))
}

func newGoogleUUIDGenerator() *GoogleUUIDGenerator {
	return &GoogleUUIDGenerator{}
}
