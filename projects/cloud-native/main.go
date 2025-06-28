package main

import (
	"cloud-native/retry"
	"log"
	"net/http"
	"net/url"
	"time"
)

func main() {
	req := http.Request{
		Method: "GET",
		URL:    &url.URL{Scheme: "http", Host: "https://run.mocky.io/v3/b9072d6c-3eae-414c-a622-a4e732833f73"},
	}

	client := http.Client{
		Timeout:   time.Second * 10,
		Transport: http.DefaultTransport,
	}

	resp, err := retry.PerformReq(client, &req, retry.DefaultRetry)
	if err != nil {
		panic(err)
	}
	log.Printf("received response status code: %d\n", resp.StatusCode)
}
