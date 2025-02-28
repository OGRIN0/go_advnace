package main 

import "fmt"

func createPointer() *int {
	x := 42
	return &x 
}

func example(){
	p := createPointer()
	fmt.Println(*p)
}