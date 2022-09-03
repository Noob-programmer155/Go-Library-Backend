package main

import (
	"amrDev/libraryBackend.com/repository"
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", "8000"),
		Handler: mux,
	}

	go func() {
		panic(server.ListenAndServe())
	}()

	repository.InitDB(server)
}
