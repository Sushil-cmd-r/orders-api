package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(check),
	}

	fmt.Println("starting server...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func check(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
