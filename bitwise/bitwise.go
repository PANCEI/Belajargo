package main

import "fmt"

func main() {
	// Variabel bertipe uint8 (8 bit)
	// Angka 12 dalam format biner 8 bit : 0000 1100
	// Angka 10 dalam format biner 8 bit : 0000 1010
	var a uint8 = 12
	var b uint8 = 10

	fmt.Printf("Nilai a = %d (biner: %08b)\n", a, a)
	fmt.Printf("Nilai b = %d (biner: %08b)\n\n", b, b)

	// 1. Bitwise AND (&)
	// Menghasilkan bit 1 HANYA JIKA kedua bit berposisi sama bernilai 1.
	// 0000 1100 (12)
	// 0000 1010 (10)
	// --------- &
	// 0000 1000 (Hasil: 8)
	fmt.Printf("a & b  = %d (Biner: %08b) -> Bitwise AND\n", a&b, a&b)

	// 2. Bitwise OR (|)
	// Menghasilkan bit 1 JIKA SALAH SATU atau KEDUANYA bernilai 1.
	// 0000 1100 (12)
	// 0000 1010 (10)
	// --------- |
	// 0000 1110 (Hasil: 14)
	fmt.Printf("a | b  = %d (Biner: %08b) -> Bitwise OR\n", a|b, a|b)

	// 3. Bitwise XOR (^)
	// Menghasilkan bit 1 JIKA NILAI BIT BERBEDA (satu 0, satu 1).
	// 0000 1100 (12)
	// 0000 1010 (10)
	// --------- ^
	// 0000 0110 (Hasil: 6)
	fmt.Printf("a ^ b  = %d (Biner: %08b) -> Bitwise XOR\n", a^b, a^b)

	// 4. Bitwise AND NOT (&^) - Operator khas Go (Bit Clear)
	// Mematikan (mengubah jadi 0) bit 'a' pada posisi di mana bit 'b' bernilai 1.
	// 0000 1100 (12)
	// 0000 1010 (10) - Bit 1 ada di posisi ke-2 dan ke-4 dari kanan
	// --------- &^
	// 0000 0100 (Hasil: 4) - Bit posisi ke-2 pada 'a' yang tadinya 1 dipaksa jadi 0
	fmt.Printf("a &^ b = %d (Biner: %08b) -> Bitwise AND NOT\n", a&^b, a&^b)

	// 5. Bitwise NOT / Complement (^) - Unary Operator
	// Membalikkan semua bit (0 jadi 1, 1 jadi 0).
	// ^0000 1100 (12)
	// -----------
	//  1111 0011 (Hasil: 243)
	fmt.Printf("^a     = %d (Biner: %08b) -> Bitwise NOT\n", ^a, ^a)

	// 6. Left Shift (<<)
	// Menggeser bit ke kiri sebanyak 2 posisi.
	// Memasukkan angka 0 di sebelah kanan. Sama seperti mengalikan dengan 2^2 (12 * 4 = 48).
	// 0000 1100 (12) << 2
	// ---------
	// 0011 0000 (Hasil: 48)
	fmt.Printf("a << 2 = %d (Biner: %08b) -> Shift Left 2 Bit\n", a<<2, a<<2)

	// 7. Right Shift (>>)
	// Menggeser bit ke kanan sebanyak 2 posisi.
	// Membuang bit paling kanan. Sama seperti membagi dengan 2^2 (12 / 4 = 3).
	// 0000 1100 (12) >> 2
	// ---------
	// 0000 0011 (Hasil: 3)
	fmt.Printf("a >> 2 = %d (Biner: %08b) -> Shift Right 2 Bit\n", a>>2, a>>2)
}
