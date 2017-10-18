/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     flfsr16.go
 * +===============================================
 */
package flfsr

import lfsr "github.com/1995parham/LFSR.go"

func NewFLFSR16() lfsr.LFSR16 {
	return &flfsr16{}
}

type flfsr16 struct {
	data uint16
	poly uint16
}

func (f *flfsr16) Init(poly uint16, seed uint16) {
	f.poly = poly
	f.data = seed
}

func (f *flfsr16) Next() uint16 {
	var bit uint16
	var i uint16

	for i = 0; i < 16; i++ {
		if f.poly&(1<<i) != 0 {
			bit ^= (f.data >> i)
		}
	}
	bit &= 0x01

	f.data = (f.data << 1) | bit

	return f.data
}
