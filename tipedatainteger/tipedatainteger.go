package main

import "fmt"

func main() {
	//tipe data integer adalah tipe data yang digunakan untuk menyimpan bilangan bulat
	//tipe bilangan bulat 10, 100, 2000
	// 32bit ->int 4 byte -> 32 bit
	//64bit -> int 8 byte ->64 bit
	// tipe data integer di golang adalah tipe signed integer (bilangan bulat bertanda) dan unsigned integer (bilangan bulat tidak bertanda)
	/*
			Tipe data signed integer (bilangan bulat bertanda) adalah tipe data integer yang bisa menyimpan nilai negatif, nol, dan positif.

		Kata "signed" mengacu pada ketersediaan tanda plus (+) atau minus (-) pada angka tersebut. Untuk menyimpan tanda negatif, komputer menggunakan 1 bit (yaitu bit paling kiri/paling depan) sebagai penanda (+ atau -).

		Perbedaan Signed vs Unsigned
		Signed Integer: Mendukung angka negatif dan positif.

		Unsigned Integer: Hanya mendukung angka nol dan positif (tidak bisa menyimpan angka negatif).
	*/
	//contoh data tipe data unsigned adalah
	// uint8 , uint16 , uint32 , uint64
	//contoh data tipe data signed adalah
	// int8 , int16 , int32 , int64
	// var angka int = 43
	// fmt.Println("nilai angka adalah ", angka)
	// var angka2 int
	// fmt.Println("nilai angka2 adalah ", angka2)
	angka := 43
	fmt.Println("nilai angka adalah ", angka)

}
