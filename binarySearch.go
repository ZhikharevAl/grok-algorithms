package main

func BinarySearch(arr []int, target int) int {
    low, high := 0, len(arr)-1
    for low <= high {
        mid := low + (high-low)/2
        switch {
        case arr[mid] == target:
            return mid
        case arr[mid] > target:
            high = mid - 1
        default:
            low = mid + 1
        }
    }
    return -1
}
