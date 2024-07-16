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

	//var wg sync.WaitGroup

	//wg.Add(1)

	//go loadingAnimation(&wg)

	knownHostsCheck(fileDir)

	gettingGithubUserName()

	//wg.Wait()
	return true
}

func loadingAnimation(wg *sync.WaitGroup) {
	marks := []string{"   ", ".  ", ".. ", "..."}
	for i := 0; i < 500; i++ {
		fmt.Printf("\rconecting to github %s", marks[i%4])
		time.Sleep(250 * time.Millisecond)
	}
	wg.Done()
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
	f, err := os.Create(filepath.Join(fileDir.homeDir, fileDir.distDir, "known_hosts"))
	if err != nil {
		fmt.Println("failed to make known_hosts file")
		os.Exit(1)
	}
	defer f.Close()
	f.Write(out)
	fmt.Println("Done")
}

func gettingGithubUserName() {
	fmt.Println("get github username")
	out, err := exec.Command("ssh", "-T", "git@github.com").CombinedOutput()
	if err != nil {
		if out != nil {
			string := string(out)
			strList := strings.Split(string, " ")
			if strList[0] == "git@github.com:" {
				fmt.Println("failed to connect to github")
				os.Exit(1)
			}
			if strList[0] == "ssh:" {
				fmt.Println("failed to connect to github")
				fmt.Println("please check network")
				os.Exit(1)
			}
			userName := strList[1]
			userName = strings.Replace(userName, "!", "", 1)
			fmt.Println("github username: " + userName)
		} else {
			fmt.Println("failed to get github username")
		}
	} else {
		fmt.Println("failed to get github username")
	}
}
