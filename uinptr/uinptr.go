package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	Angka int64 // Ukuran 8 byte
	Teks  string
}

func main() {
	d := Data{
		Angka: 100,
		Teks:  "Golang",
	}

	fmt.Println("=== 1. Mengakses Field Struct Menggunakan unsafe.Add ===")
	// Ambil pointer dasar struct 'd'
	pData := unsafe.Pointer(&d)

	// Hitung offset memori untuk field 'Teks'
	offsetTeks := unsafe.Offsetof(d.Teks)

	// POINTER ARITHMETIC AMAN:
	// unsafe.Add langsung menambah offset (byte) ke unsafe.Pointer
	pTeks := (*string)(unsafe.Add(pData, offsetTeks))

	fmt.Println("Nilai field 'Teks' via unsafe.Add :", *pTeks)

	// Ubah nilai field 'Teks' secara langsung di memori
	*pTeks = "Master Go"
	fmt.Println("Nilai 'd.Teks' setelah diubah    :", d.Teks)

	fmt.Println("\n=== 2. Akses Array dengan Pointer Arithmetic ===")
	arr := [3]int32{10, 20, 30} // int32 berukuran 4 byte

	// Ambil alamat elemen pertama (indeks 0) sebagai unsafe.Pointer
	pElem0 := unsafe.Pointer(&arr[0])

	// Pindah ke elemen kedua (indeks 1) menggunakan unsafe.Add
	// ukuran 1 elemen int32 = unsafe.Sizeof(arr[0])
	pElem1 := (*int32)(unsafe.Add(pElem0, unsafe.Sizeof(arr[0])))

	fmt.Println("Nilai arr[0]                     :", arr[0])
	fmt.Println("Nilai arr[1] via unsafe.Add      :", *pElem1)
}
