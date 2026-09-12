// true jika ekspresi terpenuhi
// false jika ekspresi tidak terpenuhiu

// AND logika  and  && a && b
// logika or or  || a || b
// logika not not  !a
package main

import "fmt"

func main() {
	// a := true
	// b := false
	// fmt.Println("a && b:", a && b)
	// fmt.Println("a || b:", a || b)
	// fmt.Println("!b:", !b)
	a := 10
	b := 14

	// Menyimpan hasil perbandingan ke dalam variabel boolean
	apakahALebihKecil := a < b     // bernilai true
	apakahBSamaDengan14 := b == 14 // bernilai true

	// 1. Logika AND (&&) - Kedua kondisi harus true
	// Contoh: Apakah (a < b) DAN (b > 20)?
	fmt.Println("a < b && b > 20  :", (a < b) && (b > 20)) // true && false -> Hasil: false

	// 2. Logika OR (||) - Salah satu kondisi bernilai true
	// Contoh: Apakah (a == 10) ATAU (b == 10)?
	fmt.Println("a == 10 || b == 10:", (a == 10) || (b == 10)) // true || false -> Hasil: true

	// 3. Logika NOT (!) - Membalikkan nilai boolean
	// Contoh: Membalikkan nilai dari (a < b)
	fmt.Println("!(a < b)         :", !(a < b)) // !(true) -> Hasil: false

	// Penggunaan variabel boolean langsung pada operator logika
	fmt.Println("Gabungan variabel:", apakahALebihKecil && apakahBSamaDengan14) // true && true -> Hasil: true
}
