package ad_tech

import (
	"bytes"
	"errors"
	"fmt"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
)

type AdTech struct {
	resp http.ResponseWriter
}

func New(w http.ResponseWriter) *AdTech {
	return &AdTech{
		resp: w,
	}
}

func (a AdTech) ReadAdImage(fp string) ([]byte, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("invalid file path: %w", err)
	}

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed decode file data: %v", err)
	}

	buff := new(bytes.Buffer)
	if err := jpeg.Encode(buff, img, nil); err != nil {
		return nil, errors.New("unable to encode image")
	}
	return buff.Bytes(), nil
}

func (a *AdTech) ConvertToJSON(base, ad string) ([]byte, error) {
	return []byte(fmt.Sprintf(base, ad)), nil
}

func (a *AdTech) Response(status int, headers map[string]string, body []byte) {
	switch {
	case status >= 500:
		a.ResponseErr(status, headers)
	case status >= 400:
		a.ResponseErr(status, headers)
	case status >= 300:
		a.ResponseRedirect(status, headers)
	case status >= 200:
		a.ResponseSuccess(status, headers, body)
	}
}

func (a *AdTech) ResponseSuccess(status int, headers map[string]string, body []byte) {
	for k, v := range headers {
		a.resp.Header().Add(k, v)
	}
	a.resp.WriteHeader(status)
	a.resp.Write(body)
	log.Printf("Response success: %v\n", status)
}

func (a *AdTech) ResponseErr(status int, headers map[string]string) {
	for k, v := range headers {
		a.resp.Header().Add(k, v)
	}
	a.resp.WriteHeader(status)
	log.Printf("Response failed: %v\n", status)
}

func (a *AdTech) ResponseRedirect(status int, headers map[string]string) {
	for k, v := range headers {
		a.resp.Header().Add(k, v)
	}
	a.resp.WriteHeader(status)
	log.Printf("Response redirect: %v\nLocation: %v\n", status, headers["Location"])
}
