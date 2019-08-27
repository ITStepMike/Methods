package main

import (
	"fmt"
	"math"
)

//y' = cos(2x + y) + 1.5 * (x - y)
//y(0) = 1.0
//[a, b] = [0, 1]

func f(x, y float64) float64 {
	return math.Cos(2*x+y) + 1.5*(x-y)
}

const (
	a, b float64 = 0, 1
	m    int     = 10
)

func main() {

	var K1, K2, K3, K4, Dy_rk []float64
	x := make([]float64, 0)
	y := make([]float64, 0)
	h := (b - a) / float64(m)
	y = append(y, 1)

	for i := 0; i < m+1; i++ {
		x = append(x, a+float64(i)*h)
	}

	for i := 0; i < m; i++ {
		k1 := h * f(x[i], y[i])
		k2 := h * f(x[i]+0.5*h, y[i]+0.5*k1)
		k3 := h * f(x[i]+0.5*h, y[i]+0.5*k2)
		k4 := h * f(x[i]+h, y[i]+k3)

		K1 = append(K1, k1)
		K2 = append(K2, k2)
		K3 = append(K3, k3)
		K4 = append(K4, k4)

		Dy_rk = append(Dy_rk, (k1+2*k2+2*k3+k4) * 0.166666667)
		y = append(y, y[i]+Dy_rk[i])
	}

	for i := 0; i < m; i++ {
		fmt.Printf("X%v: %-20v   Y%v: %-20v   dy: %-20v\n", i, x[i], i, y[i], Dy_rk[i])
	}

}
