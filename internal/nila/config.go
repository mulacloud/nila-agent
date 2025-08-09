package nila

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type Config struct {
	OdooHost     string
	OdooBaseURL  string
	OdooPort     int
	APIKey       string
	PullInterval int
	AgentID      string
	JobStore     string
	Hoster       []Hoster
}

type Hoster struct {
	Address     string `json:"address"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	StatCommand string `json:"stat_command"`
}

func InitConfig() Config {
	host := flag.String("host", "", "Odoo Host")
	port := flag.Int("port", 0, "Odoo Port")
	jobPath := flag.String("jobstore", "/tmp/nila", "Nila Job store path")
	flag.Parse()

	if *host == "" {
		log.Panic("Odoo Host not set")
	}

	if *port == 0 {
		log.Panic("Odoo port not set")
	}

	api, err := os.LookupEnv("APIKEY")
	if err == false {
		log.Panic("APIKEY Environment variable not set")
	}

	os.Mkdir(*jobPath, 0755)

	odooBaseURL := fmt.Sprintf("%s:%d", *host, *port)

	config := Config{
		OdooHost:    *host,
		OdooPort:    *port,
		OdooBaseURL: odooBaseURL,
		APIKey:      api,
		JobStore:    *jobPath,
	}

	return config
}

func (config Config) GetRegisterURL() string {
	return fmt.Sprintf("http://%s/nila/agent/register", config.OdooBaseURL)
}

func (config Config) GetPullURL() string {
	return fmt.Sprintf("http://%s/nila/job/pull", config.OdooBaseURL)
}
