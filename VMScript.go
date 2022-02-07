//REGKEYS
// HKLM\HARDWARE\DEVICEMAP\Scsi\Scsi Port 0\Scsi Bus 0\Target Id 0\Logical Unit Id 0	Identifier	VBOX
// HKLM\HARDWARE\DEVICEMAP\Scsi\Scsi Port 1\Scsi Bus 0\Target Id 0\Logical Unit Id 0	Identifier	VBOX
// HKLM\HARDWARE\DEVICEMAP\Scsi\Scsi Port 2\Scsi Bus 0\Target Id 0\Logical Unit Id 0	Identifier	VBOX
// HKLM\HARDWARE\Description\System	SystemBiosVersion	VBOX
// HKLM\HARDWARE\Description\System	VideoBiosVersion	VIRTUALBOX
// HKLM\HARDWARE\Description\System\BIOS	SystemProductName	VIRTUAL
// HKLM\SYSTEM\ControlSet001\Services\Disk\Enum	DeviceDesc	VBOX
// HKLM\SYSTEM\ControlSet001\Services\Disk\Enum	FriendlyName	VBOX
// HKLM\SYSTEM\ControlSet002\Services\Disk\Enum	DeviceDesc	VBOX
// HKLM\SYSTEM\ControlSet002\Services\Disk\Enum	FriendlyName	VBOX
// HKLM\SYSTEM\ControlSet003\Services\Disk\Enum	DeviceDesc	VBOX
// HKLM\SYSTEM\ControlSet003\Services\Disk\Enum	FriendlyName	VBOX
// HKLM\SYSTEM\CurrentControlSet\Control\SystemInformation	SystemProductName	VIRTUAL
// HKLM\SYSTEM\CurrentControlSet\Control\SystemInformation	SystemProductName	VIRTUALBOX

//CPU ID STUFF
// https://superuser.com/questions/625648/virtualbox-how-to-force-a-specific-cpu-to-the-guest

//STUFF TO HOOK
// HKLM\SYSTEM\CurrentControlSet\Enum\PCI\VEN_80EE*	Subkey has the following structure: VEN_XXXX&DEV_YYYY&SUBSYS_ZZZZ&REV_WW
// HKLM\HARDWARE\ACPI\DSDT\VBOX__
// HKLM\HARDWARE\ACPI\FADT\VBOX__
// HKLM\HARDWARE\ACPI\RSDT\VBOX__
// HKLM\SOFTWARE\Oracle\VirtualBox Guest Additions
// HKLM\SYSTEM\ControlSet001\Services\VBoxGuest
// HKLM\SYSTEM\ControlSet001\Services\VBoxMouse
// HKLM\SYSTEM\ControlSet001\Services\VBoxService
// HKLM\SYSTEM\ControlSet001\Services\VBoxSF
// HKLM\SYSTEM\ControlSet001\Services\VBoxVideo
package testBox

import (
	"fmt"

	"golang.org/x/sys/windows/registry"

	// "io/ioutil"
	"os/exec"
)

func fuckYourProcesses() {
	//kills all identifiable vbox proceesses
	exec.Command("taskkill /IM \"vboxservice.exe\" /F").Run()
	exec.Command("taskkill /IM \"vboxtray.exe\" /F").Run()
}
func fuckYourRegistry() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `HKLM\HARDWARE\Description\System\BIOS`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		fmt.Println(err)
	}
	k.SetStringValue("SystemProductName", "LmaoNope")

	k.Close()
	// ioutil.WriteFile("output.txt", []byte(err.Error()), 0644)
}
func main() {
	fmt.Println("Everything is working :)")
	fuckYourProcesses()
	fuckYourRegistry()

}
