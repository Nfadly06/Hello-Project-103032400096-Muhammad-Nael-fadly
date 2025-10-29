package main

import "fmt"

func menu() {
	fmt.Println("\n=== MENU SEA GAMES MANAGER ===")
	fmt.Println("1. Tambah Negara")
	fmt.Println("2. Edit Nama Negara")
	fmt.Println("3. Hapus Negara")
	fmt.Println("4. Edit Data Medali")
	fmt.Println("5. Hapus Medali (reset ke 0)")
	fmt.Println("6. Tampilkan Peringkat")
	fmt.Println("7. Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var dataNegara DataNegara
	var dataMedali DataMedali
	var nData, pilihan int

	for pilihan != 7 {
		menu()
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			tambahNegara(&dataNegara, &dataMedali, &nData)
		case 2:
			editNegara(&dataNegara, nData)
		case 3:
			hapusNegara(&dataNegara, &dataMedali, &nData)
		case 4:
			editMedali(dataNegara, &dataMedali, nData)
		case 5:
			hapusMedali(dataNegara, &dataMedali, nData)
		case 6:
			tampilkan(dataNegara, dataMedali, nData)
		case 7:
			fmt.Println("Terima kasih!")
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
