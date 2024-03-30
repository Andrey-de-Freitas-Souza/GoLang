/*- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
*/
package main

import (
	"fmt"
)


func main() {
	x := 200
	fmt.Printf("Variable X = %v \nDecimal = %d \nBinary = %b \nHexadecimal = %#x\n", x, x,x,x)
	y := x << 1
	fmt.Printf("\nVariable y = %v \nDecimal = %d \nBinary = %b \nHexadecimal = %#x", y, y,y,y)

}
