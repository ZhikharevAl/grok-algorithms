package main

func count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return 1 + count(arr[1:])
}