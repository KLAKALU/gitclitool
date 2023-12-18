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

		// ask what to do
		prompt := promptui.Select{
			Label: "what do you want to do?",
			Items: []string{"check", "ssh-key create"},
		}
		_, out, err := prompt.Run()
		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}
		switch out {
		case "check":
			// check
			try_login_github(fileDir)
		case "ssh-key create":
			// ssh-key create

			CreateSshKey(OS_TYPE, fileDir)

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
			if out == "no" {
				os.Exit(0)
			}
			JumpToSettingPage(OS_TYPE)
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
