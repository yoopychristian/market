//YOOPY CHRISTIAN

package menu

import (
	"challenge/services"
	"fmt"
)

func MenuTransaksi() {
	fmt.Println("======================")
	fmt.Println("  Transaksi Penjualan ")
	fmt.Println("======================")
	fmt.Println("1. Tambah Transaksi Penjualan")
	fmt.Println("2. Hapus Transaksi Penjualan")
	fmt.Println("3. Detail Transaksi Penjualan")
	fmt.Println("Pilih menu")
	var selectMenu string

	fmt.Scanln(&selectMenu)
	switch selectMenu {
	case "1":
		services.TambahTransaksi()
		break
	case "2":
		services.HapusTransaksi()
		break
	case "3":
		services.DetailTransaksi()
		break
	case "4":
		MainMenu()
		ClearScreen()
	default:
		MainMenu()
		ClearScreen()
	}

}
