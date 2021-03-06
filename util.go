package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)


// ParseBodyParam 按json格式解析请求体body
func ParseBodyParam(r *http.Request, param interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	// Reset resp.body so it can be use again
	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	err = json.Unmarshal(body, param)
	if err != nil {
		return err
	}
	return nil
}

// WrapResponse 将interface{}转为json写入http.ResponseWriter
func WrapResponse(w http.ResponseWriter, data interface{}, httpCode int) {

	result := map[string]interface{}{"data":data}
	jsonData, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpCode)
	w.Write(jsonData)
}
