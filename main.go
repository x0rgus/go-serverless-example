package main

import (
	"log"
	"net/http"

	"github.com/x0rgus/go-serverless-example/hardware"
)

func main() {
	http.HandleFunc("/devices", hardware.DevicesHandler)
	log.Println("Servidor iniciado na porta :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
