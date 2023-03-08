package main

import "fmt"

func main() {
	/* x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", (x + y), (x + y))
	fmt.Printf("%T, %v\n", (x - y), (x - y))
	fmt.Printf("%T, %v\n", (x * y), (x * y))
	fmt.Printf("%T, %v\n", (x / y), (x / y)) */

	/* z := 5.0 / 2
	fmt.Printf("%T, %v\n", z, z) */

	/* t, v := 15, 9
	fmt.Printf("%T, %v\n", (t % v), (t % v))  // kalan verir */

	// Increment ++, Decrement --    POSTFIX VARDIR SADECE SAĞINDA ++ -- OLUR

	x := 10

	fmt.Println(x)

	x = x + 1

	fmt.Println(x)

	x++ // burası !!!

	// x = ++x  // go dilinde böyle bir durum  yoktur

	fmt.Println(x)

	// fmt.Println(x++)  // bir satırda yalnızca bir statements bulunabilir

	// burda println de statements dır  bunun çözümü

	// !!! işareti kullanarak yapabiliriz

	y := 8

	y = y - 1

	fmt.Println(y)

	y--

	fmt.Println(y)

}
