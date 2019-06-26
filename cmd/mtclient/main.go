package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/weibocom/motan-go"
	"github.com/weibocom/motan-go/core"
)

var motanClient *motan.Client
var videSizeMap = map[string][]byte{"1k": nil, "5k": nil, "128k": nil, "512k": nil, "1M": nil, "10M": nil, "20M": nil}
var sizeMap = map[string]int{"1k": 1024, "5k": 1024 * 5, "128k": 1024 * 128, "512k": 1024 * 512, "1M": 1024 * 1024, "10M": 1024 * 1024 * 10, "20M": 1024 * 1024 * 20}

func video(w http.ResponseWriter, r *http.Request) {

	var reply string
	reqSize := r.URL.Query()["s"][0]
	err := motanClient.Call("JustForT", []interface{}{map[string][]byte{"cc": nil, "s": videSizeMap[reqSize]}}, &reply)
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Fprintf(w, "Hello, %q,%s", reply, reqSize)
}

func t(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "hello, %q", html.EscapeString(r.URL.Path))
}

func main() {
	defer core.HandlePanic(nil)
	motanClientContext := motan.GetClientContext("./simple-client.yaml")
	motanClientContext.Start(nil)
	motanClient = motanClientContext.GetClient("test-video")
	f, _ := os.Open("/dev/urandom")
	for sizeStr, _ := range videSizeMap {
		readData := make([]byte, sizeMap[sizeStr])
		f.Read(readData)
		videSizeMap[sizeStr] = readData
	}
	f.Close()
	http.HandleFunc("/bar", video)
	http.HandleFunc("/t", t)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
