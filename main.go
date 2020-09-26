package main

import (
	"os"
	"fmt"
)

func main(){
	if len(os.Args) < 2 {
		fmt.Println("must specify the chapter")
		return
	}
}