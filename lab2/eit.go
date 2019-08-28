package main

// import "fmt"

// const (
// 	x = 2
// 	m = 9
// )

// func main() {

// 	X := [m]float64{0, 0.2, 0.4, 0.6, 0.8, 1.0, 1.2, 1.4, 1.6}
// 	Y := [m]float64{0.0025, 0.2113, 0.4306, 0.6879, 0.9381, 1.2152, 1.5576, 1.9621, 2.4153}
// 	var P [9][9]float64

// 	for n := 0; n < m; n++ {
// 		j := n
// 		for i := 0; i < (9-n) && j < 9; i++ {
// 			if i == j {
// 				P[i][j] = Y[i]
// 			} else if i < j {
// 				t := ((P[i][j-1])*(X[j]-x) - (P[i+1][j])*(X[i]-x)) / (X[j] - X[i])
// 				P[i][j] = t
// 			}
// 			j++
// 		}
// 	}

// 	for i := 0; i < m; i++ {
// 		for j := 0; j < m; j++ {
// 			if i < j {
// 				fmt.Printf("%-22v", P[i][j])
// 			}
// 		}
// 		fmt.Println()
// 	}

// }
