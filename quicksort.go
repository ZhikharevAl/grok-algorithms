package main

func quicksort(array []int) []int {

	if len(array) < 2 {
		return array

	} else {
		pivot := array[0]
		var less []int
		var greater []int

		for _, item := range array[1:] {
			if item <= pivot {
				less = append(less, item)
			} else {
				greater = append(greater, item)
			}
		}

		less = quicksort(less)
		greater = quicksort(greater)

		return append(append(less, pivot), greater...)
	}
}
