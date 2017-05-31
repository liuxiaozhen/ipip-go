package main

import (
	"math/rand"
	"net/http"
	"runtime"
	"time"

	"github.com/liuxiaozhen/go-ipip/ipip"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ipip.Load("mydata4vipday2.datx")
	http.HandleFunc("/q", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":8085", nil)
}
