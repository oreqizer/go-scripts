// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	var acc int
	for x > 0 {
		acc++
		x = x & (x - 1)
	}
	return acc

	//var acc int
	//for i := 0; i < 64; i++ {
	//	if x & 1 != 1 {
	//		acc++
	//	}
	//	x = x >> 1
	//}
	//return acc

	//var acc int
	//for i := 0; i < 8; i++ {
	//	acc += int(pc[byte(x >> uint64(i * 8))])
	//}
	//return acc

	//return int(pc[byte(x>>(0*8))] +
	//	pc[byte(x>>(1*8))] +
	//	pc[byte(x>>(2*8))] +
	//	pc[byte(x>>(3*8))] +
	//	pc[byte(x>>(4*8))] +
	//	pc[byte(x>>(5*8))] +
	//	pc[byte(x>>(6*8))] +
	//	pc[byte(x>>(7*8))])
}

//!-
