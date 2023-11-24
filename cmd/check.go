/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
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

		isgithubconnected := false

		wg.Add(1)

		go roadinganimation(&wg)

		go tryconnecttogithub(&isgithubconnected, &wg)

		wg.Wait()

		if isgithubconnected {
			fmt.Println("connected to github successfully!")
		} else {
			fmt.Println("failed to connect to github!")
		}
	},
}

func roadinganimation(wg *sync.WaitGroup) {
	marks := []string{"   ", ".  ", ".. ", "..."}
	for i := 0; i < 500; i++ {
		fmt.Printf("\rconecting to github %s", marks[i%4])
		time.Sleep(250 * time.Millisecond)
	}
	wg.Done()
}

func tryconnecttogithub(isgithubconnected *bool, wg *sync.WaitGroup) {
	time.Sleep((3000 * time.Millisecond))
	*isgithubconnected = true
	wg.Done()
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
