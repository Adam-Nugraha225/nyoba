package main

import "fmt"
func main(){
	var nama string
	var elektronik1, elektronik2, elektronik3, fashion, musik, olahraga string
	var rina, budi, hartono string
	fmt.Scan(&nama)
	
	//produk dan kategori
	
	elektronik1 = "TV, "
	elektronik2 = "headphone, "
	fashion = "baju, "
	musik = "gitar, "
	olahraga = "sepatu, "
	elektronik3 = "kamera, "
	

	//minat dan nama
	rina = elektronik1 + elektronik2 + musik + elektronik3
	budi = fashion + musik
	hartono = olahraga + elektronik3


	if nama == "rina"{
		 fmt.Print(rina)
	} 

	if nama == "budi" {
		fmt.Print(budi)
	} 

	if nama == "hartono" {
		fmt.Print(hartono)
	}

	if nama != "rina" && nama != "budi" && nama != "hartono" {
		fmt.Print("nama tidak ditemukan")
	}

}