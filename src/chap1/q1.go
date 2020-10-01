package chap1

import (
	"fmt"
	"standard"
	"github.com/kyokomi/emoji"
)

type Q1 struct {

}

func (q *Q1) Test(){
	standard.Equal()
	assert.Equal(testing.T, "321", reverseString("123"))
	fmt.Println(reverseString("waseef"))
}

// ReverseString ..
func reverseString(input string) string {
	return input
}

