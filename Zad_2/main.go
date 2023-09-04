package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	res := ProveStrok(text)
	fmt.Println(res)
}
func ProveStrok(s string) int {
	res := 0

	m := make(map[string]int)
	dlina := 0
	s = strings.Trim(s, "\n\r")
	r := []rune(s)
	for _, i := range s {
		_ = i
		dlina++
	}
	for i := 0; i < dlina; i++ {
		switch r[i] {
		case 's':
			m["s"] = m["s"] + 1
		case 'h':
			m["h"] = m["h"] + 1
		case 'e':
			m["e"] = m["e"] + 1
		case 'r':
			m["r"] = m["r"] + 1
		case 'i':
			m["i"] = m["i"] + 1
		case 'f':
			m["f"] = m["f"] + 1
		}
	}
	for (m["s"] > 0) && (m["h"] > 0) && (m["e"] > 0) && (m["r"] > 0) && (m["i"] > 0) && (m["f"] > 1) {
		if (m["s"] > 0) && (m["h"] > 0) && (m["e"] > 0) && (m["r"] > 0) && (m["i"] > 0) && (m["f"] > 1) {
			res = res + 1
			m["s"] = m["s"] - 1
			m["h"] = m["h"] - 1
			m["e"] = m["e"] - 1
			m["r"] = m["r"] - 1
			m["i"] = m["i"] - 1
			m["f"] = m["f"] - 2
		}
	}
	return res
}
