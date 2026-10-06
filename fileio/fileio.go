/*
os
1. membuat file (os.Create)
2. membuka file (os.Open)
3. menulis file (os.Write) / os.WriteString
4. membaca file (os.Read)
5. menutup file (os.Close)

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file_kita, error := os.Open("file_kita.txt")
	if error != nil {
		fmt.Println("Terjadi error saat membuka file:", error)
	}
	defer file_kita.Close() // Menutup file setelah selesai digunakan

	scanner := bufio.NewScanner(file_kita)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Terjadi error saat membaca file:", err)
	}
	// jika ingin menjalankan program ini , pastikank di jalankan di folder yang sama dengan file_kita.txt

	// mambaca file dengan cara beda
	data_kita, err := os.ReadFile("file_kita.txt")
	if err != nil {
		fmt.Println("tidak bisa membaca file error : ", err)
		return
	}
	fmt.Println(string(data_kita))

	// memmbuat file baru dan menulis file baru
	file_baru, errorFileBaru := os.Create("data_kita.txt")
	if errorFileBaru != nil {
		fmt.Println("Gagal Membuat file baru", errorFileBaru)
	}
	defer file_baru.Close()
	_, errorTulis := file_baru.WriteString("Ini Adalah Tulisan Baruo \n")
	if errorTulis != nil {
		fmt.Println("gagal melakukan penuslisan dalam file", errorTulis)
		return
	}
	fmt.Println("berhasil melakukan penulisan pada file")

	// menambahkan tulisan baru pada file tanpa  menimpah
	kata_baru, errs := os.OpenFile("data_kita.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if errs != nil {
		fmt.Println("error saat melakukan open file", errs)
		return
	}
	defer kata_baru.Close()
	//  tambah kata
	_, errors := kata_baru.WriteString("sadasdasdasdsad \n")
	if errors != nil {
		fmt.Println("gagal menambahkan kata baru ", errors)
		return
	}
	fmt.Println("kata baru berhasil di tambahkan")
}
