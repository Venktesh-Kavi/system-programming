package task

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/api/loan/details", loanDetailsHandler)
	http.ListenAndServe(":8080", mux)
}

func loanDetailsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&loanDetails)
}
