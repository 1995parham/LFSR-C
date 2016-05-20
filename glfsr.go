/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     glfsr.go
 * +===============================================
 */
package lfsr

func NewGLFSR8() LFSR8 {
	return &glfsr8{}
}

type glfsr8 struct {
	data uint8
	mask uint8
	poly uint8
}

func (g *glfsr8) Init(poly uint8, seed uint8) {
	var shift uint
	shift = 7

	g.poly = poly | 1
	g.data = seed

	var seed_mask uint8
	seed_mask = 1
	seed_mask <<= shift

	for ; shift != 0; shift-- {
		if (poly & seed_mask) != 0 {
			g.mask = seed_mask
			break
		}
		seed_mask >>= 1
	}
}

func (g *glfsr8) Next() uint8 {
	g.data <<= 1

	if g.data&g.mask != 0 {
		g.data ^= g.poly
	}

	return g.data
}
