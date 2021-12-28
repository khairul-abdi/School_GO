package main

import "fmt"

func main() {

	//Ini komentar dengan 1 baris saja
	/*
		ini komentar dengan multiple baris
	*/

	//Contoh 1 Declare variable
	var name string   //ini istilahnya declare variable
	name = "khairul"  //istilahnya initial variable
	fmt.Println(name) //menampilkan isi variable

	var umur int
	umur = 26
	fmt.Println(umur)
	fmt.Println("---------------------------")
	//Contoh 2 Declare variable
	var name2 string = "khairul"
	fmt.Println(name2)
	fmt.Println("---------------------------")

	//Contoh 3 Declare variable
	name3 := "zulhan" //ini paling sering kugunakan
	fmt.Println(name3)
	fmt.Println("---------------------------")

	//Contoh 4: Declare multiple variable
	var x, y, z string
	x, y, z = "ini x", "ini y", "ini z"
	fmt.Println(x, y, z)

	nama, old, _ := "jhon", 23, "jl.cambridge"

	fmt.Println(nama)
	fmt.Println(old)
	fmt.Println("---------------------------")

	//Contoh 5: Declare menggunakan new
	name01 := new(string)
	fmt.Println("declare new: ", name01)  // 0x20818a220
	fmt.Println("declare new: ", *name01) // ""
	fmt.Println("---------------------------")

}
