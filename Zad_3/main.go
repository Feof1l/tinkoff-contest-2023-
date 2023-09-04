package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&b[i])
	}
	start_pos := -1
	end_pos := -1
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			if start_pos == -1 {
				start_pos = i
			}
			end_pos = i
		}
	}
	if start_pos == -1 {
		fmt.Println("YES")
		return
	}
	len := end_pos - start_pos
	sorted := make([]int, len)
	copy(sorted, a[start_pos:end_pos+1])
	sort.Ints(sorted)
	for i := start_pos; i <= end_pos; i++ {
		if sorted[i-start_pos] != b[i] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
