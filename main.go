package main

import (
	_ "checkgo/conf"
	_ "checkgo/generator"
	"log"
	"net/http"
)

const (
	defaultWebHostStr   = "127.0.0.1:8091"
	defaultWebPrefixStr = "/"
	defaultWebStaticDir = "./static/"
)

func main() {
	http.Handle(defaultWebPrefixStr, http.StripPrefix(defaultWebPrefixStr, http.FileServer(http.Dir(defaultWebStaticDir))))
	if err := http.ListenAndServe(defaultWebHostStr, nil); err != nil {
		log.Fatal(err)
	}
}
