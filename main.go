package main

import (
	"log"
	"net/http"

	"goleaky/builder"
	"goleaky/validator"

	_ "net/http/pprof"
)

func main() {
	go doWork()
	log.Println(http.ListenAndServe("localhost:6060", nil))
}

func doWork() {
	for {
		report := builder.BuildReport()
		validator.ValidateAndSave(report)
	}
}
