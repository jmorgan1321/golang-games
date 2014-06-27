// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomeHandler")
}

func serveSingle(pattern string, filename string) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filename)
	})
}

func main() {
	// Normal resources
	http.Handle("/", http.FileServer(http.Dir("./static/")))

	// // Mandatory root-based resources
	// serveSingle("/sitemap.xml", "./sitemap.xml")
	// serveSingle("/favicon.ico", "./favicon.ico")
	// serveSingle("/robots.txt", "./robots.txt")

	done := make(chan bool)
	go func() {
		http.ListenAndServe(":8080", nil)
		done <- true
	}()

	cmd := exec.Command("cmd", "/c", "start", "http://localhost:8080")
	cmd.Run()
	<-done
}
