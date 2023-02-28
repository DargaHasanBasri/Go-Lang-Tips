package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// Type Conversion type(value)  ---> int(y)

	//fmt.Println(x + y) // hatalı işlem çünkü iki değişkenin tipleri farklı oldğundan işlem yapılamaz

	fmt.Println(x + int(y)) // geçici olarak veri tipi değişti

	fmt.Printf("%v %T\n", y, y) // asıl değişkenimizin tipi değişmiyor y nin tipi float64

	// Bu tip durumlarda daha geniş bir kümeyi kapsayan tipe dönüştürülür
	z := 10
	t := 10.8

	fmt.Printf("%v %T\n", z, z)
	fmt.Printf("%v %T\n", t, t)

	fmt.Println(float64(z) + t)

	// string bir ifade int bir ifadeye dönüştürülemez

	/* 	m := 10
	   	n := "10"

	   	fmt.Println(m + int(n)) */

	// ASCII tabloda bazı ondalık sayılar karakterlerler ifade etmektedir
	// Bunu kullanarak int değeri string tipine dönüştürebiliriz

	fmt.Println()

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)
}
