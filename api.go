package main

import (
	"net/http"
)

var jsonConf *JsonConf

func init() {
	http.HandleFunc("/json/save", SaveJsonHandler)
	http.HandleFunc("/json/get", GetJsonHandler)
}

// SaveJsonHandler  保存配置
func SaveJsonHandler(w http.ResponseWriter, r *http.Request) {
	var queryParams JsonConf
	err := ParseBodyParam(r, &queryParams)
	if err != nil {
		WrapResponse(w,nil,http.StatusBadRequest)
	}
	jsonConf = &queryParams
	WrapResponse(w,jsonConf,http.StatusOK)
}

// GetJsonHandler  获取配置
func GetJsonHandler(w http.ResponseWriter, r *http.Request) {
	WrapResponse(w,jsonConf,http.StatusOK)
}


