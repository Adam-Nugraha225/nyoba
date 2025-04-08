package main

import "fmt"

func main() {
	var nama, makanan, minuman string
	var jumlahPesananMakanan, jumlahPesananMinuman int

	// Nama menu
	makanan1 := "ayamgorengkrispi"
	makanan2 := "ayampukpuk"
	makanan3 := "ayambakar"
	minuman1 := "esteh"
	minuman2 := "esjeruk"

	// Harga menu
	hargaMakanan1 := 15000
	hargaMakanan2 := 13000
	hargaMakanan3 := 20000
	hargaMinuman1 := 5000
	hargaMinuman2 := 7000

	// Pajak
	pajakMakanan := 0.05
	pajakMinuman := 0.03
	pajakTransaksi := 0.15

	fmt.Println("|          Menu           |       Harga      |")
	fmt.Println("|   Ayam goreng krispi    |       15000      |")
	fmt.Println("|   Ayam puk puk          |       13000      |")
	fmt.Println("|   Ayam bakar            |       20000      |")
	fmt.Println("|   Es teh                |       5000       |")
	fmt.Println("|   Es jeruk              |       7000       |")
	
	fmt.Print("Masukkan nama anda: ")
	fmt.Scan(&nama)

	fmt.Println("gunakan huruf kecil semua dan jangan gunakan spasi")
	fmt.Print("Masukkan nama pesanan makanan: ")
	fmt.Scan(&makanan)
	fmt.Print("Masukkan jumlah pesanan makanan: ")
	fmt.Scan(&jumlahPesananMakanan)

	fmt.Print("Masukkan nama pesanan minuman: ")
	fmt.Scan(&minuman)
	fmt.Print("Masukkan jumlah pesanan minuman: ")
	fmt.Scan(&jumlahPesananMinuman)

	var hargaMakanan float64
	if makanan == makanan1 {
		hargaMakanan = float64(hargaMakanan1)
	} else if makanan == makanan2 {
		hargaMakanan = float64(hargaMakanan2)
	} else if makanan == makanan3 {
		hargaMakanan = float64(hargaMakanan3)
	} else {
		fmt.Println("Makanan tidak tersedia.")
		return
	}

	var hargaMinuman float64
	if minuman == minuman1 {
		hargaMinuman = float64(hargaMinuman1)
	} else if minuman == minuman2 {
		hargaMinuman = float64(hargaMinuman2)
	} else {
		fmt.Println("Minuman tidak tersedia.")
		return
	}

	// Hitung total makanan dan minuman (dengan pajak makanan & minuman)
	totalMakanan := hargaMakanan * (1 + pajakMakanan) * float64(jumlahPesananMakanan)
	totalMinuman := hargaMinuman * (1 + pajakMinuman) * float64(jumlahPesananMinuman)

	// Hitung total semua (dengan pajak transaksi)
	subtotal := totalMakanan + totalMinuman
	totalKeseluruhan := subtotal * (1 + pajakTransaksi)

	// Output
	fmt.Println("===================================")
	fmt.Println("Nama Pelanggan: ",nama)
	fmt.Println("Total makanan + pajak: Rp.",totalMakanan)
	fmt.Println("Total minuman + pajak: Rp.",totalMinuman)
	fmt.Printf("Total keseluruhan (dengan pajak transaksi): Rp.%.0f\n", totalKeseluruhan)
}
