package main

//Посроение интерполяционного многочлена Ньютона
import (
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
		div = float64(x[i+divDepth] - x[i+divDepth-1])
		if div != 0 {
			result /= div
		} else {
			result = 0
		}
	} else {
		result = (Fa(i-1, divDepth+1, x, fx) - Fa(i-1, divDepth, x, fx))
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

// Multiply ...
// Просто поочерёдно умножаем члены каждого многочлена
// Результат умножения в степень равную сумме степеней членов многочлена
// Внутри слайсов они отсортированы по степеням*
func Multiply(Result *[]int, A []int, B []int, m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			(*Result)[i+j] += A[i] * B[j]
			//fmt.Print("+", A[i]*B[j], "x^", i+j)
		}
		//fmt.Println("")
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

//PolynomialCoefficents В РАБОТЕ
//Комментарии
//ADD RETURN
//Будет возвращать коэффициенты вычисленного полинома
//на основе слайсов с координатами известных точек
func PolynomialCoefficents(x []int, y []int) {
	PolCoefficents := make([]float64, len(x))
	var AnCoefficient float64

	PolCoefficents[0] = float64(y[0])

	for i := 1; i < (len(x)); i++ {

		AnCoefficient = Fa(i, 0, x, y)

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
				Multiply(&MulBuf, FirstMul, SecondMul, len(FirstMul), len(SecondMul))
			} else {
				Multiply(&MulBuf, MulBuf, SecondMul, len(MulBuf), len(SecondMul))
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
