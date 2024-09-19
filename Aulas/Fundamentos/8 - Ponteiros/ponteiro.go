package main

import "fmt"

func main() {
	var1 := 10
	var2 := var1

	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	var3 := 100
	ponteiro := &var3
	fmt.Println(var3, *ponteiro)
	var3 = 150
	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)

}
