/*
Copyright © 2023 KLAKALU klakalu438@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

type FileDirectory struct {
	homeDir    string
	distDir    string
	sshKeyName string
}

var rootCmd = &cobra.Command{
	Use:   "gitclitool",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		const OS_TYPE = runtime.GOOS
		var fileDir FileDirectory
		var err error
		fileDir.homeDir, err = os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fileDir.distDir = ".ssh"
		fileDir.sshKeyName = "id_rsa"

		isConnectArrow := checkGithubConnection(fileDir)
		if isConnectArrow {
			fmt.Println("login to github success!")
			return
		}
		fmt.Println("login to github failed!")

		// ssh-key create
		fmt.Println("create ssh key")
		CreateSshKeyFile(OS_TYPE, fileDir)

		// ask open setting page
		prompt := promptui.Select{
			Label: "Open setting page?",
			Items: []string{"yes", "no"},
		}
		_, out, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("You choose %s\n", out)
		if out == "Yes" {
			openSettingPage(OS_TYPE)
		} else {
			os.Exit(0)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
