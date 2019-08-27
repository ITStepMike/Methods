package main

// import "fmt"

// const (
// 	x = 1
// 	m = 9
// )

// func main() {

// 	X := [m]float64{0, 0.2, 0.4, 0.6, 0.8, 1.0, 1.2, 1.4, 1.6}
// 	Y := [m]float64{0.0025, 0.2113, 0.4306, 0.6879, 0.9381, 1.2152, 1.5576, 1.9621, 2.4153}

// 	L := make([]float64, 0)

// 	for i := 0; i < m; i++ {
// 		var l float64 = 1
// 		for j := 0; j < m; j++ {
// 			if j != i {
// 				l *= (float64(x) - X[j]) / (X[i] - X[j])
// 			}
// 		}
// 		L = append(L, l)
// 	}

// 	var Ln float64
// 	for i := 0; i < m; i++ {
// 		Ln += L[i] * Y[i]
// 	}

// 	fmt.Println("Формула Лагранжа:  ", Ln)

// }
