package main

import "fmt"

// 1. FUNGSI SEDERHANA (Tanpa Parameter & Tanpa Return)
func sapa() {
	fmt.Println("1. Halo, Selamat Belajar Golang!")
}

// 2. FUNGSI DENGAN PARAMETER & RETURN VALUE TUNGGAL
func tambah(a int, b int) int {
	return a + b
}

// 3. FUNGSI DENGAN MULTIPLE RETURN VALUES (Banyak Nilai Kembalian)
// Mengembalikan 2 nilai: hasil pembagian (int) dan sisa bagi (int)
func bagiDanSisa(angka int, pembagi int) (int, int) {
	hasil := angka / pembagi
	sisa := angka % pembagi
	return hasil, sisa
}

// 4. FUNGSI DENGAN NAMED RETURN VALUES
// Nama variabel kembalian ('luas' dan 'keliling') dideklarasikan di awal
func hitungPersegi(sisi int) (luas int, keliling int) {
	luas = sisi * sisi
	keliling = 4 * sisi
	return // Otomatis mengembalikan nilai 'luas' dan 'keliling'
}

// 5. VARIADIC FUNCTION (Parameter Bertipe Slice/Banyak Data)
// Tanda (...) membuat fungsi bisa menerima jumlah angka berapa saja
func totalJumlah(angka ...int) int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}

func main() {
	// Memanggil Fungsi 1
	sapa()

	// Memanggil Fungsi 2
	hasilTambah := tambah(10, 5)
	fmt.Println("2. Hasil Tambah (10 + 5)     :", hasilTambah)

	// Memanggil Fungsi 3
	hasilBagi, sisaBagi := bagiDanSisa(10, 3)
	fmt.Printf("3. Hasil Bagi 10/3           : %d (Sisa: %d)\n", hasilBagi, sisaBagi)

	// Trik: Gunakan '_' jika ada nilai kembalian yang tidak ingin digunakan
	hBagi, _ := bagiDanSisa(20, 4)
	fmt.Println("   Hasil Bagi 20/4 (sisa diabaikan):", hBagi)

	// Memanggil Fungsi 4
	l, k := hitungPersegi(5)
	fmt.Printf("4. Persegi (sisi 5)          : Luas = %d, Keliling = %d\n", l, k)

	// Memanggil Fungsi 5 (Variadic)
	total1 := totalJumlah(1, 2, 3, 4, 5) // Menerima 5 angka
	total2 := totalJumlah(10, 20)        // Menerima 2 angka
	fmt.Println("5. Total Jumlah (1+2+3+4+5)  :", total1)
	fmt.Println("   Total Jumlah (10+20)      :", total2)

	// 6. ANONYMOUS FUNCTION & CLOSURE (Fungsi Tanpa Nama)
	// Fungsi disimpan langsung ke dalam variabel 'perkalian'
	perkalian := func(x, y int) int {
		return x * y
	}
	fmt.Println("6. Anonymous Func (4 * 3)    :", perkalian(4, 3))
}
