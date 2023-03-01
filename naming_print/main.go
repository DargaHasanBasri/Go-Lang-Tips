package main

import "fmt"

func main() {

	/* // (Print - Println)  --> olduğu gibi yazdırır
	// Printf ---> formatlanmış şekilde yazdırır

	fmt.Println("Merhaba")

	//Println işleminden sonra bir alt satıra geçer

	fmt.Print("Merhaba")

	fmt.Println()

	fmt.Printf("Merhaba")
	fmt.Println()

	fmt.Println("-----------------------------")

	name := "Hasan"

	fmt.Print("Benim Adım", name) // Benim AdımHasan şeklinde yazdırır
	fmt.Println()
	fmt.Println("Benim Adım", name) // Benim Adım Hasan şeklinde yazdırır

	// fmt.Printf("Benim Adım", name) // Formatlanmış olduğu için (EXTRA string hatası verir)

	// hatayı almamak için özel karakter kullanılır
	// %v değikenin değerini döndürür
	fmt.Printf("Benim Adım %v", name) // Benim Adım Hasan şeklinde yazdırır
	fmt.Println()

	// %T değişkenin type döndürür
	fmt.Printf("Benim Adım %T", name) // Benin Adım string şeklinde yazdırır
	fmt.Println()

	fmt.Printf("Benim Adım %v %T", name, name) // Benim Adım Hasan string
	fmt.Println()

	// %X 16lık değerini döndürür
	fmt.Printf("Benim Adım %v %X", name, name) */

	x := 100
	y := 20
	z := 30

	fmt.Printf("%b", x) // 2lik sisteme göre değerini döndürür
	fmt.Println()

	fmt.Printf("%d", y) //10luk sisteme göre değerini döndürür
	fmt.Println()

	fmt.Printf("%o", z) // 8lik sisteme göre değerini döndürür
	fmt.Println()

	// Genel Kültür

	// Go türkçe karakterlerle oluşan değişkenlerde hata vermez
	// Çince Japoncayı da destekliyor

	a := 5

	if true {
		a := 10        // := demek var olan değişkeni yeniden oluşturuyor
		fmt.Println(a) // 10 yazdırıcak
	}

	fmt.Println(a) // 5 yazdırıcak

	b := 6

	if true {
		b = 12         // = var olan değişkene yeni değer atıyor
		fmt.Println(b) // 12
	}

	fmt.Println(b) // yeni değeri 12 yazdırır

}
