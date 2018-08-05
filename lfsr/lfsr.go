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

// LFSR8 8-bit Linear Feedback Shift Register
type LFSR8 interface {
	Init(poly uint8, seed uint8)
	Next() uint8
}

// LFSR16 16-bit Linear Feedback Shift Register
type LFSR16 interface {
	Init(poly uint16, seed uint16)
	Next() uint16
}

// LFSR32 32-bit Linear Feedback Shift Register
type LFSR32 interface {
	Init(poly uint32, seed uint32)
	Next() uint32
}

// LFSR64 64-bit Linear Feedback Shift Register
type LFSR64 interface {
	Init(poly uint64, seed uint64)
	Next() uint64
}
