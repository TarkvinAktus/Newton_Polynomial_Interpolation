package main

import (
	"math"
	"reflect"
	"testing"
)

func Test_PolynomMultiply(t *testing.T) {
	a := [][]int{
		{-3, 1, -2},
		{-3, 1, -2},
		{6, -1}}
	b := [][]int{
		{2, -1},
		{1, -2, 4},
		{-5, 2, -3, 2}}

	answ := [][]int{
		{-6, 5, -5, 2},
		{-3, 7, -16, 8, -8},
		{-30, 17, -20, 15, -2}}

	for i := 0; i < len(a); i++ {
		m := len(a[i])
		n := len(b[i])
		result := make([]int, n+m-1)

		Multiply(&result, a[i], b[i], m, n)

		for j := 0; j < len(result); j++ {
			if result[j] != answ[i][j] {
				t.Errorf("Incorrect")
			}
		}
	}
}

func Test_faCoefficents(t *testing.T) {
	//Источник - http://kontromat.ru/?page_id=4955
	x := []int{0, 1, 2, 3}
	fx := []int{-2, -5, 0, -4}
	FaExcepted := []float64{-3, 5, -4, 4, -4.5, -2.83}

	//Round to 2 signs because of giant floats
	if toFixed(Fa(1, 0, x, fx), 2) != FaExcepted[0] {
		t.Errorf("Fa(1,0)")
	}
	if toFixed(Fa(1, 1, x, fx), 2) != FaExcepted[1] {
		t.Errorf("Fa(1,1)")
	}
	if toFixed(Fa(1, 2, x, fx), 2) != FaExcepted[2] {
		t.Errorf("Fa(1,2)")
	}
	if toFixed(Fa(2, 0, x, fx), 2) != FaExcepted[3] {
		t.Errorf("Fa(2,0)")
	}
	if toFixed(Fa(2, 1, x, fx), 2) != FaExcepted[4] {
		t.Errorf("Fa(2,1)")
	}
	if toFixed(Fa(3, 0, x, fx), 2) != FaExcepted[5] {
		t.Errorf("Fa(3,0)")
	}
}

func Test_PolynomialCoefficents(t *testing.T) {

	x := []int{0, 1, 2, 3}
	fx := []int{-2, -5, 0, -4}

	Excepted := []float64{-2, -12.66, 12.49, 2.83}

	Result := make([]float64, len(x))

	PolynomialCoefficents(&Result, x, fx)
	//fmt.Println(Result)
	//not have to use toFixed float round
	if reflect.DeepEqual(Result, Excepted) {
		t.Errorf("Incorrect")
	}

	/*Here was testing polinom to compare its fn coefficients
	not about result coefficents

	x1 := []int{2, 5, 8, 10, 15}
	fx2 := []int{6, 3, 6, 3, 4}

	Result2 := make([]float64, len(x1))

	PolynomialCoefficents(&Result2, x1, fx2)

	if reflect.DeepEqual(Result, Excepted?) {
		t.Errorf("Incorrect")
	}
	fmt.Println(Result2)
	*/
}

func Test_fxPolynom(t *testing.T) {
	//Источник - http://kontromat.ru/?page_id=4955
	x := []int{0, 1, 2, 3}
	fx := []int{-2, -5, 0, -4}
	Xpoints := []int{0, 1, 2, 3, 4}
	expected := []float64{-2, -5, 0, -4, -34}

	answ := make([]float64, 5)

	for i := 0; i < 5; i++ {
		answ[i] = Polynom(x, fx, Xpoints[i])
		//fmt.Print(" | Polynom ", i, " = ", answ[i])
	}

	for i := 0; i < 5; i++ {
		//Вычисление тестового примера для
		//x := []int{0, 1, 2, 3}
		//fx := []int{-2, -5, 0, -4}
		testCalc := func(x int) float64 {
			var buf float64
			buf = (float64(-17)/float64(6))*(math.Pow(float64(x), 3)) + (math.Pow(float64(x), 2) * 12.5) - (float64(38) / float64(3) * float64(x)) - 2
			return math.Round(buf)
		}
		//Сравнивание с заданными и вычисленными значениями
		if answ[i] != expected[i] && answ[i] != testCalc(i) {
			t.Errorf("fxPolynom")
		}
	}
}
