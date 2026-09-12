package main

import "fmt"

func main() {
	fmt.Println(30)
	fmt.Println(12)
	// string
	fmt.Println("sadas sadsad as asdas")
	// backtick
	fmt.Println(`basris 1
	baris dua
	baris tiga
	`)
	// floating point , angka desimal
	fmt.Println(4.13)
	fmt.Println(-4.13)
	/*

		Nilai 'A' menjadi 65 karena dalam bahasa Go (Golang), karakter yang diapit oleh petik tunggal (' ') dianggap sebagai tipe data rune, bukan string.Secara internal di dalam sistem komputer:rune adalah alias untuk int32 (bilangan bulat 32-bit).Representasi Angka (Kode ASCII/Unicode): Komputer menyimpan karakter huruf sebagai kode angka. Karakter huruf kapital 'A' memiliki nilai kode ASCII / Unicode 65.Saat Anda memanggil fmt.Println('A'), Go tidak mencetak karakter fisiknya secara otomatis, melainkan mencetak nilai numerik (ASCII/Unicode) dari karakter tersebut.Perbandingan Karakter vs String di Go'A' (Petik tunggal) $\rightarrow$ Tipe rune (mencetak nilai angka: 65)"A" (Petik ganda) $\rightarrow$ Tipe string (mencetak teks: A)
	*/
	fmt.Println('A')
	// boool
	fmt.Println(false)
}
