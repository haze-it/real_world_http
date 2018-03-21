package main
// curl -F "name=haze" -F "thumnail=@haze_icon.png" http://localhost:18888

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "haze")

	fileWriter, err := writer.CreateFormFile("thumbnail", "haze_icon.png")
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("haze_icon.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	writer.Close()

	// writer.FromDataContentType() =="multipart/form-data; boundary=" + writer.Boundary()
	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		panic(err)
	}
	log.Println("status: ", resp.Status)

}
