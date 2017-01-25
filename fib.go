package main

import (
	"fmt"
	"os"
	"strconv"
)

var usage = "Usage: %s NaturalNumber\n"

func main() {
	x := 1
	if len(os.Args) == 2 {
		N, e := strconv.Atoi(os.Args[1])
		var r int
		if e == nil && N >= 0 {
			r = fibonacci(N)
			println(r)
			x = 1
		} else {
			fmt.Fprintf(os.Stderr, usage, os.Args[0])
		}
	} else {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
	}
	os.Exit(x)
}

func fibonacci(N int) (x int) {
	//{ N ≥ 0}
	n, x, y := 0, 0, 1
	// P: 0 ≤ n ≤ N ∧ x = fib.n
	// B: n ≠ N
	// { P }
	for n != N {
		// { P ∧ B }
		n, x, y = n+1, y, x+y
		// { P }
	}
	//{ x = fib.N ∧ n = N}
	return
}
