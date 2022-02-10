package main

//check if file is ran in a terminal https://github.com/mattn/go-isatty
import (
	"errors"
	"fmt"
	"os"
	"runtime"

	"github.com/mattn/go-isatty"
)

func windows() {
	userDir, _ := os.UserHomeDir()
	if _, err := os.Stat(fmt.Sprintf("%s\\VirtualBox VMs", userDir)); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("exists")
		os.Exit(0)

	} else {
		fmt.Println("Please specify a path to wher your VM's are saved")
		os.Exit(0)
	}
}
func linux() {
	userDir, _ := os.UserHomeDir()
	if _, err := os.Stat(fmt.Sprintf("%s/VirtualBox VMs", userDir)); !os.IsNotExist(err) {
	} else {
		fmt.Println("Please specify a path to wher your VM's are saved")
		os.Exit(0)
	}
}
func main() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Cygwin/MSYS2 Terminal")
	} //else {
	// os.Exit(0)
	//ereor
	//}
	// newPath := flag.String("path")
	// flag.Parse()
	// if !strings.Matches(*newPath, /[/]+/g) {
	// 	fmt.Println("Please include the full file path.")
	// 	os.Exit()
	// }
	if runtime.GOOS == "windows" {
		windows()
	} else {
		linux()
	}
}
