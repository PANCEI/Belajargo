package main

import "fmt"

func main() {
	// ==========================================
	// KELOMPOK 1: OPERATOR PENUGASAN ARITMATIKA
	// ==========================================

	// 1. Operator Penugasan Biasa (=)
	x := 10
	fmt.Println("Nilai awal x         :", x) // Output: 10

	// 2. Addition Assignment (+=) -> x = x + 5
	x += 5
	fmt.Println("x += 5  (x = 10 + 5)  :", x) // Output: 15

	// 3. Subtraction Assignment (-=) -> x = x - 5
	x -= 5
	fmt.Println("x -= 5  (x = 15 - 5)  :", x) // Output: 10

	// 4. Multiplication Assignment (*=) -> x = x * 5
	x *= 5
	fmt.Println("x *= 5  (x = 10 * 5)  :", x) // Output: 50

	// 5. Division Assignment (/=) -> x = x / 5
	x /= 5
	fmt.Println("x /= 5  (x = 50 / 5)  :", x) // Output: 10

	// 6. Modulus Assignment (%=) -> x = x % 5
	// Mencari sisa bagi dari 10 dibagi 5 (sisanya 0)
	x %= 5
	fmt.Println("x %= 5  (x = 10 % 5)  :", x) // Output: 0

	// ==========================================
	// KELOMPOK 2: OPERATOR PENUGASAN BITWISE
	// ==========================================
	// Kita ubah nilai x jadi 12 (Biner: 0000 1100) untuk operasi bitwise
	x = 12
	fmt.Printf("\n--- Operasi Bitwise (Nilai x awal = %d / %08b) ---\n", x, x)

	// 7. Bitwise AND Assignment (&=) -> x = x & 2
	// 0000 1100 (12) & 0000 0010 (2) = 0000 0000 (0)
	x &= 2
	fmt.Println("x &= 2  (x = x & 2)   :", x) // Output: 0

	// Set nilai x = 12 lagi
	x = 12

	// 8. Bitwise OR Assignment (|=) -> x = x | 2
	// 0000 1100 (12) | 0000 0010 (2) = 0000 1110 (14)
	x |= 2
	fmt.Println("x |= 2  (x = x | 2)   :", x) // Output: 14

	// Set nilai x = 12 lagi
	x = 12

	// 9. Bitwise XOR Assignment (^=) -> x = x ^ 2
	// 0000 1100 (12) ^ 0000 0010 (2) = 0000 1110 (14)
	x ^= 2
	fmt.Println("x ^= 2  (x = x ^ 2)   :", x) // Output: 14

	// Set nilai x = 12 lagi
	x = 12

	// 10. Left Shift Assignment (<<=) -> x = x << 2
	// Menggeser bit 12 ke kiri 2 langkah (sama seperti 12 * 2^2 = 48)
	x <<= 2
	fmt.Println("x <<= 2 (x = x << 2)  :", x) // Output: 48

	// 11. Right Shift Assignment (>>=) -> x = x >> 2
	// Menggeser bit 48 ke kanan 2 langkah (sama seperti 48 / 2^2 = 12)
	x >>= 2
	fmt.Println("x >>= 2 (x = x >> 2)  :", x) // Output: 12
}
