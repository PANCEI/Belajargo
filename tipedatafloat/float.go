package main

import "fmt"

func main() {
	// float literal
	//0.1 , 3.14 , 2.2,0.5
	//float 32 dan float 64
	//float 32 memiliki renge 1.18e-38 hingga 3.4e38
	//float 64 memiliki range 2.23e-308 hingga 1.8e308
	// nilai float 64 memiliki presisi 15-16 digit
	// nilai float 32 memiliki presisi 6-7 digit
	// var pi float32 = 3.14
	// var e float64 = 2.718281828459045
	// fmt.Println("nilai pi adalah ", pi)
	// fmt.Println("nilai e adalah ", e)
	//1.23e10 -> 1.23 * 10^10
	var angka_besar float64 = 1.23e10
	//1.23e-10 -> 1.23 * 10^-10
	var angka_kecil float32 = 1.23e-10
	fmt.Println("nilai angka besar adalah ", angka_besar)
	fmt.Println("nilai angka kecil adalah ", angka_kecil)
}
