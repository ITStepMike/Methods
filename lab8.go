package main

// import "fmt"

// //Диф рівняння першого порядку
// //y'' + 2xy' - y = 0.4
// //Відповідно
// //p(x) = 2x, q(x) = 1, f(x) = 0.4

// //Початкова умова
// //2y(0.3) + y'(0.3) = 1
// //y'(0.6) = 2
// //Відповідно
// //al1 = 2 bt1 = 1 ym1 = 1
// //al2 = 0 bt2 = 1 ym2 = 2

// //p(x)
// func px(x float64) float64 {
// 	return 2 * x
// }

// //q(x)
// func qx(x float64) float64 {
// 	return -1
// }

// //f(x)
// func fx(x float64) float64 {
// 	return 0.4
// }

// const (
// 	a, b          float64 = 0.3, 0.6 //Відрізок [a, b]
// 	n                     = 10       //n - частин
// 	al1, bt1, ym1 float64 = 2, 1, 1  //al1 = 2 bt1 = 1 ym1 = 1
// 	al2, bt2, ym2 float64 = 0, 1, 2  //al2 = 0 bt2 = 1 ym2 = 2
// )

// func main() {

// 	var h = (b - a) / n
// 	var x []float64
// 	var A, B, C, D []float64
// 	var P, Q, Y []float64
// 	var AA [n][n]float64

// 	fmt.Println("H: ", h)
// 	fmt.Println()

// 	//find Xi
// 	for i := 0; i < n; i++ {
// 		x = append(x, a+float64(i)*h)
// 	}

// 	//find Ai
// 	for i := 0; i < n; i++ {
// 		A = append(A, 1-(px(x[i])*h)/2)
// 	}

// 	//find Bi
// 	for i := 0; i < n; i++ {
// 		B = append(B, 1+(px(x[i])*h)/2)
// 	}

// 	//find Ci
// 	for i := 0; i < n; i++ {
// 		C = append(C, -2+h*h*qx(x[i]))
// 	}

// 	//find Di
// 	for i := 0; i < n; i++ {
// 		D = append(D, h*h*fx(x[i]))
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			if i == j {
// 				AA[i][j] = C[i]
// 			} else if i < n-1 && j == i+1 {
// 				AA[i][j] = B[i]
// 			} else if i >= 1 && j == i-1 {
// 				AA[i][j] = A[i]
// 			}
// 		}
// 	}

// 	C[0] = -bt1 + h*al1
// 	B[0] = bt1
// 	D[0] = h * ym1

// 	A[n-1] = -bt2
// 	C[n-1] = bt2 + h*al2
// 	D[n-1] = h * ym2

// 	P = append(P, B[0]/-C[0])
// 	Q = append(Q, D[0]/-C[0])

// 	for i := 1; i < n; i++ {
// 		P = append(P, B[i]/(-C[0]-A[i]*P[i-1]))
// 		Q = append(Q, (A[i]*Q[i-1]-D[i])/(-C[i]-A[i]*P[i-1]))
// 	}

// 	Y = make([]float64, n)

// 	Y[n-1] = (A[n-1]*Q[n-2] - D[n-1]) / (-C[n-1] - A[n-1]*P[n-2])

// 	for i := n - 1; i > 0; i-- {
// 		Y[i-1] = P[i-1]*Y[i] + Q[i-1]
// 	}

// 	for i := 0; i < n; i++ {
// 		for j := 0; j < n; j++ {
// 			fmt.Print(" ", AA[i][j])
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println()
// 	fmt.Println(Y)
// 	fmt.Println()
// 	fmt.Println(D)

// 	fmt.Println("Перевірка")
// 	fmt.Println("al2*Yn + bt2*(Yn - Yn-1)/h  =  ym2")
// 	fmt.Println("al2*Yn + bt2*(Yn - Yn-1)/h = ", al2*Y[n-1]+bt2*(Y[n-1]-Y[n-2])/h)
// 	fmt.Println("ym2  =  ", ym2)

// 	if al2*Y[n-1]+bt2*(Y[n-1]-Y[n-2])/h == ym2 {
// 		fmt.Println("Fits")
// 	}

// }
