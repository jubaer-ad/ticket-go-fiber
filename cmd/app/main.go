package main

import (
	"fmt"
	"net/http"

	"github.com/ticket-go/routes"
)

func main() {
	router := routes.NewRouter()
	port := 4000
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server is listening on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}
