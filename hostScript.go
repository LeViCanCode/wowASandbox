package main

//check if file is ran in a terminal https://github.com/mattn/go-isatty
import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/klauspost/cpuid/v2"
	"github.com/mattn/go-isatty"
)

func windows() {
	userDir, _ := os.UserHomeDir()
	name := cpuid.CPU.BrandName
	physCores := cpuid.CPU.PhysicalCores
	threadsPC := cpuid.CPU.ThreadsPerCore
	logicCores := cpuid.CPU.LogicalCores
	family := cpuid.CPU.Family
	model := cpuid.CPU.Model
	vendorID := cpuid.CPU.VendorID
	vm := cpuid.CPU.VM()
	fmt.Printf(name, "\n", physCores, "\n", threadsPC, "\n", logicCores, "\n", family, "\n", model, "\n", vendorID)
	if vm {
		fmt.Println("cpu id not spoofed")
	}

	if _, err := os.Stat(fmt.Sprintf("%s\\VirtualBox VMs", userDir)); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("exists")
		list := exec.Command("vboxmanage", "list").Run()
		fmt.Println(list)
		// os.Exit(0)

	} else {
		fmt.Println("Please specify a path to where your VMs are saved")
		os.Exit(0)
	}
}
func linux() {
	userDir, _ := os.UserHomeDir()
	if _, err := os.Stat(fmt.Sprintf("%s\\VirtualBox VMs", userDir)); !errors.Is(err, os.ErrNotExist) {
		fmt.Println("exists")
		os.Exit(0)
	} else {
		fmt.Println("Please specify a path to where your VMs are saved")
		os.Exit(0)
	}
}
func main() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Is Cygwin/MSYS2 Terminal")
	} else {
		os.Exit(0)
		//later stuff
		// newPath := flag.String("path")
		// flag.Parse()
	}
	if runtime.GOOS == "windows" {
		windows()
	} else {
		linux()
	}
}
