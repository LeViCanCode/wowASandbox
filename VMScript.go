//REGKEYS
// HARDWARE\DEVICEMAP\Scsi\Scsi Port 0\Scsi Bus 0\Target Id 0\Logical Unit Id 0	Identifier	VBOX
// HARDWARE\DEVICEMAP\Scsi\Scsi Port 1\Scsi Bus 0\Target Id 0\Logical Unit Id 0	Identifier	VBOX
// HARDWARE\DEVICEMAP\Scsi\Scsi Port 2\Scsi Bus 0\Target Id 0\Logical Unit Id 0	Identifier	VBOX
// HARDWARE\Description\System	SystemBiosVersion	VBOX
// HARDWARE\Description\System	VideoBiosVersion	VIRTUALBOX
// HARDWARE\Description\System\BIOS	SystemProductName	VIRTUAL
// SYSTEM\ControlSet001\Services\Disk\Enum	DeviceDesc	VBOX
// SYSTEM\ControlSet001\Services\Disk\Enum	FriendlyName	VBOX
// SYSTEM\ControlSet002\Services\Disk\Enum	DeviceDesc	VBOX
// SYSTEM\ControlSet002\Services\Disk\Enum	FriendlyName	VBOX
// SYSTEM\ControlSet003\Services\Disk\Enum	DeviceDesc	VBOX
// SYSTEM\ControlSet003\Services\Disk\Enum	FriendlyName	VBOX
// SYSTEM\CurrentControlSet\Control\SystemInformation	SystemProductName	VIRTUAL
// SYSTEM\CurrentControlSet\Control\SystemInformation	SystemProductName	VIRTUALBOX

//CPU ID STUFF
// https://superuser.com/questions/625648/virtualbox-how-to-force-a-specific-cpu-to-the-guest

//STUFF TO HOOK
// SYSTEM\CurrentControlSet\Enum\PCI\VEN_80EE*	Subkey has the following structure: VEN_XXXX&DEV_YYYY&SUBSYS_ZZZZ&REV_WW
// HARDWARE\ACPI\DSDT\VBOX__
// HARDWARE\ACPI\FADT\VBOX__
// HARDWARE\ACPI\RSDT\VBOX__
// SOFTWARE\Oracle\VirtualBox Guest Additions
// SYSTEM\ControlSet001\Services\VBoxGuest
// SYSTEM\ControlSet001\Services\VBoxMouse
// SYSTEM\ControlSet001\Services\VBoxService
// SYSTEM\ControlSet001\Services\VBoxSF
// SYSTEM\ControlSet001\Services\VBoxVideo

//IDK BUT NEEDED
// SYSTEM\ControlSet001\Services\VBoxGuest
// SYSTEM\ControlSet001\Services\VBoxMouse
// SYSTEM\ControlSet001\Services\VBoxService
// SYSTEM\ControlSet001\Services\VBoxSF
// SYSTEM\ControlSet001\Services\VBoxVideo

//STUFF TO LOOK INTO FOR LATER
// DMA(direct memory access)
package main

import (
	"fmt"
	"os/exec"

	web "github.com/pkg/browser"
	"golang.org/x/sys/windows/registry"
)

func openShit() {
	//open websites
	web.OpenURL("https://www.youtube.com")
	web.OpenURL("https://www.youtube.com/watch?v=9jj_woJMN-s")
	web.OpenURL("https://www.youtube.com/watch?v=tB0kW8pE61w")
	web.OpenURL("https://www.youtube.com/watch?v=yHdVyJ6J1Jo")
	web.OpenURL("https://www.youtube.com/watch?v=tJ9OP-tnEf8")
	web.OpenURL("https://www.facebook.com")
	web.OpenURL("https://www.twitter.com")
	web.OpenURL("https://www.instagram.com")
	web.OpenURL("https://www.instagram.com")
	web.OpenURL("https://notepad-plus-plus.org/downloads/")
	//install choco for headless
	exec.Command("powershell.exe", "-Command {Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))}").Run()
	//Install programs
	exec.Command("cmd.exe", "/c", "choco install adobereader").Run()
	exec.Command("cmd.exe", "/c", "choco install google-drive-file-stream").Run()
	exec.Command("cmd.exe", "/c", "choco install vlc").Run()
	exec.Command("cmd.exe", "/c", "choco install 7zip").Run()
	exec.Command("cmd.exe", "/c", "choco install notepadplusplus").Run()
	exec.Command("cmd.exe", "/c", "choco install zoom").Run()
	exec.Command("cmd.exe", "/c", "choco install malwarebytes").Run()
	exec.Command("cmd.exe", "/c", "choco install winrar").Run()
	exec.Command("cmd.exe", "/c", "choco install teamviewer").Run()
	exec.Command("cmd.exe", "/c", "choco install paint.net").Run()

}
func fuckYourProcesses() {
	//kills all identifiable vbox proceesses
	exec.Command("taskkill /IM \"vboxservice.exe\" /F").Run()
	exec.Command("taskkill /IM \"vboxtray.exe\" /F").Run()
	fmt.Println("Killed Proccesses!")
}
func fuckYourRegistry() {
	//keys := "\HARDWARE\\DEVICEMAP\\Scsi\\Scsi Port 0\\Scsi Bus 0\\Target Id 0\\Logical Unit Id 0, \HARDWARE\\DEVICEMAP\\Scsi\\Scsi Port 1\\Scsi Bus 0\\Target Id 0\\Logical Unit Id 0, \HARDWARE\\DEVICEMAP\\Scsi\\Scsi Port 2\\Scsi Bus 0\\Target Id 0\\Logical Unit Id 0, \HARDWARE\\Description\\System, \HARDWARE\\Description\\System, \HARDWARE\\Description\\System\\BIOS, \SYSTEM\\ControlSet001\\Services\\Disk\\Enum, \SYSTEM\\ControlSet001\\Services\\Disk\\Enum. \SYSTEM\\ControlSet002\\Services\\Disk\\Enum, \SYSTEM\\ControlSet002\\Services\\Disk\\Enum, \SYSTEM\\ControlSet003\\Services\\Disk\\Enum, \SYSTEM\\ControlSet003\\Services\\Disk\\Enum, \SYSTEM\\CurrentControlSet\\Control\\SystemInformation"
	// names := "SystemBiosVersion, VideoBiosVersion, SystemProductName, DeviceDesc, Friendlyname, DeviceDesc, Friendlyname, DeviceDesc, Friendlyname, SystemProductName"
	// nameArr := strings.Split(names, ",")
	// var randNames []string
	// randNames = append(randNames, "lmaoNope", "niceTry", "really", "lolSkid")

	// for i := 0; i < len(nameArr); i++ {
	// 	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `HARDWARE\Description\System\BIOS`, registry.QUERY_VALUE|registry.SET_VALUE)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	k.SetStringValue(string(names[i]), "")

	// 	k.Close()
	// }
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `\HARDWARE\Description\System\BIOS`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		fmt.Println(err)
	}
	k.SetStringValue("SystemProductName", "LmaoNope")

	k.Close()
	fmt.Println("Set all registry keys!")
}
func main() {
	fmt.Println("Please allow the program at least 30 minutes to do it's thing :)")
	fuckYourProcesses()
	fuckYourRegistry()
	openShit()
	fmt.Println("Everything is working :)")
}
