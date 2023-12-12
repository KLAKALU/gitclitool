/*
Copyright © 2023 KLAKALU klakalu438@gmail.com
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KLAKALU/gitclitool/cmd/create_ssh_key"
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

		// ask what to do
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

		create_ssh_key.CreateSshKey(isShowMsgTrue)
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
