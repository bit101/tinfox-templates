// Package main is the main program
package main

import (
	"fmt"
	"github.com/bit101/bitlib/random"
)

//////////////////////////////
// MAIN
//////////////////////////////

func main() {
	r := random.FloatRange(0, 100)
	fmt.Printf("random float: %f\n", r)
}
