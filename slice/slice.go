package main

import "fmt"

func main() {
	// ==========================================
	// 1. DEKLARASI SLICE LANGSUNG (Literal)
	// ==========================================
	// Berbeda dengan Array, kita TIDAK menentukan ukurannya di dalam kurung siku []
	buah := []string{"Apel", "Jeruk", "Mangga"}
	fmt.Println("1. Slice Buah Awal     :", buah)
	fmt.Printf("   Panjang (len): %d, Kapasitas (cap): %d\n", len(buah), cap(buah))

	// ==========================================
	// 2. MENAMBAH ELEMEN DENGAN append()
	// ==========================================
	// Ukuran slice akan bertambah secara otomatis
	buah = append(buah, "Pisang", "Melon")
	fmt.Println("\n2. Setelah append()     :", buah)
	fmt.Printf("   Panjang (len): %d, Kapasitas (cap): %d\n", len(buah), cap(buah))

	// ==========================================
	// 3. MEMBUAT SLICE DENGAN FUNGSI make()
	// ==========================================
	// Format: make([]TipeData, length, capacity)
	angka := make([]int, 3, 5) // Panjang awal 3, kapasitas maksimum 5
	angka[0] = 10
	angka[1] = 20
	angka[2] = 30
	fmt.Println("\n3. Slice dari make()   :", angka)
	fmt.Printf("   Panjang (len): %d, Kapasitas (cap): %d\n", len(angka), cap(angka))

	// ==========================================
	// 4. PEMOTONGAN (SLICING) DARI ARRAY/SLICE LAIN
	// ==========================================
	// Format: data[indeks_awal : indeks_akhir_dikecualikan]
	hewan := [5]string{"Kucing", "Anjing", "Kelinci", "Burung", "Ikan"}

	sliceHewan1 := hewan[1:4] // Mengambil indeks 1, 2, dan 3 ("Anjing", "Kelinci", "Burung")
	fmt.Println("\n4. Potongan hewan[1:4]  :", sliceHewan1)

	// Sifat Referensi: Jika elemen Slice diubah, Array aslinya JUGA ikut berubah!
	sliceHewan1[0] = "Serigala"
	fmt.Println("   Array Asli Terdampak :", hewan) // "Anjing" berubah jadi "Serigala"

	// ==========================================
	// 5. MENGODOPY SLICE DENGAN copy()
	// ==========================================
	// copy() menyalin elemen agar tidak saling memengaruhi memori (Pass by Value)
	asal := []int{1, 2, 3}
	tujuan := make([]int, len(asal)) // Harus disiapkan ukurannya dulu
	copy(tujuan, asal)

	tujuan[0] = 99                                   // Mengubah slice tujuan
	fmt.Println("\n5. Slice Asal           :", asal) // Tetap [1 2 3]
	fmt.Println("   Slice Tujuan Copy    :", tujuan) // Jadi [99 2 3]
}
