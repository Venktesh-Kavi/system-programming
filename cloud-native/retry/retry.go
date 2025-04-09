package retry

import (
	"errors"
	"log"
	"math"
	"net/http"
	"slices"
	"time"
)

type RetryConfig struct {
	interval           time.Duration
	maxRetries         int
	backOffCoefficient float64
}

type RetryClient struct {
	client  http.Client
	config  RetryConfig
	request http.Request
}

var DefaultRetry RetryConfig = RetryConfig{
	interval:           time.Second * 10,
	maxRetries:         3,
	backOffCoefficient: 2,
}

var clientSideStatusCodes = []int{400, 401, 403, 404, 409, 429}
var serverSideStatusCodes = []int{500, 502, 503, 504}

func contains(code int, codes []int) bool {
	return slices.Contains(codes, code)
}

func PerformReq(client http.Client, req *http.Request, config RetryConfig) (*http.Response, error) {
	ch := make(chan *http.Response, 1)
	for idx := range config.maxRetries {
		if idx > 0 {
			// Calculate backoff duration for retries after first attempt
			backoffDuration := config.interval * time.Duration(int64(math.Pow(config.backOffCoefficient, float64(idx-1))))
			log.Printf("Waiting %v before retry attempt %d", backoffDuration, idx+1)
			time.Sleep(backoffDuration)
		}

		log.Printf("Attempt %d/%d with config: %+v", idx+1, config.maxRetries, config)
		go makeReq(client, req, ch)

		// Wait for response or timeout
		select {
		case resp := <-ch:
			if resp.StatusCode == http.StatusOK {
				return resp, nil
			} else if contains(resp.StatusCode, clientSideStatusCodes) {
				// Don't retry on client-side errors
				return resp, nil
			} else if contains(resp.StatusCode, serverSideStatusCodes) {
				// Continue retrying on server-side errors
				continue
			}
			// For any other status code, continue retrying
			continue
		case <-time.After(config.interval): // Add reasonable timeout
			// Timeout occurred, continue to next retry
			log.Printf("timer fired, making request - retry count: %d\n", idx+1)
			continue
		}

	}

	return nil, errors.New("max retries exceeded")
}

func makeReq(client http.Client, req *http.Request, ch chan *http.Response) (*http.Response, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	ch <- resp
	return resp, err
}
