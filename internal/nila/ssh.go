package nila

import (
	"encoding/base64"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/crypto/ssh"
)

func fireCmd(config Config, address, username, password, job, cmd string) {
	res, err := sendSSH(address, username, password, cmd)
	var state = "D"
	var errMsg = ""
	if err != nil {
		state = "E"
		errMsg = fmt.Sprintf("%s", err)
	}
	jobPath := filepath.Join(config.JobStore, fmt.Sprintf("%s.job", job))
	resultPath := filepath.Join(config.JobStore, fmt.Sprintf("%s.result", job))
	const tmp = `{"job_id": "%s", "output": "%s","error": "%s", "state": "%s" }`
	resData := base64.StdEncoding.EncodeToString([]byte(res))
	errData := base64.StdEncoding.EncodeToString([]byte(errMsg))
	ret := fmt.Sprintf(tmp, job, resData, errData, state)
	os.WriteFile(resultPath, []byte(ret), 0655)
	os.Remove(jobPath)
}

func sendSSH(address, username, password, cmd string) (string, error) {
	sshconfig := &ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", address, sshconfig)
	if err != nil {
		return "", fmt.Errorf("Failed to dial: %v", err)
	}
	defer conn.Close()

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
