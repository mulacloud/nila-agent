package nila

import (
	"encoding/base64"
	"fmt"
	"log"

	"github.com/tidwall/gjson"
)

func RegisterAgent(config *Config) {

	mid := genMachineID()
	ip := getLocalIP(config)
	const tmp = `{"params": {"name": "%s", "ip": "%s"}}`
	jsonData := []byte(fmt.Sprintf(tmp, mid, ip))
	var hosterArray []Hoster

	body := sendRequest("POST", config.GetRegisterURL(), config.APIKey, jsonData)
	pull_interval := gjson.GetBytes(body, "result.pull_interval").Int()
	hoster := gjson.GetBytes(body, "result.hoster")
	code := gjson.GetBytes(body, "result.code").String()
	if code == "201" {
		log.Printf("New agent %s registered successfuly", mid)
	} else {
		log.Printf("Agent %s already registered, continue operation", mid)

		hoster.ForEach(func(key, value gjson.Result) bool {
			address := gjson.Get(value.String(), "address").String()
			username := gjson.Get(value.String(), "username").String()
			password := gjson.Get(value.String(), "password").String()
			statCommand := gjson.Get(value.String(), "stat_command").String()
			decodedStat, _ := base64.StdEncoding.DecodeString(statCommand)
			hs := Hoster{
				Address:     address,
				Username:    username,
				Password:    password,
				StatCommand: string(decodedStat)}
			hosterArray = append(hosterArray, hs)
			return true
		})
	}
	log.Printf("Detect total %d hoster", len(hosterArray))
	log.Printf("Setting pull interval to = %d seconds", pull_interval)
	config.PullInterval = int(pull_interval)
	config.AgentID = mid
	config.Hoster = hosterArray
}
