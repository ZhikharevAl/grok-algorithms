package main

func isPalindromeRecursive(s string) bool {
    if len(s) <= 1 {
        return true
    }

    if s[0] != s[len(s)-1] {
        return false
    }

    return isPalindromeRecursive(s[1 : len(s)-1])
}