package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func sendSSH(cmd string) (string, error) {
	// Replace with your SSH server details
	sshAddress := "192.168.0.11:22"
	username := "root"
	password := "m" // Or use a private key for authentication

	config := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		// InsecureHostKeyCallback is for development only; for production,
		// you should verify the host key.
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", sshAddress, config)
	if err != nil {
		return "", fmt.Errorf("Failed to dial: %v", err)
	}
	defer conn.Close()

	// Example: Run a command
	session, err := conn.NewSession()
	if err != nil {
		return "", fmt.Errorf("Failed to create session: %v", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", fmt.Errorf("Failed to run command: %v", err)
	}
	return string(output), nil
}
