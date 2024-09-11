package main

func minBitFlips(start int, goal int) int {
	xor := start ^ goal
	count := 0
	for xor != 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}
