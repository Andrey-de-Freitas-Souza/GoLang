/*- Crie um tipo. O tipo subjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
- Na função main:
    1. Demonstre o valor da variável "x"
    2. Demonstre o tipo da variável "x"
    3. Atribua 42 à variável "x" utilizando o operador "="
    4. Demonstre o valor da variável "x"
    - Para os aventureiros: https://golang.org/ref/spec#Types */


package main

import (
	"fmt"
)
type mytype int 
var x mytype 


func main(){
	fmt.Printf("Variable X = Valor %v Type %T \n", x, x)
    x = 42
    fmt.Printf("Variable X = Valor %v Type %T \n", x, x)
}