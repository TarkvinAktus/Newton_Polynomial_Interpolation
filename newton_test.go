package main

import (
	"math"
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
		//fmt.Println("result = ", result)

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

	if Fa(1, 0, x, fx) != FaExcepted[0] {
		t.Errorf("Fa(1,0)")
	}
	if Fa(1, 1, x, fx) != FaExcepted[1] {
		t.Errorf("Fa(1,1)")
	}
	if Fa(1, 2, x, fx) != FaExcepted[2] {
		t.Errorf("Fa(1,2)")
	}
	if Fa(2, 0, x, fx) != FaExcepted[3] {
		t.Errorf("Fa(2,0)")
	}
	if Fa(2, 1, x, fx) != FaExcepted[4] {
		t.Errorf("Fa(2,1)")
	}
	if Fa(3, 0, x, fx) != FaExcepted[5] {
		t.Errorf("Fa(3,0)")
	}

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
	//fmt.Println("")
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
