
package main

import (
	"fmt"
)

func main() {
	i := 9
	switch{
		
	case i > 10:
		fmt.Println("I é maior que 10")
	case i == 10:
		fmt.Println("I é igual a 10")
	case i < 10:
		fmt.Println("I é menor que 10")

	}
}
