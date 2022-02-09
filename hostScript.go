//check if file is ran in a terminal https://github.com/mattn/go-isatty
import (
	"fmt"
	"runtime"
	"os"
	"flag"
	"strings"
	"github.com/mattn/go-isatty"
)
var exists bool

func windows() {
	userDir, _ := os.UserHomeDir()
	if _, err := os.Stat(fmt.Sprintf("%s\\VirtualBox VMs", userDir)) !os.IsNotExist(err) {
        exists = true
	} else {
		fmt.Println("Please specify a path to wher your VM's are saved")
		os.Exit()
	}
}
func linux() {
	userDir, _ := os.UserHomeDir()
	if _, err := os.Stat(fmt.Sprintf("%s/VirtualBox VMs", userDir) !os.IsNotExist(err) {
        exists = true
	} else {
		fmt.Println("Please specify a path to wher your VM's are saved")
		os.Exit()
	}
}
func main() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Cygwin/MSYS2 Terminal")
	} else {
		os.Exit()
	}
	newPath := flag.String("path")
	flag.Parse()
	if !strings.Matches(*newPath, /[/]+/g) {
		fmt.Println("Please include the full file path.")
		os.Exit()
	}
	if runtime.GOOS == "windows" {
		windows()
	} else {
		linux()
	}
}
