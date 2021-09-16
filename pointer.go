package main

import "fmt"

func main() {
	var a int = 13
	var b *int = &a
	fmt.Println(a)
	fmt.Println(*b)

}
