package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"gonum.org/v1/gonum/mat"
)

func newRandDense(n int) *mat.Dense {
	d := make([]float64, n*n)
	for i := range d {
		d[i] = rand.NormFloat64()
	}
	return mat.NewDense(n, n, d)
}

func main() {
	n := 500
	iters := 100

	A := newRandDense(n)
	B := newRandDense(n)
	C := mat.NewDense(n, n, nil)

	start := time.Now()
	for i := 0; i < iters; i++ {
		C.Mul(A, B)
	}
	elapsed := time.Since(start)
	fmt.Printf("n=%d, iters=%d, elapsed=%v\n", n, iters, elapsed)
}
