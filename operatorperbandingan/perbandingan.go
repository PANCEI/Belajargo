// operator perbandingan
// == -> a == b
// != -> a != b
// > -> a > b
// < -> a < b
// >= -> a >= b
// <= -> a <= b

package main

import "fmt"

func main() {
	a := 10
	b := 20

	// 1. Memeriksa kesamaan (Equal to)
	fmt.Println("a == b :", a == b) // Hasil: false

	// 2. Memeriksa ketidaksamaan (Not equal to)
	fmt.Println("a != b :", a != b) // Hasil: true

	// 3. Lebih besar dari (Greater than)
	fmt.Println("a > b  :", a > b) // Hasil: false (10 tidak lebih besar dari 20)

	// 4. Lebih kecil dari (Less than)
	fmt.Println("a < b  :", a < b) // Hasil: true (10 lebih kecil dari 20)

	// 5. Lebih besar dari atau sama dengan (Greater than or equal to)
	fmt.Println("a >= b :", a >= b) // Hasil: false

	// 6. Lebih kecil dari atau sama dengan (Less than or equal to)
	fmt.Println("a <= b :", a <= b) // Hasil: true
}
