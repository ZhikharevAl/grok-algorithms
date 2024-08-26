package main

func findSmallest(arr []int) int {
	smallestIndex := 0
	smallest := arr[0]
	for i := 1; i < len(arr); i++ {
        if arr[i] < smallest {
            smallest = arr[i]
            smallestIndex = i
        }
	}
	return smallestIndex
}
