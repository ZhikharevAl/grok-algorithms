package main


func alternateMax(a []int) int {
    if len(a) == 0 {
        return -1
    }

    for i := range a {
        isMax := true
        for j := range a {
            if a[i] < a[j] {
                isMax = false
                break
            }
        }
        if isMax {
            return a[i]
        }
    }
    return a[0]
}