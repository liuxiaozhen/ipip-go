package main

import (
	"fmt"
	//	"net/http"
	"runtime"

	"github.com/liuxiaozhen/ipip-go/ipip"
)

const (
	LISTEN_PORT string = ":8085"
	DATA_FILE   string = "E:/mygo/data/mydata4vipday2.dat"
)

var (
	oip *ipip.Ipip
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	oip := ipip.NewIpip()
	oip.Load(DATA_FILE)
	var ipstr string = "202.106.187.173"

	ipinfo, err := oip.Find(ipstr)
	if err != nil {
		fmt.Println("Oops, some error occurred!")
		fmt.Println(err)
		return
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
