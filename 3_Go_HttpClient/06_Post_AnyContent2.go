package main

// curl -T main.go -H "Content-Type: text/plain" http://localhost:18888

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("TEXT TEXT TEXT")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		panic(err)
	}
	log.Println("status:", resp.Status)
}
