package jump_to_settingpage

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
)

func JumpToSettingPage(OS_TYPE string) {
	const SETTING_PAGE_URL = "https://github.com/settings/ssh/new"

	prompt := promptui.Select{
		Label: "Open setting page?",
		Items: []string{"yes", "no"},
	}
	_, out, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose %s\n", out)
	if out == "no" {
		return
	}

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
