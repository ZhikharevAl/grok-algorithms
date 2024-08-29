package main

import "fmt"

var voted = make(map[string]bool)

func checkVoter(name string) {
	if voted[name] {
		fmt.Println("Выгнать его!")
	} else {
		voted[name] = true
		fmt.Println("Допустить к голосованию!")
	}
}
