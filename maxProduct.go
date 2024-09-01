package main


func maxProduct(nums []int) int {
    myMax, second := nums[0], nums[1]
	if myMax < second {
		myMax, second = second, myMax
	}

	for _, val := range nums[2:] {
		if myMax < val {
			myMax, second = val, myMax
		} else if second < val {
			second = val
		}
	}
	return (myMax - 1) * (second - 1)
}