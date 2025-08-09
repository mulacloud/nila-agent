package nila

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tidwall/gjson"
)

func PullJob(config Config) {
	const restmp = `{"params": {"name": "%s", "result": [%s], "stat": [%s]}}`
	var resArray []string
	var statArray []string

	resFiles, _ := filepath.Glob(fmt.Sprintf("%s/*.result", config.JobStore))
	for _, fl := range resFiles {
		cont, _ := os.ReadFile(fl)
		resArray = append(resArray, string(cont))
		os.Remove(fl)
	}

	for _, hs := range config.Hoster {
		statRE, _ := sendSSH(hs.Address, hs.Username, hs.Password, hs.StatCommand)
		stats := strings.Split(statRE, "\n")
		if len(stats) > 0 {
			stats = stats[:len(stats)-1]
		}
		for _, st := range stats {
			statArray = append(statArray, st)
		}
	}

	resData := fmt.Sprintf(
		restmp,
		config.AgentID,
		strings.Join(resArray, ","),
		strings.Join(statArray, ","),
	)

	body := sendRequest("POST", config.GetPullURL(), config.APIKey, []byte(resData))

	jobs := gjson.GetBytes(body, "result.job")

	jobs.ForEach(func(key, value gjson.Result) bool {
		jobName := gjson.Get(value.String(), "name").String()
		address := gjson.Get(value.String(), "hoster.address").String()
		username := gjson.Get(value.String(), "hoster.username").String()
		password := gjson.Get(value.String(), "hoster.password").String()
		shell := gjson.Get(value.String(), "shell").String()
		jobPath := filepath.Join(config.JobStore, fmt.Sprintf("%s.job", jobName))
		os.WriteFile(jobPath, []byte(value.String()), 0655)

		go fireCmd(config, address, username, password, jobName, shell)

		return true // keep iterating
	})

	log.Printf("Pulling ==> %d jobs || Reporting ==> %d job || Stat ==> %d zones",
		len(jobs.Array()), len(resArray), len(statArray),
	)
}
