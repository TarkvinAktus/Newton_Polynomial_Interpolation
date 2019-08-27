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
//На вход i - номер коэффициента; div_depth - шаг деления
//x - иксы, fx - игреки
func fa(i int, div_depth int, x []int, fx []int) float64 {
	var result float64
	var div float64
	if i == 1 {
		result = float64( fx[i+div_depth] - fx[i+div_depth-1] )
		div = float64(x[i+div_depth] - x[i+div_depth-1])
		if div != 0 {
			result /= div
		} else {
			result = 0
		}
	} else {
		result = (fa(i-1,div_depth+1,x,fx) - fa(i-1,div_depth,x,fx) )
		div= float64(x[i+div_depth] - x[0 + div_depth])
		if div != 0 {
			result /= div
			result = toFixed(result,2)
		} else {
			result = 0
		}
	}
	return result
}


func polynom(x []int,y []int, x_point int) float64{
	var Polynom_sum float64
	var buf float64

	Polynom_sum += float64(y[0])

	for i := 1;i<(len(x)); i++ {
		buf = fa(i,0,x,y) 
		for j := 0;j<i; j++ {
			buf *= float64(x_point-x[j]) 
		}
		Polynom_sum += buf
	}
	return math.Round(Polynom_sum)
}


func check(x int) float64{
	var buf float64
	buf = (float64(-17)/float64(6))*(math.Pow(float64(x),3)) + (math.Pow(float64(x),2)*12.5) - (float64(38)/float64(3) * float64(x)) - 2
	return math.Round(buf) 

}

func main() {

	x := []int{0,1,2,3}
	fx := []int{-2,-5,0,-4}
/*
	fmt.Println("fa(10) = ",fa(1,0,x,fx))
	fmt.Println("fa(11) = ",fa(1,1,x,fx))
	fmt.Println("fa(12) = ",fa(1,2,x,fx))
	fmt.Println("fa(20) = ",fa(2,0,x,fx))
	fmt.Println("fa(21) = ",fa(2,1,x,fx))
	fmt.Println("fa(30) = ",fa(3,0,x,fx))
*/

	
	fmt.Println("___check 0 ",check(0))
	fmt.Println("___check 1 = ",check(1))
	fmt.Println("___check 2 = ",check(2))
	fmt.Println("___check 3 = ",check(3))
	fmt.Println("___check 3 = ",check(4))
	
	fmt.Println("polynom 0 = ",polynom(x,fx,0))
	fmt.Println("polynom 1 = ",polynom(x,fx,1))
	fmt.Println("polynom 2 = ",polynom(x,fx,2))
	fmt.Println("polynom 3 = ",polynom(x,fx,3))
	fmt.Println("polynom 4 = ",polynom(x,fx,4))

	var scan string
	fmt.Scan(&scan)
}