package main

import (
	"log"

	"github.com/asadhayat1068/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8086")
	log.Fatal(srv.ListenAndServe())
}
