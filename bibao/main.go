package main

import "fmt"

func add(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f1 := add(5)
	f2 := add(-2)
	fmt.Println(f1(2))
	fmt.Println(f1(8))
	fmt.Println(f2(9))
	fmt.Println(f2(18))
}