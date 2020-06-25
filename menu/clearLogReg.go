//YOOPY CHRISTIAN

package menu

import (
	"os"
	"os/exec"
	"runtime"
)

func ClearScreenLoginRegister() {
	osRunning := runtime.GOOS
	if osRunning == "linux" || osRunning == "darwin" {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else if osRunning == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	LandingLogin()
}
