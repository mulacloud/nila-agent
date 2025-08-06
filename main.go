package main

import "log"

func main() {
	cmd := "vmadm list"
	out, err := sendSSH(cmd)
	log.Printf("Command out : %s => Error %s", out, err)
}
