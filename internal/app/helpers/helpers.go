package helpers

import (
	"bytes"
	"io"
	"log"
	"net/http"
)


func ReadRequestBody(r *http.Request) *bytes.Buffer {
	
	b := new(bytes.Buffer)
	_, err := io.Copy(b, r.Body)

	if err != nil {
		log.Fatal(err)
	}

	return b
}

func ReadResponseBody(r *http.Response) *bytes.Buffer {
	
	b := new(bytes.Buffer)
	_, err := io.Copy(b, r.Body)

	if err != nil {
		log.Fatal(err)
	}

	return b
}
