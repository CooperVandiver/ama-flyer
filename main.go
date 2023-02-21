package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func servePDF(w http.ResponseWriter, r *http.Request) {
	filename := "data/dandelionflyer.pdf"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	defer f.Close()

	w.Header().Set("Content-type", "application/pdf")

	if _, err := io.Copy(w, f); err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
}

func main() {
	listenAddr := "0.0.0.0:" + os.Getenv("PORT")
	http.HandleFunc("/dandelion-and-co-flyer", servePDF)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		fmt.Println(err)
	}
}
