package main

import "fmt"

var packVar = "Package Scope"

var funcVar = "Func(Package) Scope"

// PackVar2 := "Package Scope2"  // kısa gösterim package değişkenlerinde kullanılmaz hata verir

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar) // kendi scope içinde olan değişkeni yazıdırıyor
	}

	// fmt.Println(blockVar) // kendi scope içinde olmadığı için yazdıramaz hata verir

	funcVar := "Func Scope" // kısa gösterim kullanılabiliyor

	fmt.Println(funcVar) // kendi scope içinde olan değişkeni yazdırıyor
	// main içindeki değişkeni yazdırır aynı isimde package değişkeni olsa bile

	fmt.Println(packVar) // kendi scope içinde olmasada kendi package içindeki değişkeni yazdırıyor

}

func myFunc() {
	fmt.Println(packVar) // package ye ait değişkenlere her yerden ulaşabiliyorum
	// fmt.Println(funcVar)  // kendi scope dışından değişkene ulaşamayız hata verir
}

func myFunc2() {
	fmt.Println(funcVar) // main içindekini değil package içindeki değişkeni yazdırır

}
