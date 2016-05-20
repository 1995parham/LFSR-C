/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 20-05-2016
 * |
 * | File Name:     glfsr_test.go
 * +===============================================
 */
package lfsr

import (
	"testing"
)

func TestGLFSR(t *testing.T) {
	var l LFSR8
	l = &glfsr8{}
	l.Init()
}
