package main

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func main() {
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
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	log.Println("SSH connection established successfully!")

	// Example: Run a command
	session, err := conn.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	output, err := session.CombinedOutput("ls -l")
	if err != nil {
		log.Fatalf("Failed to run command: %v", err)
	}
	log.Printf("Command output:\n%s", output)
}
