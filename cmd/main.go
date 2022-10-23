package main

import (
	"fmt"
	"log"
	"net/http"

	// internal
	"go-qr-reader/internal/app/server"

	// import gif, jpeg, png
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	// import bmp, tiff, webp
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"
)

var (
	addr string
)

func init() {
	addr = server.NewConfig()
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/file", server.FileHandler)
	mux.HandleFunc("/url", server.UrlHandler)
	httpServer := &http.Server{Addr: addr, Handler: mux}
	fmt.Println("Server running on: ", addr)
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
