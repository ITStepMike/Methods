package main

// import (
// 	"fmt"
// 	"math"
// )

// //y' = cos(2x + y) + 1.5 * (x - y)
// //y(0) = 1.0
// //[a, b] = [0, 1]

// var Dy_e []float64

// func f(x, y float64) float64 {
// 	return math.Cos(2*x+y) + 1.5*(x-y)
// }

// const (
// 	a, b float64 = 0, 1
// 	m    int     = 10
// )

// func main() {

// 	h := (b - a) / float64(m)
// 	x := make([]float64, m+1)
// 	y := make([]float64, 0)
// 	y = append(y, 1)

// 	for i := 0; i < m+1; i++ {
// 		x[i] = a + float64(i)*h
// 	}

// 	for i := 0; i < m; i++ {
// 		dy := h * f(x[i], y[i])
// 		Dy_e = append(Dy_e, dy)
// 		y = append(y, y[i]+dy)
// 	}

// 	for i := 0; i < m; i++ {
// 		fmt.Printf("X%v: %-20v   Y%v: %-20v   dy: %-20v\n", i, x[i], i, y[i], Dy_e[i])
// 	}

// }
