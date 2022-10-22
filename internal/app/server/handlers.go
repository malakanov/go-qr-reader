package server

import (
	"bytes"
	"fmt"
	"go-qr-reader/internal/app/libs"
	"io"
	"log"
	"net/http"
)

func ScanHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b := new(bytes.Buffer)
	if _, err := io.Copy(b, r.Body); err != nil {
		msg := fmt.Sprintf("Failed to read request body: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	res, err := libs.Scan(b.Bytes())
	if err != "" {
		msg := fmt.Sprintf("Internal server error: %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	_, err2 := w.Write([]byte(res))
	if err2 != nil {
		log.Fatal(err2)
	}
}
