package main

import (
	"github.com/golovers/prometheus/my-restapi/service"
	"log"
	"net/http"
)

func main() {

	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
