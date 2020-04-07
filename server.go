package main

import (
	_ "fmt"
	"log"
	"io"
	"net/http"

	"caracal-pty/xterm"
)

func ptyHandler(w http.ResponseWriter, r *http.Request) {
	xterm.ShellWs(w, r)
}

func home(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "It's home path.")
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/pty", ptyHandler)

	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
