package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	configFileLoc := flag.String("config", "./config.yaml", "Path to your config file")
	flag.Parse()

	config := loadConfig(*configFileLoc)

	log.Println((time.Now()).Format("Mon Jan 2 15:04:05 MST 2006"), " Checking Status")
	for _, service := range config.Services {
		go checkForDowntime(service)
	}
	for range time.Tick(time.Duration(config.Interval) * time.Second) {
		log.Println((time.Now()).Format("Mon Jan 2 15:04:05 MST 2006"), " Checking Status")
		for _, service := range config.Services {
			go checkForDowntime(service)
		}
	}
}

// checkForDowntime does a GET request and compares resp status code to expected code.
func checkForDowntime(service Service) {
	resp, err := http.Get(service.URL)
	if err != nil {
		log.Println(service.Name, " Error: ", err)
		sendAlert(service)
	} else {
		if resp.StatusCode != service.RespCode {
			log.Println(service.Name, " Expected Resp Code ", service.RespCode, " got ", resp.StatusCode)
			sendAlert(service)
		}
		// No need to log if everything is ok.
	}

}

// sendAlert if service is down.
func sendAlert(service Service) {
	if !service.Alert {
		return
	}

	postParams := url.Values{}

	for key, val := range service.AlertFormParams {
		postParams.Add(key, val)
	}

	_, err := http.PostForm(service.AlertWebHook, postParams)
	if err != nil {
		log.Println(service.Name, " Error while sending alert ", err)
	}
}
