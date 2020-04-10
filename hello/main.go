package main

import "fmt"

func main() {
	var a [4]int

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %2d  ", j, i, i*j)
		}
		fmt.Println()
	}
	fmt.Scanf("%d %d %d %d", &a[0], &a[1], &a[2], &a[3])
	fmt.Printf("%d", a[0])
}
