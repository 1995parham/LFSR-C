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

type LFSR8 interface {
	Init(poly uint8, seed uint8)
	Next() uint8
}
