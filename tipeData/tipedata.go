package main

import "fmt"

func main() {
	// yang mempenagruhi tipe data adalah
	// 1. ukuran memori
	//2. operasi
	// 3. cara data tersebut di interpretasikan

	// tipe data di golang
	// 1. tipe data numerik
	// int , int8 , int16 , int32 , int64 -> 1. -200, 2. 0 , 3. 200
	// uint , uint8 , uint16 , uint32 , uint64 -> 1. 0 , 2. 200
	// float32 , float64 -> 1. -200.12 , 2. 0.0 , 3. 200.12
	// complex64 , complex128 -> 1. -200.12+5i , 2. 0+0i , 3. 200.12+5i
	// byte -> alias dari uint8
	// rune -> alias dari int32
	// 2. tipe data string
	// 3. tipe data boolean
	// bool -> true , false
	// 4. tipe data array
	// 5. tipe data slice
	// 6. tipe data struct
	// 7. tipe data pointer
	// 8. tipe data function
	// 9. tipe data interface

	// cara mendeklasrikan variable di golang

	// var nama string = "baco"
	// // fmt.Println("nama saya adalah ", nama)
	// // var umur int = 20
	// // var sudahBesar bool
	// // var haji float32
	// // sudahBesar = true
	// // haji = 100.12
	// // fmt.Println("umur saya adalah ", umur)
	// // fmt.Println(sudahBesar)
	// // fmt.Println(haji)
	// name := "baco2"
	// fmt.Println("nama saya adalah ", nama)
	// fmt.Println("nama saya adalah ", name)
	// var nilai_pertama, nilai_kedua int = 10, 20
	// fmt.Println("nilai pertama adalah ", nilai_pertama)
	// fmt.Println("nilai kedua adalah ", nilai_kedua)
	x, y, z := 10.3, "sadasd", true
	fmt.Println("nilai x adalah ", x)
	fmt.Println("nilai y adalah ", y)
	fmt.Println("nilai z adalah ", z)
}
