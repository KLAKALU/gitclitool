package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func try_login_github(DIST_DIR string, SSH_KEY_NAME string, homeDir string) {

	const USER = "git"

	const HOST = "github.com"

	key, err := os.ReadFile(filepath.Join(homeDir, DIST_DIR, SSH_KEY_NAME))
	if err != nil {
		fmt.Println("failed to read private key")
		os.Exit(1)
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		fmt.Println("failed to parse private key")
		os.Exit(1)
	}

	config := &ssh.ClientConfig{
		User: USER,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	_, err = ssh.Dial("tcp", HOST+":22", config)
	if err != nil {
		fmt.Println("failed to connect to GitHub")
		os.Exit(1)
	}

	fmt.Println("Connected to GitHub")
}
