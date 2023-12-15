package jump_to_settingpage

import (
	"fmt"
	"os"
	"os/exec"
)

func JumpToSettingPage(OS_TYPE string) {
	const SETTING_PAGE_URL = "https://github.com/settings/ssh/new"

	// open setting page
	switch OS_TYPE {
	case "darwin":
		_, err := exec.Command("open", SETTING_PAGE_URL).CombinedOutput()
		if err != nil {
			fmt.Println("failed to open setting page")
			os.Exit(1)
		}
	case "linux":
		_, err := exec.Command("xdg-open", SETTING_PAGE_URL).CombinedOutput()
		if err != nil {
			fmt.Println("failed to open setting page")
			os.Exit(1)
		}
	case "windows":
		_, err := exec.Command("powershell", "start", SETTING_PAGE_URL).CombinedOutput()
		if err != nil {
			fmt.Println("failed to open setting page")
			os.Exit(1)
		}
	default:
		fmt.Print("sorry, this os is not supported yet")
		os.Exit(1)
	}
}
