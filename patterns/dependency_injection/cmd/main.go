package main

import (
	di "dependency_injection/di"
	"log"
	"net/http"
)

func main() {
	sds := di.NewSimpleDataStore()
	l := di.NewSimpleLogic(sds, di.LoggerAdapter(di.LogOutput))
	c := di.NewController(di.LoggerAdapter(di.LogOutput), l)
	http.HandleFunc("/hello", c.SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
