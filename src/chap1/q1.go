package chap1

import (
	"fmt"
)

type Q1 struct {

}

func (q *Q1) Test(){
	fmt.Println(reverseString("waseef"))
}

// ReverseString ..
func reverseString(input string) string {
	return input
}

