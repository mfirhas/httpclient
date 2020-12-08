package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	_httpserver "github.com/mfathirirhas/httpserver"
)

func main() {
	srv := _httpserver.New(&_httpserver.Opts{
		Port:         8080,
		EnableLogger: true,
	})
	srv.GET("/getwithparams", GetHandlerWithParams)
	srv.GET("/getwithparams/:param1", GetHandlerWithPathParams)
	srv.GET("/getwithuriparamsandpathparams/:param1", GetHandlerWithURIParamsAndPathParams)
	srv.POST("/post", Post)
	srv.POST("/postjson", PostJSON)
	srv.POST("/postmulti", PostMultiPart)
	srv.Run()
}

func GetHandlerWithParams(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"param1": r.URL.Query().Get("param1"),
		"param2": r.URL.Query().Get("param2"),
		"param3": r.URL.Query().Get("param3"),
		"param4": r.URL.Query().Get("param4"),
	}
	_httpserver.ResponseJSON(w, r, http.StatusOK, resp)
}

func GetHandlerWithPathParams(w http.ResponseWriter, r *http.Request) {
	resp := r.URL.Query().Get("param1")
	_httpserver.ResponseString(w, r, http.StatusOK, resp)
}

func GetHandlerWithURIParamsAndPathParams(w http.ResponseWriter, r *http.Request) {
	resp := map[string]interface{}{
		"param1": r.URL.Query().Get("param1"),
		"param2": r.URL.Query().Get("param2"),
		"param3": r.URL.Query().Get("param3"),
		"param4": r.URL.Query().Get("param4"),
		"param5": r.URL.Query().Get("param5"),
	}
	_httpserver.ResponseJSON(w, r, http.StatusOK, resp)
}

func Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp := map[string]interface{}{
		"param1": r.FormValue("param1"),
		"param2": r.FormValue("param2"),
		"param3": r.FormValue("param3"),
		"param4": r.FormValue("param4"),
	}
	_httpserver.ResponseJSON(w, r, http.StatusOK, resp)
}

func PostJSON(w http.ResponseWriter, r *http.Request) {
	resp := struct {
		Param1 string `json:"param1"`
		Param2 string `json:"param2"`
		Param3 string `json:"param3"`
		Param4 string `json:"param4"`
	}{}
	if err := json.NewDecoder(r.Body).Decode(&resp); err != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err)
		return
	}
	fmt.Println(resp)
	_httpserver.ResponseJSON(w, r, http.StatusOK, resp)
}

func PostMultiPart(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err)
		return
	}
	file1, header1, err1 := r.FormFile("file1")
	defer file1.Close()
	if err1 != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err1)
		return
	}

	file2, header2, err2 := r.FormFile("file2")
	defer file2.Close()
	if err2 != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err2)
		return
	}

	dest1 := "/Users/mfathirirhas/code/go/src/github.com/mfathirirhas/httpclient/example/to/" + header1.Filename
	f1, err := os.Create(dest1)
	defer f1.Close()
	if err != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err)
		return
	}
	if _, err := io.Copy(f1, file1); err != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err)
		return
	}

	dest2 := "/Users/mfathirirhas/code/go/src/github.com/mfathirirhas/httpclient/example/to/" + header2.Filename
	f2, err := os.Create(dest2)
	defer f2.Close()
	if err != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err)
		return
	}
	if _, err := io.Copy(f2, file2); err != nil {
		_httpserver.ResponseString(w, r, http.StatusInternalServerError, err)
		return
	}

	_httpserver.ResponseString(w, r, http.StatusOK, "Success upload")
}
