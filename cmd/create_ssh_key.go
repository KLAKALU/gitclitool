package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

func CreateSshKey(OS_TYPE string, fileDir FileDirectory) {

	// check ssh directory exist
	if _, err := os.Stat(filepath.Join(fileDir.homeDir, fileDir.distDir)); os.IsNotExist(err) {
		// ~/.ssh directory not exist
		if err := os.Mkdir(filepath.Join(fileDir.homeDir, fileDir.distDir), 0755); err != nil {
			fmt.Println("ssh-key directory create error")
			os.Exit(1)
		}
		fmt.Println("ssh-key directory created")
	} else {
		fmt.Println("ssh-key directory already exist")
	}

	// check ssh-key exist
	if _, err := os.Stat(filepath.Join(fileDir.homeDir, fileDir.distDir, fileDir.sshKeyName)); os.IsNotExist(err) {
		// ~/.ssh/id_rsa file not exist
		out, err := exec.Command("ssh-keygen", "-t", "ed25519", "-N", "", "-f", filepath.Join(fileDir.homeDir, fileDir.distDir, fileDir.sshKeyName)).CombinedOutput()
		fmt.Printf("\nssh-keygen result: %s", string(out))
		if err != nil {
			fmt.Println("ssh-keygen error")
			os.Exit(1)
		}
		fmt.Println("ssh-keygen success")
	} else {
		fmt.Println("ssh-key already exist")
	}

	// ask copy ssh-key to clipboard
	prompt := promptui.Select{
		Label: "Copy ssh-key to clipboard?",
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

	//get ssh-key value
	sshKey, err := os.ReadFile(filepath.Join(fileDir.homeDir, fileDir.distDir, fileDir.sshKeyName+".pub"))
	if err != nil {
		panic(err)
	}
	fmt.Println("ssh-key get success")

	//copy ssh-key to clipboard
	switch OS_TYPE {
	case "darwin":
		//mac
		_, err := exec.Command("osascript", "-e", "set the clipboard to \""+string(sshKey)+"\"").CombinedOutput()
		if err != nil {
			fmt.Println("ssh-key copy error")
			os.Exit(1)
		}
	case "linux":
		//linux
		/*
			_, err := exec.Command("xclip", "-selection", "c", "-i").CombinedOutput()
			if err != nil {
				fmt.Println("ssh-key copy error")
				os.Exit(1)
			}
		*/
	case "windows":
		//windows
		_, err := exec.Command("powershell", "\"", string(sshKey), "\"", "|", "clip").CombinedOutput()
		if err != nil {
			fmt.Println("ssh-key copy error")
			os.Exit(1)
		}
	default:
		fmt.Print("sorry, this os is not supported yet")
		os.Exit(1)
		//その他
	}
	fmt.Println("ssh-key copy success")
}
