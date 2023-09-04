package main

import "fmt"

func main() {
	var n, s int
	fmt.Scan(&n, &s)
	a := make([]int, n)
	max := 0
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
		if (a[i] > max) && (a[i] <= s) {
			max = a[i]

		}
	}
	fmt.Println(max)
}
