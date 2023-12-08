// package main

// import "fmt"

// func main() {
// 	hargaBarang := 100.0
// 	jumlahPembelian := 3
// 	totalPembayaran := hargaBarang * float64(jumlahPembelian)
// 	var diskon float64
// 	var bayarDiskon float64

// 	if totalPembayaran <= 250 {
// 		diskon = 0.3
// 	} else if totalPembayaran <= 150 {
// 		diskon = 0.2
// 	} else {
// 		diskon = 0.1
// 	}
// 	fmt.Printf("%.2f\n", totalPembayaran)
// 	bayarDiskon = totalPembayaran * diskon
// 	fmt.Printf("%.2f", bayarDiskon)
// 	totalBayar := totalPembayaran - bayarDiskon
// 	fmt.Printf("\nTotal bayar %.2f ", totalBayar)
// }