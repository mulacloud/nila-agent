package main

import (
	"time"

	"mula.cloud/nila-agent/internal/nila"
)

func main() {
	config := nila.InitConfig()
	nila.RegisterAgent(&config)
	for {
		nila.PullJob(config)
		time.Sleep(time.Duration(config.PullInterval) * time.Second)
	}
	// cmd := "vmadm list"
	// out, err := nila.SendSSH(config, cmd)
	// log.Printf("Command out : %s => Error %s", out, err)
}
