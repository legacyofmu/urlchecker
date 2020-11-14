package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	c := make(chan result)

	urls := []string{
		"https://www.airbnb.com",
		"https://www.google.com",
		"https://www.amazon.com",
		"https://www.reddit.com",
		"https://soundcloud.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) { // chan<- : send only channel
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		status = "FAILED"
	}

	c <- result{url: url, status: status}
}
