package main

import "fmt"

func main() {
	/*Değişken tanımının sytanx
	<var anahtar kelimesi> <değişken ismi> <static tipi> */

	/* var name string
	name = "Hasan Basri"

	var  firstName, lastName string = "Hasan Basri", "DARGA"

	var age int
	age = 23

	var isMarried bool
	isMarried = false



	fmt.Println(firstName,lastName)
	fmt.Println(age)
	fmt.Println(isMarried) */

	name := "Arin"
	age := 40
	isMarried := true

	// çoklu değişken oluşturma

	var (
		myName      string = "Hasan"
		myAge       int    = 22
		myIsMarried bool   = false

		myWeight float32 = 68.5
		myHeight int     = 178
	)

	//var youName, youAge, youIsMarried, youWeight, youHeight = "Arin", 40, true, 72.5, 172

	youName, youAge, youIsMarried, youWeight, youHeight := "Arin", 40, true, 72.5, 172

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)

	fmt.Println("-----------------------")
	fmt.Println(myName)
	fmt.Println(myAge)
	fmt.Println(myIsMarried)
	fmt.Println(myHeight)
	fmt.Println(myWeight)

	fmt.Println("-----------------------")
	fmt.Println(youName)
	fmt.Println(youAge)
	fmt.Println(youIsMarried)
	fmt.Println(youHeight)
	fmt.Println(youWeight)

	fmt.Println("-----------------------")
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isMarried)

	fmt.Println("-----------------------")

	/* Zero Values --- Değişkenlerin Varsayılan Değerleri

	string --->  "" boş değer döndürür

	numeric ---> 0 döndürür

	bool ----> false
	*/

	var temp string
	fmt.Println(temp)

	var temp2 int
	fmt.Println(temp2)

	var temp3 bool
	fmt.Println(temp3)

}
