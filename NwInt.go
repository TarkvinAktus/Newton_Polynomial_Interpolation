package main
//Посроение интерполяционного многочлена Ньютона
import (
	"fmt"
	"github.com/shopspring/decimal"
)


//Доделать Decimal

//Функция f(a) - коэффициентов полинома
//Составляется из разделённых разностей полинома (см теорию http://matlab.exponenta.ru/spline/book1/14.php)
//fa(Xi,Xi+1) = f(Xi+1)-f(Xi) / Xi+1 - Xi
//fa(X1,X2,X3) = fa(X2,X3) - fa(X1,X2) / X3 - X1
//И т.д.
//На вход i - номер коэффициента; div_depth - шаг деления
//x - иксы, fx - игреки
func fa(i int, div_depth int, x []int, fx []int) int {
	var result int
	var div int
	var N_D decimal.New(0,0)
	//N_D := decimal.New(0,0)
	N_D2 := decimal.New(0,0)
	if i == 1 {
		result = ( fx[i+div_depth] - fx[i+div_depth-1] )
		div = (x[i+div_depth] - x[i+div_depth-1])
		if div != 0 {
			result /= div
		} else {
			result = 0
		}
	} else {
		result = (fa(i-1,div_depth+1,x,fx) - fa(i-1,div_depth,x,fx) )
		div= (x[i+div_depth] - x[0 + div_depth])
		if div != 0 {
			N_D = decimal.New(int64(result),0)
			N_D2 = decimal.New(int64(div),0)
			N_D = N_D.Div(N_D2)
			fmt.Println(N_D.StringFixed(2))
			result /= div
		} else {
			result = 0
		}
		

		
	}
	return result
	
}

func main() {

	x := []int{0,1,2,3}
	fx := []int{-2,-5,0,-4}
	/*
	x := []int{0,1,2,3}
	fx := []int{-2,-5,0,-4}

	var Polynom_sum int
	var buf int

	Polynom_sum += fx[0]

	for i := 1;i<len(x); i++ {
		if x_point == x[i]
			break
		buf = fa(i,0) 
		for j := 0;j<i; j++ {
			buf *= (x_point-x[j]) 
		}
		Polynom_sum += buf
	}*/

	fmt.Println("fa(21) = ",fa(3,0,x,fx))

}