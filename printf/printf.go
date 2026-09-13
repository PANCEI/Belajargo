package main

import "fmt"

type Karyawan struct {
	Nama string
	Usia int
}

func main() {
	nama := "Budi"
	umur := 25
	gaji := 7500000.50
	isAktif := true
	k := Karyawan{Nama: "Siti", Usia: 28}

	fmt.Println("=== DEMO FORMAT PRINTF DI GO ===")

	// 1. %v (Value): Mencetak nilai variabel dalam bentuk standar
	fmt.Printf("1. %%v  : %v\n", nama)

	// 2. %+v (Plus Value): Sangat berguna untuk Debugging Struct!
	// Mencetak nama field beserta nilainya
	fmt.Printf("2. %%+v : %+v\n", k)

	// 3. %#v (Go Syntax Value): Mencetak variabel sesuai sintaks kode Go
	fmt.Printf("3. %%#v : %#v\n", k)

	// 4. %T (Type): Mencetak TIPE DATA dari variabel (Sangat penting untuk Debugging)
	fmt.Printf("4. %%T  : Tipe data 'gaji' adalah %T\n", gaji)

	// 5. %d, %f, %t, %s (Tipe Spesifik)
	fmt.Printf("5. %%d  : Angka bulat = %d\n", umur)
	fmt.Printf("   %%.2f: Angka desimal = Rp %.2f\n", gaji)
	fmt.Printf("   %%t  : Boolean = %t\n", isAktif)

	// 6. %p (Pointer): Mencetak alamat memori (Debugging Pointer)
	fmt.Printf("6. %%p  : Alamat memori 'k' = %p\n", &k)
}
