package main

import (
	"chap1"
	"os"
	"fmt"
	"standard"

	
)

func main(){
	if len(os.Args) < 3 {
		fmt.Println("ERROR: must specify the chapter and question")
		return
	}

	c := parseArgument(os.Args[1])
	q := parseArgument(os.Args[2])

	runners := []standard.Runner{&chap1.Chap1{}}

	runners[c-1].Run(q-1)
}

func parseArgument(input string) int {
	lastChar := input[len(input)-1]
	return int(lastChar - '0')
}