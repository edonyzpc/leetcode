package main

import "fmt"

// to find assembly code of slice address, run command line `env GOSSAFUNC=main go build main.go`
func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := a[0:4]

	fmt.Printf("a=%v, len=%d, cap=%d\n", a, len(a), cap(a))
	fmt.Printf("b=%v, len=%d, cap=%d\n", b, len(b), cap(b))

	b = append(b, 10, 11, 12)

	fmt.Printf("a=%v, len=%d, cap=%d\n", a, len(a), cap(a))
	fmt.Printf("b=%v, len=%d, cap=%d\n", b, len(b), cap(b))

	var c []int
	c = append(c, a...)
	// or
	// c := append([]int(nil), a...)
	// or(
	// **NOTE**:
	//    1. copy will not do copy if dst length equals 0)
	//    2. copy is faster than append
	// c1 := make([]int, len(a[0:4]))
	// copy(c1, a[0:4])
	// fmt.Printf("c1 = %v\n", c1)
	fmt.Printf("a=%v, len=%d, cap=%d\n", a, len(a), cap(a))
	fmt.Printf("c=%v, len=%d, cap=%d\n", c, len(c), cap(c))

	c = append(c, 13, 14, 15)

	fmt.Printf("a=%v, len=%d, cap=%d\n", a, len(a), cap(a))
	fmt.Printf("c=%v, len=%d, cap=%d\n", c, len(c), cap(c))
}
