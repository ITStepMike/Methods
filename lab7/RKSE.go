package main

// import (
// 	"fmt"
// 	"math"
// )

// //dx(t)/dt = sin(t + x - y^2)
// //dy(t)/dt = t^2 - xy

// //x(0.5) = 1
// //y(0.1) = 2

// //t0, t = 2, 3.5

// const (
// 	t0, T float64 = 0.5, 3
// 	m             = 10
// )

// func fxt(x, y, t float64) float64 {
// 	return math.Sin(t + x - y*y)
// }

// func fyt(x, y, t float64) float64 {
// 	return t*t - x*y
// }

// func main() {

// 	h := float64((T - t0) / m)
// 	x := make([]float64, 0)
// 	y := make([]float64, 0)
// 	t := make([]float64, 0)
// 	dy := make([]float64, 0)
// 	dx := make([]float64, 0)

// 	for i := 0; i < m+1; i++ {
// 		t = append(t, t0+float64(i)*h)
// 	}

// 	x = append(x, 1)
// 	y = append(y, 2)

// 	for i := 0; i < m; i++ {
// 		k1x := h * fxt(x[i], y[i], t[i])
// 		k1y := h * fyt(x[i], y[i], t[i])
// 		k2x := h * fxt(x[i]+0.5*k1x, y[i]+0.5*k1y, t[i]+0.5*h)
// 		k2y := h * fyt(x[i]+0.5*k1x, y[i]+0.5*k1y, t[i]+0.5*h)
// 		k3x := h * fxt(x[i]+0.5*k2x, y[i]+0.5*k2y, t[i]+0.5*h)
// 		k3y := h * fyt(x[i]+0.5*k2x, y[i]+0.5*k2y, t[i]+0.5*h)
// 		k4x := h * fxt(x[i]+0.5*k3x, y[i]+0.5*k3y, t[i]+h)
// 		k4y := h * fyt(x[i]+0.5*k3x, y[i]+0.5*k3y, t[i]+h)

// 		dx = append(dx, (k1x+2*k2x+2*k3x+k4x)*0.166666667)
// 		x = append(x, x[i]+dx[i])
// 		dy = append(dy, (k1y+2*k2y+2*k3y+k4y)*0.166666667)
// 		y = append(y, y[i]+dy[i])
// 	}

// 	for i := 0; i < m; i++ {
// 		fmt.Printf("X%v: %-20v   Y%v: %-20v   dy: %-20v   dx: %-22v   t%v: %-20v\n", i, x[i], i, y[i], dy[i], dx[i], i, t[i])
// 	}

// }
