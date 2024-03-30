/* - Escreva expressões utilizando os operadores, e atribua seus valores a variáveis.
	- ==
	- !=
	- <
	- >
	- <=
	- >=
- Demonstre estes valores.*/
package main

import(
	"fmt"
)

func main(){
    value1 := 100
	value2 := 150
	a := (value1 == value2)
	b := (value1 != value2)
	c := (value1 < value2)
	d := (value1 > value2)
	e := (value1 <= value2)
	f := (value1 >= value2)
	fmt.Printf("Value1 = %v \nValue2 = %v \nvalue1 == value2 = %v \nvalue1 != value2 = %v \nvalue1 < value2 = %v \nvalue1 > value2 = %v \nvalue1 <= value2 = %v \nvalue1 >= value2 = %v \n",value1,value2, a,b,c,d,e,f)
}