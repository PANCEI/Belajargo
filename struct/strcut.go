package main

import "fmt"

// 1. DEKLARASI STRUCT STANDAR
type Alamat struct {
	Kota     string
	Provinsi string
}

// 2. EMBEDDED STRUCT (Struct di dalam Struct)
type Karyawan struct {
	ID     int
	Nama   string
	Gaji   float64
	Alamat Alamat // Mengambil bidang dari struct Alamat
}

// 3. STRUCT METHOD (Fungsi yang menempel pada Struct Karyawan)
// Receiver (k Karyawan) membuat fungsi ini bisa dipanggil langsung oleh instance Karyawan
func (k Karyawan) TampilkanInfo() {
	fmt.Printf("   [ID: %d] Nama: %s | Gaji: Rp %.0f | Lokasi: %s, %s\n",
		k.ID, k.Nama, k.Gaji, k.Alamat.Kota, k.Alamat.Provinsi)
}

// Method dengan Pointer Receiver (untuk mengubah isi struct)
func (k *Karyawan) NaikGaji(persen float64) {
	k.Gaji += k.Gaji * (persen / 100)
}

func main() {
	// ==========================================
	// 1. INISIALISASI STRUCT
	// ==========================================
	// Cara A: Menggunakan nama field (Sangat Direkomendasikan)
	karyawan1 := Karyawan{
		ID:   101,
		Nama: "Budi Santoso",
		Gaji: 7500000,
		Alamat: Alamat{
			Kota:     "Jakarta Selatan",
			Provinsi: "DKI Jakarta",
		},
	}

	// Cara B: Posisional (harus urut sesuai urutan field di struct)
	karyawan2 := Karyawan{102, "Siti Rahma", 8500000, Alamat{"Bandung", "Jawa Barat"}}

	fmt.Println("1. Memanggil Method Struct:")
	karyawan1.TampilkanInfo()
	karyawan2.TampilkanInfo()

	// ==========================================
	// 2. MENGUBAH VALUE STRUCT VIA METHOD (POINTER)
	// ==========================================
	fmt.Println("\n2. Mengubah Data via Pointer Method (Naik Gaji 10%):")
	karyawan1.NaikGaji(10) // Gaji karyawan1 bertambah 10%
	karyawan1.TampilkanInfo()

	// ==========================================
	// 3. ANONYMOUS STRUCT (Struct Tanpa Nama Tipe)
	// ==========================================
	// Cocok digunakan untuk data sekali pakai (misal response API sementara)
	userSementara := struct {
		Username string
		IsActive bool
	}{
		Username: "user_guest",
		IsActive: true,
	}

	fmt.Printf("\n3. Anonymous Struct         : User = %s | Active = %t\n",
		userSementara.Username, userSementara.IsActive)
}
