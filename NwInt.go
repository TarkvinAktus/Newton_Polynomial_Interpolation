package main

//Посроение интерполяционного многочлена Ньютона
import (
	"fmt"
	"math"
	"strconv"
)

//Функция f(a) - коэффициентов полинома
//Составляется из разделённых разностей полинома (см теорию http://matlab.exponenta.ru/spline/book1/14.php)
//fa(Xi,Xi+1) = f(Xi+1)-f(Xi) / Xi+1 - Xi
//fa(X1,X2,X3) = fa(X2,X3) - fa(X1,X2) / X3 - X1
//И т.д.
//На вход i - номер коэффициента; divDepth - шаг деления
//x - иксы, fx - игреки
func fa(i int, divDepth int, x []int, fx []int) float64 {
	var result float64
	var div float64
	if i == 1 {
		result = float64(fx[i+divDepth] - fx[i+divDepth-1])
		div = float64(x[i+divDepth] - x[i+divDepth-1])
		if div != 0 {
			result /= div
		} else {
			result = 0
		}
	} else {
		result = (fa(i-1, divDepth+1, x, fx) - fa(i-1, divDepth, x, fx))
		div = float64(x[i+divDepth] - x[0+divDepth])
		if div != 0 {
			result /= div
			result = toFixed(result, 2)
		} else {
			result = 0
		}
	}
	return result
}

// A[] represents coefficients of first polynomial 
// B[] represents coefficients of second polynomial 
// m and n are sizes of A[] and B[] respectively 
func multiply(A []int, B []int, m int, n int) []int { 
   prod := [m+n-1]int 
   // Initialize the porduct polynomial 
   for i := 0; i<m+n-1; i++ {
	   prod[i] = 0 
   }
   // Multiply two polynomials term by term 
  
   // Take ever term of first polynomial 
   for i:=0; i<m; i++ { 
     // Multiply the current term of first polynomial 
     // with every term of second polynomial. 
     for (int j=0; j<n; j++){
		prod[i+j] += A[i]*B[j]
	 }  
   } 
  
   return prod
} 

func polynom(x []int, y []int, Xpoint int) float64 {
	var PolynomSum float64
	var buf float64
	var resultStr string

	resultStr += string(strconv.Itoa(y[0]))

	PolynomSum += float64(y[0])

	for i := 1; i < (len(x)); i++ {
		resultStr += "+"
		
		buf = fa(i, 0, x, y)
		s := fmt.Sprintf("%f", buf)
		resultStr += s
		
		
		for j := 0; j < i; j++ {
			buf *= float64(Xpoint - x[j])
			resultStr += "(x-"
			resultStr += strconv.Itoa(x[j])
			resultStr += ")"

		}
		PolynomSum += buf
	}
	fmt.Println("result = ",resultStr)
	return math.Round(PolynomSum)
}

func check(x int) float64 {
	var buf float64
	buf = (float64(-17)/float64(6))*(math.Pow(float64(x), 3)) + (math.Pow(float64(x), 2) * 12.5) - (float64(38) / float64(3) * float64(x)) - 2
	return math.Round(buf)

}

func main() {

	x := []int{0, 1, 2, 3}
	fx := []int{-2, -5, 0, -4}
	/*
		fmt.Println("fa(10) = ",fa(1,0,x,fx))
		fmt.Println("fa(11) = ",fa(1,1,x,fx))
		fmt.Println("fa(12) = ",fa(1,2,x,fx))
		fmt.Println("fa(20) = ",fa(2,0,x,fx))
		fmt.Println("fa(21) = ",fa(2,1,x,fx))
		fmt.Println("fa(30) = ",fa(3,0,x,fx))
	*/

	fmt.Println("___check 0 ", check(0))
	fmt.Println("___check 1 = ", check(1))
	fmt.Println("___check 2 = ", check(2))
	fmt.Println("___check 3 = ", check(3))
	fmt.Println("___check 3 = ", check(4))

	fmt.Println("polynom 0 = ", polynom(x, fx, 0))
	fmt.Println("polynom 1 = ", polynom(x, fx, 1))
	fmt.Println("polynom 2 = ", polynom(x, fx, 2))
	fmt.Println("polynom 3 = ", polynom(x, fx, 3))
	fmt.Println("polynom 4 = ", polynom(x, fx, 4))

	var scan string
	fmt.Scan(&scan)
}
