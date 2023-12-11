/*
Copyright © 2023 KLAKALU klakalu438@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/manifoldco/promptui"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitclitool",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		distDir := ".ssh"

		sshKeyName := "id_rsa.pub"

		prompt := promptui.Select{
			Label: "what do you want to do?",
			Items: []string{"check", "ssh-key create"},
		}
		_, out, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("%s\n", out)
		config, _ := cmd.Flags().GetBool("config")
		isShowMsgTrue, _ := cmd.Flags().GetBool("showmsg")
		if config {
			fmt.Println("config is true")
		} else {
			fmt.Println("config is false")

		}

		// check ssh directory exist
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if _, err := os.Stat(filepath.Join(homedir, distDir)); os.IsNotExist(err) {
			// ~/.ssh directory not exist
			if err := os.Mkdir(filepath.Join(homedir, distDir), 0755); err != nil {
				fmt.Println("ssh-key directory create error")
				os.Exit(1)
			}
			fmt.Println("ssh-key directory created")
		} else {
			if isShowMsgTrue {
				fmt.Println("ssh-key directory already exist")
			}
		}

		// check ssh-key exist
		if _, err := os.Stat(filepath.Join(homedir, distDir, sshKeyName)); os.IsNotExist(err) {
			// ~/.ssh/id_rsa file not exist
			out, err := exec.Command("ssh-keygen", "-t", "ed25519", "-N", "", "-f", filepath.Join(homedir, distDir, sshKeyName)).CombinedOutput()
			fmt.Printf("\nssh-keygen result: %s", string(out))
			if err != nil {
				fmt.Println("ssh-keygen error")
				os.Exit(1)
			}
			fmt.Println("ssh-keygen success")
		} else {
			fmt.Println("ssh-key already exist")
		}
		prompt2 := promptui.Select{
			Label: "Copy ssh-key to clipboard?",
			Items: []string{"yes", "no"},
		}
		_, out, err = prompt2.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose %s\n", out)
		if out == "no" {
			os.Exit(1)
		}

		osType := runtime.GOOS
		//get ssh-key value
		var sshKey []byte
		switch osType {
		case "darwin":
			//mac
			var err error
			sshKey, err = exec.Command("cat", filepath.Join(homedir, distDir, sshKeyName)).CombinedOutput()
			if err != nil {
				fmt.Println("ssh-key copy error")
				os.Exit(1)
			}
		case "linux":
			//linux
			var err error
			sshKey, err = exec.Command("cat", filepath.Join(homedir, distDir, sshKeyName)).CombinedOutput()
			if err != nil {
				fmt.Println("ssh-key copy error")
				os.Exit(1)
			}
		case "windows":
			//windows
			var err error
			sshKey, err = exec.Command("powershell", "cat", filepath.Join(homedir, distDir, sshKeyName)).CombinedOutput()
			if err != nil {
				fmt.Println("ssh-key copy error")
				os.Exit(1)
			}
		default:
			//その他
			fmt.Print("sorry, this os is not supported yet")
			os.Exit(1)
		}
		fmt.Println("ssh-key get success")
		//copy ssh-key to clipboard
		switch osType {
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
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("config", "c", false, "congiure")
	rootCmd.PersistentFlags().BoolP("showmsg", "s", false, "show message")
}
