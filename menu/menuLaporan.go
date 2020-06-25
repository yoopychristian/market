//YOOPY CHRISTIAN

package menu

import (
	"challenge/services"
	"fmt"
)

func MenuLaporan() {
	fmt.Println("=========================")
	fmt.Println("     Laporan Penjualan   ")
	fmt.Println("=========================")
	fmt.Println("1. Laporan Harian")
	fmt.Println("2. Laporan Bulanan")
	fmt.Println("3. Semua Laporan")
	fmt.Println("Pilih menu")

	var selectMenu string

	fmt.Scanln(&selectMenu)
	switch selectMenu {
	case "1":
		services.LaporanHarian()
		break
	case "2":
		services.LaporanBulanan()
		break
	case "3":
		services.SemuaLaporan()
		break
	case "4":
		MainMenu()
		ClearScreen()
	default:
		MainMenu()
		ClearScreen()
	}

}
