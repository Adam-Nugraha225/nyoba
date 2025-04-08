package main

import "fmt"

func main() {
	var kue string

	// Ceritanya, Abah menunjukkan hasil foto dari Xiao kepada Yanfei
	// Maka input dari pengguna adalah hasil dari foto yang diambil oleh Xiao
	fmt.Println("Masukkan hasil foto dari Xiao (crt):")
	fmt.Scan(&kue)

	// Hasil foto menunjukkan kue masih ada saat Xiao datang
	// Urutan datang: Ningguang, Hutao, Xiao, Childe
	// Jadi:
	// - Ningguang memeriksa kue, tidak mengambil.
	// - Hutao langsung memberi kado tanpa lihat kue, bisa jadi mencuri.
	// - Xiao memotret dan hasil fotonya menunjukkan kue masih ada -> artinya Hutao tidak mencuri
	// - Childe datang terakhir, jadi dialah yang paling mungkin mengambil kue setelah Xiao

	if kue == "ada" || kue == "masih" || kue == "utuh" || kue == "sempurna"{
		fmt.Println("Pelaku yang paling mungkin mengambil kue adalah: Childe")	q
	} else {
		fmt.Println("Informasi tidak cukup atau foto tidak menunjukkan kue, analisis gagal!.")
	}
}
