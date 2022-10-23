package server

import (
	"encoding/json"
	"fmt"
	"go-qr-reader/internal/app/helpers"
	"go-qr-reader/internal/app/libs"
	"io/ioutil"
	"log"
	"net/http"
)

type QrUrl struct {
	Url string `json:"url"`
}

func FileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b := helpers.ReadRequestBody(r)

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

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	reqBody, _ := ioutil.ReadAll(r.Body)
	var qrUrl QrUrl
	json.Unmarshal(reqBody, &qrUrl)

	fmt.Println("Send request by url: ", qrUrl.Url)

	resp, err := http.Get(qrUrl.Url)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Got response, with status code: %d\n", resp.StatusCode)

	b := helpers.ReadResponseBody(resp)

	res, scanError := libs.Scan(b.Bytes())
	if scanError != "" {
		msg := fmt.Sprintf("Internal server error: %v", scanError)
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	_, writeError := w.Write([]byte(res))
	if writeError != nil {
		msg := "Internal server error"
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}
