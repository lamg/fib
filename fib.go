// Fibonacci numbers calculator
package main

import (
	"flag"
)

func main() {
	var N = flag.Uint64("N", 0, "N ≥ 0")

	flag.Parse()
	r := fibonacci(*N)
	println(r)
}

func fibonacci(N uint64) (x uint64) {
  var n, y uint64
	//{ N ≥ 0}
	n, x, y = 0, 0, 1
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
