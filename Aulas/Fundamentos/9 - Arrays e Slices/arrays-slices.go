package main

import "fmt"

func main() {
	var a [5]int
	a[2] = 7
	println(a[2])

	b := [5]int{1, 2, 3, 4, 5}
	println(b[2])

	var c []int
	c = append(c, 1)
	c = append(c, 2)
	println(c[1])

	d := []int{1, 2, 3, 4, 5}
	println(d[3])

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := d[1:3]
	fmt.Println(slice2)

	slice2[1] = 100
	fmt.Println(slice2)

	slice3 := make([]float32, 10, 15)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 10)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
