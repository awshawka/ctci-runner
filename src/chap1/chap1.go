package chap1

import (
	"standard"
)

type Chap1 struct{

}

func (chap1 *Chap1) Run(q int) {
	solutions := []standard.Solution{&Q1{}}

	solutions[q].Test()
}



