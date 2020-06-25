//YOOPY CHRISTIAN

package landing

import (
	"challenge/config"
	"challenge/menu"
	"challenge/services"
	"challenge/utils"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

type Login struct {
	username string
	password string
}

func LoginPengguna() {
	var username, password string
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("======================================================")
	fmt.Println("                    ENIGMA MART LOGIN                 ")
	fmt.Println("======================================================")
	fmt.Print("Username : ")
	fmt.Scan(&username)
	fmt.Print("Password : ")
	fmt.Scan(&password)
	rows, err := db.Query("Select * from login where username=? and password=?", username, password)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var users []*Login
	for rows.Next() {
		var user = Login{}
		var err = rows.Scan(&user.username, &user.password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &user)
	}
	isUser := false
	for _, u := range users {
		if u.username != "" && u.password != "" {
			isUser = true
		} else {
			isUser = false
		}
	}
	ValidLogin(isUser, db)

}
func ValidLogin(isUser bool, db *sql.DB) {
	dbEngine, dbSource := config.Env_conn()
	db, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	if isUser {
		fmt.Println("Login Berhasil")
		time.Sleep(1 * time.Second)
		GoToMainMenu()
	} else {
		fmt.Println("Username/Password yang anda masukan salah, silahkan coba lagi")
		time.Sleep(1 * time.Second)
		menu.LandingLogin()
	}
}

func GoToMainMenu() {
	dbEngine, dbSource := config.Env_conn()
	_, err := utils.ConnDB(dbEngine, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	menu.MainMenu()
	for {
		var selectMenu string
		fmt.Scanln(&selectMenu)
		switch selectMenu {
		case "1":
			menu.MenuProduk()
			for {
				var selectedProductMenu string
				fmt.Scanln(&selectedProductMenu)
				switch selectedProductMenu {
				case "1":
					services.TambahProduct()
					break
				case "2":
					services.HapusProduct()
					break
				case "3":
					services.DetailProduct()
					break
				}
				break
			}
		case "2":
			menu.MenuTransaksi()
			for {
				var selectedBillMenu string
				fmt.Scanln(&selectedBillMenu)
				switch selectedBillMenu {
				case "1":
					services.TambahTransaksi()
					break
				case "2":
					services.HapusTransaksi()
					break
				case "3":
					services.DetailTransaksi()
				}
				break
			}
		case "3":
			menu.MenuLaporan()
			for {
				var selectedReportMenu string
				fmt.Scanln(&selectedReportMenu)
				switch selectedReportMenu {
				case "1":
					services.LaporanHarian()
					break
				case "2":
					services.LaporanBulanan()
					break
				case "3":
					services.SemuaLaporan()
				}
				break
			}

		case "4":
			fmt.Print("LOGGING OUT")
			time.Sleep(1 * time.Second)
			os.Exit(0)
		default:
			menu.ClearScreenLoginRegister()
		}
	}
}
