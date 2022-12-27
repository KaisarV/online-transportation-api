package main

import "fmt"

func main() {

	a := "Kaisar"
	b := &a
	*b = "Febi"
	fmt.Println(*b)

}
