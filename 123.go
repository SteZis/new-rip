package main

import "fmt"

func main() {
	add(2, 2)
}
func add(x, c int) {
	var z = x + c
	fmt.Println("2 + 2 =", z)
}
