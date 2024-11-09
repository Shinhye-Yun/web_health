package app

import (
	"log"
	"net/http"
	"strings"
	"time"
)

// Add "https://" prefix if missing
func ifHttps(s string) string {
	if strings.HasPrefix(s, "https://") {
		return s
	} else {
		return "https://" + s
	}
}

// Get http status for a given url
func getHttpStatus(url string) {
    var (
        err      error
        res		*http.Response
        retries  int = 30
    )

	if len(url) == 0 {
		log.Printf("[ERROR] GIVEN url IS EMPTY")
		return
	}

	url = ifHttps(url)

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	// GET status code within retry
    for retries > 0 {
		res, err = c.Get(url)
        if err != nil {
            log.Printf("[ERROR] Cannot get status from %s: %v", url, err)
            retries -= 1
        } else {
			//TODO
            break
        }
    }

    if res != nil {
		// Print status code
		status := res.StatusCode
		if status < 200{
			log.Printf("[INFO] %s - Status info %d\n", url, status)
		} else if status >= 200 && status < 300{
			log.Printf("[INFO] %s - Status success %d", url, status)
		} else if status >= 300 && status < 400{
			log.Printf("[INFO] %s - Status redirect %d", url, status)
		} else if status >= 400 && status < 500{
			log.Printf("[INFO] %s - Status client error %d", url, status)
			//TODO send notification
		} else {
			log.Printf("[INFO] %s - Status server error %d", url, status)
			//TODO
		}
    }
}

func CheckStatus(domain string, servers []string){
	for i := 0; i < len(servers); i++ {
		getHttpStatus("https://"+servers[i]+domain)
	}
}