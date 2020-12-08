package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"time"

	_httpclient "github.com/mfathirirhas/httpclient"
)

func main() {
	// get()
	// getWithURIParams()
	// getWithPathParams()
	// getWithURIParamsAndPathParams()
	// Post()
	// PostJSON()
	PostMultiPart()
}

// get simple get request
func get() {
	resp := _httpclient.Get(nil, &_httpclient.Request{
		BaseURL: "http://example.com",
	})
	str, err := resp.String()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(str)
}

// getWithURIParams get request with uri params
func getWithURIParams() {
	baseURL := "http://localhost:8080/getwithparams"
	urlValues := make(url.Values)
	urlValues.Set("param1", "value1")
	urlValues.Set("param2", "value2")
	urlValues.Set("param3", "value3")
	urlValues.Set("param4", "value4")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp := _httpclient.Get(ctx, &_httpclient.Request{
		BaseURL:   baseURL,
		URLValues: urlValues,
	})
	s := struct {
		Param1 string `json:"param1"`
		Param2 string `json:"param2"`
		Param3 string `json:"param3"`
		Param4 string `json:"param4"`
	}{}
	m := make(map[string]interface{})
	fmt.Println(resp.Header)
	fmt.Println(resp.String())
	if err := resp.Scan(&m); err != nil {
		fmt.Println("error map: ", err)
	}
	fmt.Println("map: ", m)
	if err := resp.Scan(&s); err != nil {
		fmt.Println("error struct: ", err)
	}
	fmt.Println("struct: ", s)
}

// getWithPathParams get request with params in path /:param1
func getWithPathParams() {
	baseURL := "http://localhost:8080/getwithparams/123"
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp := _httpclient.Get(ctx, &_httpclient.Request{
		BaseURL: baseURL,
	})
	fmt.Println(resp.Header)
	fmt.Println(string(resp.Body))
	fmt.Println(resp.String())
}

// getWithURIParamsAndPathParams get request with uri params and path params combined
func getWithURIParamsAndPathParams() {
	baseURL := "http://localhost:8080/getwithuriparamsandpathparams/value1"
	urlValues := make(url.Values)
	urlValues.Set("param2", "value2")
	urlValues.Set("param3", "value3")
	urlValues.Set("param4", "value4")
	urlValues.Set("param5", "value5")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp := _httpclient.Get(ctx, &_httpclient.Request{
		BaseURL:   baseURL,
		URLValues: urlValues,
	})
	s := struct {
		Param1 string `json:"param1"`
		Param2 string `json:"param2"`
		Param3 string `json:"param3"`
		Param4 string `json:"param4"`
		Param5 string `json:"param5"`
	}{}
	m := make(map[string]interface{})
	fmt.Println(resp.Header)
	fmt.Println(resp.String())
	if err := resp.Scan(&m); err != nil {
		fmt.Println("error map: ", err)
	}
	fmt.Println("map: ", m)
	if err := resp.Scan(&s); err != nil {
		fmt.Println("error struct: ", err)
	}
	fmt.Println("struct: ", s)
}

// Post post request using x-www-form payload
func Post() {
	baseURL := "http://localhost:8080/post"
	body := map[string]string{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
		"param4": "value4",
	}
	resp := _httpclient.PostForm(nil, &_httpclient.Request{
		BaseURL: baseURL,
		Body:    body,
	})
	s := struct {
		Param1 string `json:"param1"`
		Param2 string `json:"param2"`
		Param3 string `json:"param3"`
		Param4 string `json:"param4"`
	}{}
	if err := resp.Scan(&s); err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(s)
}

// PostJSON port request using json payload
func PostJSON() {
	baseURL := "http://localhost:8080/postjson"
	body := map[string]string{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
		"param4": "value4",
	}
	resp := _httpclient.PostJSON(nil, &_httpclient.Request{
		BaseURL: baseURL,
		Body:    body,
	})
	s := struct {
		Param1 string `json:"param1"`
		Param2 string `json:"param2"`
		Param3 string `json:"param3"`
		Param4 string `json:"param4"`
	}{}
	if err := resp.Scan(&s); err != nil {
		fmt.Println("error:--> ", err)
	}
	fmt.Println(s)
}

// PostMultiPart post request using multipart payload of non-binary and binary data
// binary data come from directory ./example/from/ uploaded to ./example/to/
// uploaded data should be put into ./example/to/ directory named file1.pdf and file2.jpg.
func PostMultiPart() {
	baseURL := "http://localhost:8080/postmulti"
	body := map[string]string{
		"param1": "value1",
		"param2": "value2",
	}
	path1 := "/Users/mfathirirhas/code/go/src/github.com/mfathirirhas/httpclient/example/from/leis-icde2013.pdf"
	data1, _ := ioutil.ReadFile(path1)
	file1 := _httpclient.File{
		FieldName: "file1",
		FileName:  "file1.pdf",
		Data:      data1,
	}
	path2 := "/Users/mfathirirhas/code/go/src/github.com/mfathirirhas/httpclient/example/from/gopher.jpg"
	data2, _ := ioutil.ReadFile(path2)
	file2 := _httpclient.File{
		FieldName: "file2",
		FileName:  "file2.jpg",
		Data:      data2,
	}
	resp := _httpclient.PostMultipart(nil, &_httpclient.Request{
		BaseURL: baseURL,
		Body:    body,
		Files: []_httpclient.File{
			file1,
			file2,
		},
	})
	if resp.Err() != nil {
		fmt.Println(resp.Err())
		return
	}
	fmt.Println(resp.String())
}
