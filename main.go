package main

import (
	"example.com/pingui/app"
	"time"
)

func main() {
	// TODO: pass in env var or read from config
	domain := ".staging.mist-federal.com"
	web_servers := []string{"manage"}
	servers := []string{"ztp","ep-terminator","nac-terminator"}
	servers = append(web_servers, servers...)

	keepRunning := true
	counter := 0

	for keepRunning{
		if counter > 10 {
			break
		}
		app.CheckStatus(domain, web_servers)
		app.CheckCert(domain, servers)
		counter += 1
		time.Sleep(10 * time.Second) 
	}
}
