/*
Copyright © 2023 KLAKALU klakalu438@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gitclitool",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		config, _ := cmd.Flags().GetBool("config")
		showMsg, _ := cmd.Flags().GetBool("showmsg")
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
			if showMsg {
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
			if showMsg {
				fmt.Println("ssh-key already exist")
			}
			os.Exit(1)
		}
		//ssh-keyをクリップボードにコピー
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitclitool.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().BoolP("config", "c", false, "congiure")
	rootCmd.PersistentFlags().BoolP("showmsg", "s", false, "show message")
}
