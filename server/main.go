package main

import (
	"fmt"
	//	"net/http"
	"runtime"

	"github.com/liuxiaozhen/ipip-go/ipip"
)

const (
	LISTEN_PORT string = ":8085"
	DATA_FILE   string = "E:/mygo/data/mydata4vipday2.datx"
)

var (
	oip *ipip.Ipipx
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	oip := ipip.NewIpipx()
	oip.Load(DATA_FILE)
	var ipstr string = "123.125.26.232"

	ipinfo, err := oip.Find(ipstr)
	if err != nil {
		fmt.Println("Oops, some error occurred!")
	}

	fmt.Println(ipinfo)
	ipjson := ipip.JsonLocationInfo(ipinfo)
	fmt.Println(ipjson)

	/**
	http.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		ipstr := r.FormValue("ip")
		ipinfo, err := oip.Find(ipstr)
		if err != nil {
			w.Write([]byte("Oops, some error occurred!"))
		}
		w.Write([]byte(ipinfo))
	})
	http.ListenAndServe(LISTEN_PORT, nil)
	*/
}
