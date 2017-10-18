/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     flfsr8.go
 * +===============================================
 */
package flfsr

import lfsr "github.com/1995parham/LFSR.go"

func NewFLFSR8() lfsr.LFSR8 {
	return &flfsr8{}
}

type flfsr8 struct {
	data uint8
	poly uint8
}

func (f *flfsr8) Init(poly uint8, seed uint8) {
	f.poly = poly
	f.data = seed
}

func (f *flfsr8) Next() uint8 {
	var bit uint8
	var i uint8

	for i = 0; i < 8; i++ {
		if f.poly&(1<<i) != 0 {
			bit ^= (f.data >> i)
		}
	}
	bit &= 0x01

	f.data = (f.data << 1) | bit

	return f.data
}
