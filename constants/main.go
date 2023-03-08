package main

import (
	"fmt"
	"math"
)

func main() {

	/* 	r := 3.0
	   	const pi float64 = 3.14

	   	areaOfCircle := pi * (math.Pow(r, 2))

	   	fmt.Println(areaOfCircle) */

	/* const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)
	*/

	// const lar compile time durumunda oluşur
	// değişkenler run time durumunda oluşturulur

	var x = math.Pow(3, 4)
	fmt.Println(x)

	/* const y = math.Pow(3, 4)
	// sabit y compile time durumunda oluşurken
	// Pow işlemi run time durumunda oluşur bu yüzden hata alırız
	fmt.Println(y) */

	// const oluşturduğumuz ilklendirmemiz gereklidir ilk değeri olmalı

	/* 	y := 3

	   	const x = y // hata alırız biri compile time da oluşur x
	   	// diğeri run time da oluşur y */

}
