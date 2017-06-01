package main

import (
	"net/http"
	"runtime"

	"github.com/liuxiaozhen/ipip-go/ipip"
)

const (
	LISTEN_PORT string = ":8085"
	DATA_FILE   string = "E:/mygo/src/github.com/liuxiaozhen/ipip-go/data/mydata4vipday2.datx"
)

var (
	oip *ipip.Ipip
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	oip := ipip.NewIpip()
	oip.Load(DATA_FILE)

	http.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {
		ipstr := r.FormValue("ip")
		ipinfo, err := oip.Find(ipstr)
		if err != nil {
			w.Write([]byte("Oops, some error occurred!"))
		}
		w.Write([]byte(ipinfo))
	})
	http.ListenAndServe(LISTEN_PORT, nil)
}
