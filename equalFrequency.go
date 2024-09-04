package main



func equalFrequency(word string) bool {
    freq := make(map[rune]int)
    
    for _, c := range word {
        freq[c]++
    }
    
    for key := range freq {
        freq[key]--
        newFreq := make(map[int]int)
        for _, v := range freq {
            if v > 0 {
                newFreq[v]++
            }
        }
        if len(newFreq) == 1 {
            return true
        }
        freq[key]++
    }
    
    return false
}