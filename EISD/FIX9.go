package main

import "fmt"

func main() {
	var n1, n2, n3, n4, n5, n6, n7, n8, n9, n10 string
	var nama1, nama2, nama3, nama4, nama5, nama6, nama7, nama8, nama9, nama10 string
	var jumlah1, jumlah2, jumlah3, jumlah4, jumlah5, jumlah6, jumlah7, jumlah8, jumlah9, jumlah10 int

	fmt.Println("Masukkan 10 nama yang disebut anak-anak (satu per satu):")
	fmt.Print("Nama ke-1: ")
	fmt.Scan(&n1)
	fmt.Print("Nama ke-2: ")
	fmt.Scan(&n2)
	fmt.Print("Nama ke-3: ")
	fmt.Scan(&n3)
	fmt.Print("Nama ke-4: ")
	fmt.Scan(&n4)
	fmt.Print("Nama ke-5: ")
	fmt.Scan(&n5)
	fmt.Print("Nama ke-6: ")
	fmt.Scan(&n6)
	fmt.Print("Nama ke-7: ")
	fmt.Scan(&n7)
	fmt.Print("Nama ke-8: ")
	fmt.Scan(&n8)
	fmt.Print("Nama ke-9: ")
	fmt.Scan(&n9)
	fmt.Print("Nama ke-10: ")
	fmt.Scan(&n10)

	// Step 1: Masukkan nama pertama
	nama1 = n1
	jumlah1 = 1

	// Step 2: Cek nama kedua
	if n2 == nama1 {
		jumlah1++
	} else {
		nama2 = n2
		jumlah2 = 1
	}

	// Step 3: Cek nama ketiga
	if n3 == nama1 {
		jumlah1++
	} else if n3 == nama2 {
		jumlah2++
	} else {
		nama3 = n3
		jumlah3 = 1
	}

	// Step 4: Cek nama keempat
	if n4 == nama1 {
		jumlah1++
	} else if n4 == nama2 {
		jumlah2++
	} else if n4 == nama3 {
		jumlah3++
	} else {
		nama4 = n4
		jumlah4 = 1
	}

	// Step 5: Cek nama kelima
	if n5 == nama1 {
		jumlah1++
	} else if n5 == nama2 {
		jumlah2++
	} else if n5 == nama3 {
		jumlah3++
	} else if n5 == nama4 {
		jumlah4++
	} else {
		nama5 = n5
		jumlah5 = 1
	}

	// Step 6: Cek nama keenam
	if n6 == nama1 {
		jumlah1++
	} else if n6 == nama2 {
		jumlah2++
	} else if n6 == nama3 {
		jumlah3++
	} else if n6 == nama4 {
		jumlah4++
	} else if n6 == nama5 {
		jumlah5++
	} else {
		nama6 = n6
		jumlah6 = 1
	}

	// Step 7: Cek nama ketujuh
	if n7 == nama1 {
		jumlah1++
	} else if n7 == nama2 {
		jumlah2++
	} else if n7 == nama3 {
		jumlah3++
	} else if n7 == nama4 {
		jumlah4++
	} else if n7 == nama5 {
		jumlah5++
	} else if n7 == nama6 {
		jumlah6++
	} else {
		nama7 = n7
		jumlah7 = 1
	}

	// Step 8: Cek nama kedelapan
	if n8 == nama1 {
		jumlah1++
	} else if n8 == nama2 {
		jumlah2++
	} else if n8 == nama3 {
		jumlah3++
	} else if n8 == nama4 {
		jumlah4++
	} else if n8 == nama5 {
		jumlah5++
	} else if n8 == nama6 {
		jumlah6++
	} else if n8 == nama7 {
		jumlah7++
	} else {
		nama8 = n8
		jumlah8 = 1
	}

	// Step 9: Cek nama kesembilan
	if n9 == nama1 {
		jumlah1++
	} else if n9 == nama2 {
		jumlah2++
	} else if n9 == nama3 {
		jumlah3++
	} else if n9 == nama4 {
		jumlah4++
	} else if n9 == nama5 {
		jumlah5++
	} else if n9 == nama6 {
		jumlah6++
	} else if n9 == nama7 {
		jumlah7++
	} else if n9 == nama8 {
		jumlah8++
	} else {
		nama9 = n9
		jumlah9 = 1
	}

	// Step 10: Cek nama kesepuluh
	if n10 == nama1 {
		jumlah1++
	} else if n10 == nama2 {
		jumlah2++
	} else if n10 == nama3 {
		jumlah3++
	} else if n10 == nama4 {
		jumlah4++
	} else if n10 == nama5 {
		jumlah5++
	} else if n10 == nama6 {
		jumlah6++
	} else if n10 == nama7 {
		jumlah7++
	} else if n10 == nama8 {
		jumlah8++
	} else if n10 == nama9 {
		jumlah9++
	} else {
		nama10 = n10
		jumlah10 = 1
	}

	// Cari jumlah maksimum
	max := jumlah1
	if jumlah2 > max {
		max = jumlah2
	}
	if jumlah3 > max {
		max = jumlah3
	}
	if jumlah4 > max {
		max = jumlah4
	}
	if jumlah5 > max {
		max = jumlah5
	}
	if jumlah6 > max {
		max = jumlah6
	}
	if jumlah7 > max {
		max = jumlah7
	}
	if jumlah8 > max {
		max = jumlah8
	}
	if jumlah9 > max {
		max = jumlah9
	}
	if jumlah10 > max {
		max = jumlah10
	}

	// Tampilkan hasil
	if max == 1 {
		fmt.Println("Semuanya anak baik")
	} else {
		fmt.Print("Anak Nakal: ")
		tambahSpasi := false
		if jumlah1 == max {
			fmt.Print(nama1)
			tambahSpasi = true
		}
		if jumlah2 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama2)
			tambahSpasi = true
		}
		if jumlah3 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama3)
			tambahSpasi = true
		}
		if jumlah4 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama4)
			tambahSpasi = true
		}
		if jumlah5 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama5)
			tambahSpasi = true
		}
		if jumlah6 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama6)
			tambahSpasi = true
		}
		if jumlah7 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama7)
			tambahSpasi = true
		}
		if jumlah8 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama8)
			tambahSpasi = true
		}
		if jumlah9 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama9)
			tambahSpasi = true
		}
		if jumlah10 == max {
			if tambahSpasi {
				fmt.Print(" dan ")
			}
			fmt.Print(nama10)
		}
		fmt.Println()
	}
}
