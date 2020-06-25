//YOOPY CHRISTIAN

package menu

import (
	"fmt"
	"os"
)

func MainMenu() {
	fmt.Println("====================")
	fmt.Println("      MAIN MENU     ")
	fmt.Println("====================")
	fmt.Println("1. Produk Barang")
	fmt.Println("2. Transaksi Penjualan")
	fmt.Println("3. Laporan Penjualan")
	fmt.Println("4. Exit")
	fmt.Println("Pilih menu")
	for {
		var menu string

		fmt.Scanln(&menu)
		switch menu {
		case "1":
			MenuProduk()
			break
		case "2":
			MenuTransaksi()
			break
		case "3":
			MenuLaporan()
			break
		case "4":
			os.Exit(0)
		default:
			ClearScreen()
		}

	}
}
