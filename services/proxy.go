package services

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func DoProxy(w http.ResponseWriter, r *http.Request) {
	requestPath := r.RequestURI
	serviceName := strings.Split(requestPath, "/")[1]
	ipAndPort := GetIpAndPort(serviceName)
	targetUrl := "http://" + ipAndPort + strings.Replace(requestPath, "/"+serviceName, "", 1)
	client := &http.Client{}
	myRequest, err := http.NewRequest(r.Method, targetUrl, r.Body)
	myRequest.Header = r.Header
	myResponse, err := client.Do(myRequest)
	if err != nil {
		return
	}
	w.WriteHeader(myResponse.StatusCode)
	_, err = io.Copy(w, myResponse.Body)
	if err != nil {
		log.Printf("在复制响应内容时出错: %v", err)
	}
	err = myRequest.Body.Close()
	if err != nil {
		return
	}
	err = myResponse.Body.Close()
	if err != nil {
		return
	}
}
