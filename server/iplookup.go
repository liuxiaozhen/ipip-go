package main

import (
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/liuxiaozhen/ipip-go/ipip"
)

const (
	LISTEN_PORT string = ":8085"
	DATA_FILE   string = "E:/mygo/data/mydata4vipday2.datx" //请自行修改
	LOG_FILE    string = "E:/mygo/log/iplookup.log"         //使用标准库的log.请自行修改
)

var (
	oip *ipip.IpipX
)

func logdebug(str string) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)
	log.Println(str)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logFile, err := os.Create(LOG_FILE)
	defer logFile.Close()
	if err != nil {
		log.Fatalln("open file error !")
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC)

	oip := ipip.NewIpipX()
	oip.Load(DATA_FILE)

	if err != nil {
		log.Fatalln("load data error | " + err.Error())
		return
	}

	http.HandleFunc("/iplookup", func(w http.ResponseWriter, r *http.Request) {
		ipstr := r.FormValue("ip")
		ipinfo, err := oip.Find(ipstr)
		ipjson := ipip.JsonLocationInfo(ipinfo)
		if err != nil {
			w.Write([]byte("Oops, some error occurred!"))
		}
		w.Write([]byte(ipjson))
		log.Println(ipstr + "||" + ipjson)
	})
	http.ListenAndServe(LISTEN_PORT, nil)
}
