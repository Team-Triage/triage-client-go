package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func makeRequestStruct(triageNetworkAddress, grpcPort, authToken string) http.Request {
	reqHeaders := http.Header{"grpcport": []string{grpcPort}, "Authorization": []string{authToken}}
	req := http.Request{Method: "GET", URL: &url.URL{Scheme: "http", Host: triageNetworkAddress, Path: "/consumers"}, Header: reqHeaders}
	return req
}

func RequestConnection(triageNetworkAddress string, grpcPort string, authToken string) *http.Response {
	httpClient := http.Client{}
	request := makeRequestStruct(triageNetworkAddress, grpcPort, authToken)
	res, err := httpClient.Do(&request)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	if !(res.StatusCode >= 200 && res.StatusCode < 400) {
		b, _ := ioutil.ReadAll(res.Body)
		log.Fatalln(res.Status, "TRIAGE CLIENT: Unable to connect to Triage:", string(b))
	}

	return res
}
