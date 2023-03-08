package main

import "fmt"

func main() {

	/* // veri tipi belirli değil

	const x = 5

	// go varsayılan bir veri tipi atıyor

	fmt.Printf("%T, %v\n", x, x) */

	/* 	const x int8 = 3
	   	var y int16 = 12

	   	fmt.Printf("%T, %v\n", x, x)
	   	fmt.Printf("%T, %v\n", y, y)

	   	// veri tipleri aynı değil hata verir
	   	fmt.Printf("%T, %v\n", x+y, x+y) */

	// typeless const yani veri tipi belirsiz
	/* const x = 3
	var y int16 = 12

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)

	// hata almamak için varsayılan olarak int16(x) + y yapıyor
	fmt.Printf("%T, %v\n", x+y, x+y)

	// işi olmayınca typeless durumda varsayılan değer atıyor
	// beklediğimiz int16 gelmez
	fmt.Printf("%T, %v\n", x, x) */

	const z = 3
	const t = 5.6

	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

	// normalde tipleri farklı olan değişkenlerle işlem yapamıyoruz
	// fakat sabitlerde typeless durumunda hata almamak için
	// varsayılan tipleri go verir
	fmt.Printf("%T, %v", z+t, z+t)

}
