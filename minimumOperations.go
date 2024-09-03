package main


func minimumOperations(nums []int) int {
	op:=0
	for _, i:=range(nums) {
		if i % 3 !=0 {
			op+=1
		}
	}
	return op
}