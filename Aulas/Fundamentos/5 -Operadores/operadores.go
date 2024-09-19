package main

import "fmt"

func main() {

	//ARITMÉTICOS
	soma := 1 + 1
	subtracao := 1 - 1
	divisao := 1 / 1
	multiplicacao := 1 * 1
	restoDaDivisao := 1 % 1

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	//RELACIONAIS - ( < | > | <= | >= | != )
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//LÓGICOS - ( && | || | ! )
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//UNÁRIOS ( ++ | -- | * | / | %)
	numero := 10
	numero++
	fmt.Println(numero)

}
