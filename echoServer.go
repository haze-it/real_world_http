// Real World HTTP
//  Page: 4-5

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// handler
func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}

	// リクエストの内容を出力する
	fmt.Println(string(dump))

	// HTML出力
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

// main
func main() {
	var httpServer http.Server
	// "/"にアクセスがあったらhandlerが呼ばれる
	http.HandleFunc("/", handler)
	log.Println("start http listening :18888")
	// Port: 18888
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}
