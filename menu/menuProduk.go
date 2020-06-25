//YOOPY CHRISTIAN

package menu

import (
	"challenge/services"
	"fmt"
)

func MenuProduk() {
	fmt.Println("======================")
	fmt.Println("Produk Barang")
	fmt.Println("======================")
	fmt.Println("1. Tambah Produk")
	fmt.Println("2. Hapus Produk")
	fmt.Println("3. Detail Produk")
	fmt.Println("Pilih menu")
	var selectMenu string

	fmt.Scanln(&selectMenu)
	switch selectMenu {
	case "1":
		services.TambahProduct()
		break
	case "2":
		services.HapusProduct()
		break
	case "3":
		services.DetailProduct()
		break
	case "4":
		MainMenu()
		ClearScreen()
	default:
		MainMenu()
		ClearScreen()
	}

}
