package main

func max(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	subMax := max(arr[1:])
	if arr[0] > subMax {
		return arr[0]
	}
	return subMax
}
