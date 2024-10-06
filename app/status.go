package app

import (
    "log"
    "net/http"
	"time"
	"strings"
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
func GetHttpStatus(url string) {
    var (
        err      error
        res		*http.Response
        retries  int = 30
    )

	if len(url) == 0 {
		log.Println("GIVEN url IS EMPTY")
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
            log.Println(url, err)
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
			log.Println(url, "info", status)
		} else if status >= 200 && status < 300{
			log.Println(url, "success", status)
		} else if status >= 300 && status < 400{
			log.Println(url, "redirect", status)
		} else if status >= 400 && status < 500{
			log.Println(url, "client error", status)
			//TODO send notification
		} else {
			log.Println(url, "server error", status)
			//TODO
		}
    }
}