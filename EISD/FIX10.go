package main

import (
	"fmt"
)

func main() {
	var totalPayment, totalPurchase, change int

	fmt.Print("Masukkan total pembayaran: ")
	fmt.Scan(&totalPayment)
	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalPurchase)

	if totalPayment < totalPurchase {
		fmt.Println("Error: Pembayaran kurang dari total belanja!")
		return
	}

	change = totalPayment - totalPurchase
	fmt.Println("Output:")

	if change >= 100000 {
		fmt.Printf("100000: %d\n", change/100000)
		change %= 100000
	}
	if change >= 50000 {
		fmt.Printf("50000: %d\n", change/50000)
		change %= 50000
	}
	if change >= 20000 {
		fmt.Printf("20000: %d\n", change/20000)
		change %= 20000
	}
	if change >= 10000 {
		fmt.Printf("10000: %d\n", change/10000)
		change %= 10000
	}
	if change >= 5000 {
		fmt.Printf("5000: %d\n", change/5000)
		change %= 5000
	}
	if change >= 2000 {
		fmt.Printf("2000: %d\n", change/2000)
		change %= 2000
	}
	if change >= 1000 {
		fmt.Printf("1000: %d\n", change/1000)
		change %= 1000
	}
	if change >= 500 {
		fmt.Printf("500: %d\n", change/500)
		change %= 500
	}
	if change >= 200 {
		fmt.Printf("200: %d\n", change/200)
		change %= 200
	}
	if change >= 100 {
		fmt.Printf("100: %d\n", change/100)
		change %= 100
	}
}
