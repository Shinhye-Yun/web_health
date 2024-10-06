package main

import (
	"example.com/pingui/app"
	"time"
)

func main() {
	url := []string{"manage.staging.mist-federal.com"}
	keepRunning := true
	counter := 0
	for keepRunning{
		if counter > 10 {
			break
		}
		for i := 0; i < len(url); i++ {
			app.GetHttpStatus("https://"+url[i])
			app.ValidateCert(url[i], "443")
		}
		counter += 1
		time.Sleep(10 * time.Second) 
	}
}