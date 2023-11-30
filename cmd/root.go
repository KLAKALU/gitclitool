/*
Copyright © 2023 KLAKALU klakalu438@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
		config, _ := cmd.Flags().GetBool("config")
		isShowMsgTrue, _ := cmd.Flags().GetBool("showmsg")
		if config {
			fmt.Println("config is true")
		} else {
			fmt.Println("config is false")

		}
		homedir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		directlyName := "/.ssh"
		if _, err := os.Stat(homedir + directlyName); os.IsNotExist(err) {
			// ~/.ssh directory not exist
			if err := os.Mkdir(homedir+directlyName, 0755); err != nil {
				fmt.Println("ssh-key directory create error")
				os.Exit(1)
			}
			fmt.Println("ssh-key directory created")
		} else {
			if isShowMsgTrue {
				fmt.Println("ssh-key directory already exist")
			}
		}
		sshKeyName := "id_rsa"
		if _, err := os.Stat(homedir + directlyName + "/" + sshKeyName); os.IsNotExist(err) {
			// ~/.ssh/id_rsa file not exist
			out, err := exec.Command("ssh-keygen", "-t", "ed25519", "-N", "", "-f", homedir+directlyName+"/"+sshKeyName).CombinedOutput()
			fmt.Printf("\nssh-keygen result: %s", string(out))
			if err != nil {
				fmt.Println("ssh-keygen error")
				os.Exit(1)
			}
			fmt.Println("ssh-keygen success")
		} else {
			fmt.Println("ssh-key already exist")
		}
		prompt := promptui.Select{
			Label: "Select Type",
			Items: []string{"yes", "no"},
		}
		_, out, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}
		fmt.Printf("You choose %s\n", out)
		if out == "no" {
			os.Exit(1)
		}

		osType := runtime.GOOS
		//ssh-keyを取得
		var sshKey []byte
		switch osType {
		case "darwin":
			//mac
			var err error
			sshKey, err = exec.Command("cat", homedir+directlyName+"/"+sshKeyName+".pub").CombinedOutput()
			if err != nil {
				fmt.Println("ssh-key copy error")
				os.Exit(1)
			}
		case "linux":
			//linux
		case "windows":
			//windows
		default:
			//その他
		}
		//ssh-keyをクリップボードにコピー
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
		case "windows":
			//windows
		default:
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
