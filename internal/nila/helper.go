package nila

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/denisbrodbeck/machineid"
)

func genMachineID() string {
	did, _ := machineid.ID()
	hash := md5.Sum([]byte(did))
	res := hex.EncodeToString(hash[:])
	return res
}

func getLocalIP(config *Config) net.IP {
	conn, err := net.Dial("tcp", config.OdooBaseURL) // Connect to a public DNS server (Google DNS)
	if err != nil {
		log.Fatal(err) // Handle error if connection fails
	}
	defer conn.Close() // Ensure the connection is closed

	localAddress := conn.LocalAddr().(*net.TCPAddr) // Get the local address
	return localAddress.IP                          // Return the IP address
}

func sendRequest(method, url, apikey string, data []byte) []byte {
	reader := bytes.NewReader(data)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apikey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}
