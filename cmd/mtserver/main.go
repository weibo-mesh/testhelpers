package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/weibocom/motan-go"
	mc "github.com/weibocom/motan-go/core"
)

func startMotanServer() {
	motanServerCtx := motan.GetMotanServerContext("./simple-server.yaml")
	motanServerCtx.RegisterService(&MTService{}, "")
	motanServerCtx.Start(nil)
	motanServerCtx.ServicesAvailable()
}

// test HTTPServer
func startHTTPServer() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "http server provider by golang.")
	})
	http.ListenAndServe(":8080", serverMux)
}

func main() {
	go startMotanServer()
	startHTTPServer()
}

// MTService for type test Server
type MTService struct{}

// Hello test func for string map
func (mt *MTService) Hello(args map[string]string) string {
	b := mc.NewBytesBuffer(1000)
	x := b.Bytes()
	b.WriteZigzag64(64)
	b.WriteZigzag64(1)
	b.WriteZigzag32(64)
	b.WriteZigzag32(1)
	f := b.Bytes()
	return fmt.Sprintf("%+v-------%+v", x, f)
}

// HelloW test func for string map
func (mt *MTService) HelloW(args map[string]string) string {
	return "HelloW"
}

func pow(x, n int64) (ret int64) {
	ret = 1
	for n != 0 {
		if n%2 != 0 {
			ret = ret * x
		}
		n /= 2
		x = x * x
	}
	return ret
}

// HelloX test func for multi args
func (mt *MTService) HelloX(strArg string, inT64 int64, inT32 int64, nameArr []string) string {
	return fmt.Sprintf("strArg:%s-inT64:%d-int32:%d-%+v", strArg, inT64, inT32, nameArr)
}

// HelloY test func for multi args
func (mt *MTService) HelloY(strArg string, nameArr []string) string {
	return fmt.Sprintf("strArg:%s-%+v", strArg, nameArr)
}

// HelloZ test func for multi args
func (mt *MTService) HelloZ(strArg string, inT64 int64, inT32 int64, nameArr []string) int64 {
	return 92233720
	// return 9223372036854775808
}

// HelloZ32 test func for multi args
func (mt *MTService) HelloZ32(strArg string, inT64 int64, inT32 int64, nameArr []string) int32 {
	return 2222222
	// return 9223372036854775808
}

// TimeOutErr test
func (mt *MTService) TimeOutErr(strMap map[string]string) (res string) {
	time.Sleep(5 * time.Second)
	return "Time Out Test."
}
