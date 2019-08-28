package main

// import (
// 	"fmt"
// 	"math"
// )

// var X = [9]float64{0, 0.2, 0.4, 0.6, 0.8, 1.0, 1.2, 1.4, 1.6}
// var Y = [9]float64{0.0025, 0.2113, 0.4306, 0.6879, 0.9381, 1.2152, 1.5576, 1.9621, 2.4153}

// func fact(f int) int {
// 	if f == 0 {
// 		return 1
// 	}
// 	return f * fact(f-1)
// }

// func first(x float64, a []float64) float64 {
// 	Pn := a[0]
// 	for i := 1; i < len(X); i++ {
// 		var M float64 = 1
// 		for m := 0; m < i; m++ {
// 			M *= (x - X[m])
// 		}
// 		Pn += a[i] * M
// 	}

// 	return Pn
// }

// func main() {

// 	n := len(X)
// 	h := 0.2
// 	XX := []float64{0.27, 0.92, 1.65}

// 	DY := [9][9]float64{{0.0025, 0.2113, 0.4306, 0.6879, 0.9381, 1.2152, 1.5576, 1.9621, 2.4153}}

// 	for i := 1; i < n; i++ {
// 		for j := 0; j < n-i; j++ {
// 			DY[i][j] = (DY[i-1][j+1] - DY[i-1][j])
// 		}
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n-i; j++ {
// 			fmt.Printf("%-19v ", DY[i][j])
// 		}
// 		fmt.Println()
// 	}

// 	a := make([]float64, 0)
// 	for i := 0; i < n; i++ {
// 		a = append(a, DY[i][0]/(float64(fact(i))*(math.Pow(h, float64(i)))))
// 	}

// 	YY := make([]float64, 0)

// 	for i := 0; i < len(XX); i++ {
// 		YY = append(YY, first(XX[i], a))
// 	}

// 	fmt.Println("Перша інтерполяційна формула Ньютона: ", YY)

// }
