package main

import (
	"docker/Company"
	"docker/http_server"
	"log"
	"net/http"
)

func main() {
	company := Company.NewCompany()
	hadler := http_server.NewHandler(company)
	mux := http.NewServeMux()

	mux.HandleFunc("POST /employees", hadler.AddHTTP)
	mux.HandleFunc("GET /employees", hadler.GetWorkersHTTP)
	mux.HandleFunc("DELETE /employees/{id}", hadler.RemoveHTTP)
	if err := http_server.StartHTTPServer(mux); err != nil {
		log.Fatal(err)
	}
}
