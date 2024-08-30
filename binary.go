package main

import "strconv"



func binary(n int) string {
	if n == 0 {
		return "0"
	}

	binary := ""
	for n > 0 {
		binary = strconv.Itoa(n%2) + binary
		n /= 2

	}
	return binary
}