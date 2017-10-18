/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 20-05-2016
 * |
 * | File Name:     flfsr8_test.go
 * +===============================================
 */
package flfsr

import (
	"testing"

	lfsr "github.com/1995parham/LFSR.go"
)

func Test8One(t *testing.T) {
	t.Logf("polynomial: x^7 + x^5 + x^4 + x^3 + 1, 1011 1000 = 0xB8\n")

	var f, d lfsr.LFSR8
	/* FLFSR initiation */
	f = NewFLFSR8()
	f.Init(0xB8, 0x40)

	/* DLFSR initiation */
	d = NewDummyLFSR8()
	d.Init(0x0, 0x40)

	for i := 0; i < 16; i++ {
		var dd, fd uint8
		dd = d.Next()
		fd = f.Next()

		if fd != dd {
			t.Errorf("Fail on %d != %d", dd, fd)
		}
	}
}

func NewDummyLFSR8() lfsr.LFSR8 {
	return &dlfsr8{}
}

type dlfsr8 struct {
	data uint8
}

func (d *dlfsr8) Init(poly uint8, seed uint8) {
	d.data = seed
}

func (d *dlfsr8) Next() uint8 {
	d.data = ((((d.data >> 7) ^ (d.data >> 5) ^ (d.data >> 4) ^ (d.data >> 3)) & 0x01) | (d.data << 1))
	return d.data
}
