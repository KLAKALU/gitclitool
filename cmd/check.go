package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func checkGithubConnection(fileDir FileDirectory) bool {

	// sshでぃれくとりが存在するか確認

	var wg sync.WaitGroup

	stopChan := make(chan struct{})

	wg.Add(1)

	go loadingAnimation(stopChan, &wg)

	knownHostsCheck(fileDir)

	isGithubConnection := gettingGithubUserName()

	close(stopChan)
	wg.Wait()
	return isGithubConnection
}

func loadingAnimation(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	marks := []string{"   ", ".  ", ".. ", "..."}
	for i := 0; i < 500; i++ {
		select {
		case <-stopChan:
			return
		default:
			fmt.Printf("\r%s", marks[i%4])
			time.Sleep(250 * time.Millisecond)
		}
	}
}

func knownHostsCheck(fileDir FileDirectory) {
	fmt.Println("check known_hosts")
	// check known_hosts exist
	if _, err := os.Stat(filepath.Join(fileDir.homeDir, fileDir.distDir, "known_hosts")); os.IsNotExist(err) {
		// ~/.ssh/known_hosts file not exist
		makeKnownHosts(fileDir)
	} else {
		// ~/.ssh/known_hosts file exist
		fmt.Println("Done")
	}
}

func makeKnownHosts(fileDir FileDirectory) {
	fmt.Println("make known_hosts")
	out, err := exec.Command("ssh-keyscan", "github.com").CombinedOutput()
	if err != nil {
		fmt.Println("failed to make known_hosts list")
		os.Exit(1)
	}
	if f, err := os.Stat(filepath.Join(fileDir.homeDir, fileDir.distDir, "known_hosts")); os.IsNotExist(err) || !f.IsDir() {
		err := os.Mkdir(filepath.Join(fileDir.homeDir, fileDir.distDir), 0755)
		if err != nil {
			fmt.Println("failed to make .ssh directory")
			os.Exit(1)
		}

	}

	f, err := os.Create(filepath.Join(fileDir.homeDir, fileDir.distDir, "known_hosts"))
	if err != nil {
		fmt.Println("failed to make known_hosts file")
		os.Exit(1)
	}
	defer f.Close()
	f.Write(out)
	fmt.Println("Done")
}

func gettingGithubUserName() bool {
	fmt.Println("get github username")
	out, err := exec.Command("ssh", "-T", "git@github.com").CombinedOutput()
	if err != nil {
		if out != nil {
			string := string(out)
			strList := strings.Split(string, " ")
			if strList[1] == "Permission" {
				return false
			}
			if strList[0] == "git@github.com:" {
				fmt.Println("failed to connect to github")
			}
			if strList[0] == "ssh:" {
				fmt.Println("failed to connect to github")
				fmt.Println("please check network")
				os.Exit(1)
			}
			userName := strList[1]
			userName = strings.Replace(userName, "!", "", 1)
			fmt.Println("\rgithub username: " + userName)
			return true
		} else {
			fmt.Println("failed to get github username")
			return false
		}
	} else {
		fmt.Println("failed to get github username")
		return false
	}
}
