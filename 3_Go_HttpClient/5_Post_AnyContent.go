package main
// curl -T main.go -H "Content-Type: text/plain" http://localhost:18888

import (
	"log"
	"net/http"
		"os"
)
func main() {
	file, err := os.Open("5_Post_AnyContent.go")
	if err != nil {
		panic(err)
	}

	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		panic(err)
	}
	log.Println("status: ", resp.Status)
}
