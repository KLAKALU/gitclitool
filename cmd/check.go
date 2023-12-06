/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup

		isGithubConnected := false

		wg.Add(1)

		go loadingAnimation(&wg)

		go tryConnectToGithub(&isGithubConnected, &wg)

		wg.Wait()

		if isGithubConnected {
			fmt.Println("\nconnected to github successfully!")
		} else {
			fmt.Println("\nfailed to connect to github!")
		}
	},
}

func loadingAnimation(wg *sync.WaitGroup) {
	marks := []string{"   ", ".  ", ".. ", "..."}
	for i := 0; i < 500; i++ {
		fmt.Printf("\rconecting to github %s", marks[i%4])
		time.Sleep(250 * time.Millisecond)
	}
	wg.Done()
}

func tryConnectToGithub(isGithubConnected *bool, wg *sync.WaitGroup) {
	out, err := exec.Command("ssh", "-T", "git@github.com").CombinedOutput()
	fmt.Printf("\nls result: %s", string(out))
	if err.Error() == "exit status 1" {
		*isGithubConnected = true
	} else {
		*isGithubConnected = false
	}
	wg.Done()
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
