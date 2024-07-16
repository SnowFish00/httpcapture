package sdefault

import (
	"context"
	"fmt"
	"ids/packet"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	eventCh = make(chan interface{}, 1024)
	//test
	eventUCh    = make(chan interface{}, 1024)
	ctx, cancel = context.WithCancel(context.Background())
)

func StartCapture() {
	go shutdown()
	//test
	srv := packet.NewPacketHandle(ctx, "en0", eventCh, eventUCh)
	srv.SetBpf("tcp port 80")
	//test
	srv.SetBpfU("tcp and not port 80 or udp")
	srv.SetEventHandle(5, handle) // option
	//test
	srv.SetEventHandleU(5, handleU)
	srv.SetPromisc(true)          // option
	srv.SetFlushTime(time.Minute) // option
	go func() {
		if err := srv.Listen(); err != nil {
			log.Println(err.Error())
		}
	}()
	go func() {
		if err := srv.CaptureExp80(); err != nil {
			log.Println(err.Error())
		}
	}()

}

func handle(req *http.Request, resp *http.Response, doc map[string]interface{}) {
	//test
	reqBody, _ := ioutil.ReadAll(req.Body)
	respBody, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	defer resp.Body.Close()

	// 创建一个map来存储请求和响应的信息
	httpData := make(map[string]string)

	// 填充map
	httpData["Request Headers"] = fmt.Sprintf("%v", req.Header)
	httpData["Request Body"] = string(reqBody)
	httpData["Response Headers"] = fmt.Sprintf("%v", resp.Header)
	httpData["Response Body"] = string(respBody)

	doc["HttpPayload"] = httpData

	fmt.Println(doc)

}

// test
func handleU(doc map[string]interface{}) {
	fmt.Println("doc", doc)
}

// test
func shutdown() {
	time.Sleep(time.Hour * 1)
	fmt.Println("------time enough------")
	cancel()
}
