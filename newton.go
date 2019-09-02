package main

//Посроение интерполяционного многочлена Ньютона
import (
	"fmt"
	"math"
)

//Fa Функция F(a) - коэффициентов полинома
//Составляется из разделённых разностей полинома (см теорию http://matlab.exponenta.ru/spline/book1/14.php)
//Fa(Xi,Xi+1) = f(Xi+1)-f(Xi) / Xi+1 - Xi
//Fa(X1,X2,X3) = Fa(X2,X3) - Fa(X1,X2) / X3 - X1
//И т.д.
//На вход i - номер коэффициента; divDepth - шаг деления
//x - иксы, fx - игреки
func Fa(i int, divDepth int, x []int, fx []int) float64 {
	var result float64
	var div float64
	if i == 1 {
		result = float64(fx[i+divDepth] - fx[i+divDepth-1])
		div = float64(x[i+divDepth] - x[divDepth])
		if div != 0 {
			result /= div
		} else {
			result = 0
		}
	} else {
		result = (Fa(i-1, divDepth+1, x, fx) - Fa(i-1, divDepth, x, fx))
		div = float64(x[i+divDepth] - x[divDepth])
		if div != 0 {
			result /= div
			result = result
		} else {
			result = 0
		}
	}
	return result
}

// Multiply ...
// Просто поочерёдно умножаем члены каждого многочлена
// Результат умножения в степень равную сумме степеней членов многочлена
// Внутри слайсов они отсортированы по степеням*
func Multiply(Result *[]int, A []int, B []int, m int, n int) {
	bufLen := make([]int, len(*Result))
	for i := 0; i < m; i++ {
		if A[i] != 0 {
			for j := 0; j < n; j++ {
				bufLen[i+j] += (A[i] * B[j])
			}
		}
	}
	for i := 0; i < len(*Result); i++ {
		(*Result)[i] = bufLen[i]
	}
}

//Polynom возвращает координату точки с известным х (Xpoint)
//на основе слайсов с координатами остальных точек
//ПРОТЕСТИРОВАНО
func Polynom(x []int, y []int, Xpoint int) float64 {
	var PolynomSum float64
	var buf float64
	//var resultStr string

	PolynomSum += float64(y[0])
	//resultStr += string(strconv.Itoa(y[0]))
	for i := 1; i < (len(x)); i++ {
		buf = Fa(i, 0, x, y)
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

//PolynomialCoefficents ПЕРЕСМОТРЕТЬ РАБОТУ С FLOAT
//Будет возвращать коэффициенты вычисленного полинома
//на основе слайсов с координатами известных точек
func PolynomialCoefficents(Result *[]float64, x []int, y []int) {
	var AnCoefficient float64

	(*Result)[0] = float64(y[0])

	for i := 1; i < (len(x)); i++ {

		polLen := len(x)
		PolBuf := make([]int, polLen)

		AnCoefficient = Fa(i, 0, x, y)
		fmt.Println("Fa^", i, " - ", AnCoefficient)

		//Pol member = Fn(x-n1)(x-n2)...(x-n-1)
		for j := 0; j < i; j++ {
			SecondMul := make([]int, 2)
			SecondMul[0] = -x[j]
			SecondMul[1] = 1
			if j == 0 {
				PolBuf[0] += -x[j]
				PolBuf[1]++
			} else {
				Multiply(&PolBuf, PolBuf, SecondMul, len(PolBuf), len(SecondMul))
			}
		}
		fmt.Println("PolBuf - ", PolBuf)
		for i := 0; i < polLen; i++ {
			//if coefficient ==0 => stop?
			(*Result)[i] += float64(PolBuf[i]) * AnCoefficient
		}
		//fmt.Println("Result ", i, " = ", (*Result))
	}
}

//Lagrange ??
func Lagrange(x []int, y []int, Xpoint int) int {
	Sum := 0
	PartSum := 0

	for i := 0; i < len(x); i++ {
		PartSum = y[i]
		for n := 0; n < len(x); n++ {
			if x[n] != Xpoint {
				PartSum *= (Xpoint - x[n])
			}
		}
		for m := 0; m < len(x); m++ {
			if i != m {
				PartSum /= (x[i] - x[m])
			}
		}
		Sum += PartSum
	}
	return Sum
}
