package testlib

import "strings"

/* This is a proto-example of a simple library of functions.  The purpose of the
 * "library" is simply to provide a few simple functions that can be tested from a
 * separate "testing" repository / project
 */

// Positive returns a positive integer.
func Positive(a, b int) int {
	r := b
	if a > b {
		r = a
	}
	if r < 0 {
		r = -1
	}
	return r
}

// Clean returns a string with both leading and trailing white space stipped.
func Clean(instr string) string {
	return strings.Trim(instr, " \t\r\n")
}
