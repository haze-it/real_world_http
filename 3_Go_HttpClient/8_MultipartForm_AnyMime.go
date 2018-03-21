package main

// curl -F "name=haze" -F "thumnail=@haze_icon.png" http://localhost:18888

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "haze")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="haze_icon.png"`)

	fileWriter, err := writer.CreatePart(part)
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
