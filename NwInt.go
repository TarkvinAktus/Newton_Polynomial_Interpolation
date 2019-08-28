package main

//Посроение интерполяционного многочлена Ньютона
import (
	"fmt"
	"math"
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

// НЕ ПРОТЕСТИРОВАНО
// A[] represents coefficients of first polynomial
// B[] represents coefficients of second polynomial
// m and n are sizes of A[] and B[] respectively
func multiply(Result *[]int, A []int, B []int, m int, n int) {
	prod := make([]int, m+n-1)
	// Initialize the porduct polynomial
	for i := 0; i < m+n-1; i++ {
		prod[i] = 0
	}
	// Multiply two polynomials term by term

	// Take ever term of first polynomial
	for i := 0; i < m; i++ {
		// Multiply the current term of first polynomial
		// with every term of second polynomial.
		for j := 0; j < n; j++ {
			prod[i+j] += A[i] * B[j]
		}
	}
	for i := 0; i < m+n-1; i++ {
		(*Result)[i] += prod[i]
	}
}

//polynom возвращает координату точки с известным х (Xpoint)
//на основе слайсов с координатами остальных точек
//ПРОТЕСТИРОВАНО
func polynom(x []int, y []int, Xpoint int) float64 {
	var PolynomSum float64
	var buf float64
	//var resultStr string

	PolynomSum += float64(y[0])
	//resultStr += string(strconv.Itoa(y[0]))
	for i := 1; i < (len(x)); i++ {
		buf = fa(i, 0, x, y)
		//resultStr += "+"
		//s := fmt.Sprintf("%f", buf)
		//resultStr += s
		for j := 0; j < i; j++ {
			buf *= float64(Xpoint - x[j])
			//resultStr += "(x-"
			//resultStr += strconv.Itoa(x[j])
			//resultStr += ")"
		}
		PolynomSum += buf
	}
	//fmt.Println("result = ",resultStr)
	return math.Round(PolynomSum)
}

//В РАБОТЕ
//Комментарии
//ADD RETURN
//polynomialCoefficents Будет возвращать коэффициенты вычисленного полинома
//на основе слайсов с координатами известных точек
func polynomialCoefficents(x []int, y []int) {
	PolCoefficents := make([]float64, len(x))
	var AnCoefficient float64

	PolCoefficents[0] = float64(y[0])

	for i := 1; i < (len(x)); i++ {

		AnCoefficient = fa(i, 0, x, y)

		PolBuf := make([]float64, len(x)*2-1)

		//Pol member = Fn(x-n1)(x-n2)...(x-n-1)
		for j := 0; j < i-1; j++ {
			MulBuf := make([]int, len(x)*2-1)
			//Проверить корректность умножения 0 как члена полинома

			SecondMul := make([]int, 2)
			SecondMul[0] = -x[j+1]
			SecondMul[1] = 1
			if j == 0 {
				FirstMul := make([]int, 2)
				FirstMul[0] = -x[j]
				FirstMul[1] = 1
				multiply(&MulBuf, FirstMul, SecondMul, len(FirstMul), len(SecondMul))
			} else {
				multiply(&MulBuf, MulBuf, SecondMul, len(MulBuf), len(SecondMul))
			}

			for i := 0; i < len(MulBuf); i++ {
				PolBuf[i] += float64(MulBuf[i])
				//if coefficient ==0 => stop?
			}
		}
		for i := 0; i < len(PolBuf); i++ {
			//if coefficient ==0 => stop?
			PolCoefficents[i] += PolBuf[i] * AnCoefficient
		}
	}
}

//Вычисление тестового примера для
//x := []int{0, 1, 2, 3}
//fx := []int{-2, -5, 0, -4}
//Источник - http://kontromat.ru/?page_id=4955
func check(x int) float64 {
	var buf float64
	buf = (float64(-17)/float64(6))*(math.Pow(float64(x), 3)) + (math.Pow(float64(x), 2) * 12.5) - (float64(38) / float64(3) * float64(x)) - 2
	return math.Round(buf)

}

func test() {

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
