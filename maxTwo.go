package main

func largestTwo(A []int) (int, int) {
	if len(A) < 2 {
		panic("Слайс должен содержать как минимум два элемента")
	}

	myMax, second := A[0], A[1]
	if myMax < second {
		myMax, second = second, myMax
	}

	for _, val := range A[2:] {
		if myMax < val {
			myMax, second = val, myMax
		} else if second < val {
			second = val
		}
	}

	return myMax, second
}
