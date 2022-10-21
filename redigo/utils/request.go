package utils

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"io"
	"log"
	"net/http"
	"time"
)

func GenURL(host string, port string, path string) string {
	return "http://" + host + ":" + port + path
}

func RequestGet(url string, params map[string]any) (resp *http.Response, err error) {
	client := &http.Client{Timeout: 5 * time.Second}
	url += "?"
	for k, v := range params {
		url += fmt.Sprintf("%v=%v&", k, v)
	}
	resp, err = client.Get(url)
	return resp, err
}

func RequestGetData(url string, params map[string]any) map[string]any {
	resp, err := RequestGet(url, params)
	if err != nil {
		log.Fatal(err)
	}
	return ReadResponse2Map(resp)
}

func GetResponseBuffer(resp *http.Response) *bytes.Buffer {
	defer resp.Body.Close()
	var buffer [512]byte
	buf := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		buf.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}
	return buf
}

func ReadResponse2String(resp *http.Response) string {
	buf := GetResponseBuffer(resp)
	return buf.String()
}

func ReadResponse2Map(resp *http.Response) map[string]any {
	buf := GetResponseBuffer(resp)
	data := make(map[string]any)
	err := json.Unmarshal(buf.Bytes(), &data)
	if err != nil {
		panic(err)
	}
	return data
}

func ReadResponse2Obj(resp *http.Response, obj any) {
	buf := GetResponseBuffer(resp)
	err := json.Unmarshal(buf.Bytes(), obj)
	if err != nil {
		panic(err)
	}
}
