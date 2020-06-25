//YOOPY CHRISTIAN

package main

import (
	"challenge/landing"
	"challenge/menu"
	"fmt"
	"os"
)

func main() {
	landingBaru()
}

func landingBaru() {
	menu.LandingLogin()
	for {
		var selectedLanding string
		fmt.Scanln(&selectedLanding)
		switch selectedLanding {
		case "1":
			landing.LoginPengguna()
			break
		case "2":
			menu.RegisterMenus()
			for {
				var selectedRegister string
				fmt.Scanln(&selectedRegister)
				switch selectedRegister {
				case "1":
					landing.TambahPengguna()
					break
				case "2":
					landing.HapusPengguna()
					break
				case "3":
					landing.DetailPengguna()
					break
				}
				break
			}
		case "3":
			os.Exit(0)
		default:
			menu.ClearScreenLoginRegister()
		}
	}
}
