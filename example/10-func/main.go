package main

import "fmt"

func add(a int, b int) (int, bool) {
	return a + b, true
}

// func add2(a, b int) int {
// 	return a + b
// }

func exists(m map[string]string, k string) (v string, ok bool) {
	v, ok = m[k]
	return v, ok
}

func main() {
	res, ok := add(1, 2)
	fmt.Println(res, ok) // 3

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok) // A True
}
