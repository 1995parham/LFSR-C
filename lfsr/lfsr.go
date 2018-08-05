/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     lfsr.go
 * +===============================================
 */

package lfsr

// LFSR8 Linear Feedback Shift Register 8bit
type LFSR8 interface {
	Init(poly uint8, seed uint8)
	Next() uint8
}

// LFSR16 Linear Feedback Shift Register 16bit
type LFSR16 interface {
	Init(poly uint16, seed uint16)
	Next() uint16
}

// LFSR32 Linear Feedback Shift Register 32bit
type LFSR32 interface {
	Init(poly uint32, seed uint32)
	Next() uint32
}

// LFSR64 Linear Feedback Shift Register 64bit
type LFSR64 interface {
	Init(poly uint64, seed uint64)
	Next() uint64
}
