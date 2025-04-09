package retry

import "testing"

func TestPerformRetry(t *testing.T) {

}

// I should be able to mock the makeReq call with a random response, as I dont want to make a http call.
// Timers, channels & waiting can be still run via tests
// I should see if it is 5xx status code, I am retrying after an interval
// If it is 4xx status code, I am not retrying
// If it is 2xx status code, I am not retrying
