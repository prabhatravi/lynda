package main

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func upTown(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	var dogName string
	fs := strings.Split(req.URL.Path, "/")
	if len(fs) >= 2 {
		dogName = fs[1]
	}
	// the image doesn't serve
	io.WriteString(res, `
	<h1>Dog Name: `+dogName+`</h1><br>
	<img src="/toby.jpg">
	`)
}

func dogPic(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(res, "file not found", 404)
		return
	}

	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)
}

func main() {
	http.HandleFunc("/", upTown)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":9000", nil)
}

/*
All material is licensed under the Apache License Version 2.0, January 2004
http://www.apache.org/licenses/LICENSE-2.0
*/
