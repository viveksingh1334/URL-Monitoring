package main

import (
	"fmt"
	"net/http"
	"time"
)

type URLStatus struct {
	URL     string
	Status  int
	Success bool
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.twitter.com",
		"https://www.github.com",
		"https://www.linkedin.com",
	}

	// Create a channel to receive URLStatus
	statusChan := make(chan URLStatus)

	// Start monitoring URLs concurrently
	for _, url := range urls {
		go monitorURL(url, statusChan)
	}

	// Wait for the monitoring to complete
	for i := 0; i < len(urls); i++ {
		status := <-statusChan
		fmt.Printf("URL: %s, Status Code: %d, Success: %t\n", status.URL, status.Status, status.Success)
	}
}

func monitorURL(url string, statusChan chan<- URLStatus) {
	for {
		status := URLStatus{URL: url}

		// Send an HTTP GET request to the URL
		resp, err := http.Get(url)
		if err != nil {
			status.Success = false
			fmt.Println("Error:", err)
		} else {

			status.Status = resp.StatusCode
			status.Success = resp.StatusCode >= 200 && resp.StatusCode < 300

			resp.Body.Close()
		}

		// Send the status to the channel
		statusChan <- status

		// Wait for 5 seconds before checking the URL again
		time.Sleep(5 * time.Second)
	}
}
