package main

import (
	"os"
	"fmt"

	"chap1"
)

func main(){
	if len(os.Args) < 3 {
		fmt.Println("ERROR: must specify the chapter and question")
		return
	}

	chapter := parseArgument(os.Args[1])
	question := parseArgument(os.Args[2])

	chap1.
}

func parseArgument(input string) int {
	lastChar := input[len(input)-1]
	return int(lastChar - '0')
}