package media

import (
	"log"
	"net/http"
)

type Media struct {
	resp http.ResponseWriter
}

func New(w http.ResponseWriter) *Media {
	return &Media{
		resp: w,
	}
}

func (m *Media) Response(status int, headers map[string]string, body []byte) {
	switch {
	case status >= 500:
		m.ResponseErr(status, headers)
	case status >= 400:
		m.ResponseErr(status, headers)
	case status >= 300:
		m.ResponseRedirect(status, headers)
	case status >= 200:
		m.ResponseSuccess(status, headers, body)
	}
}

func (m *Media) ResponseSuccess(status int, headers map[string]string, body []byte) {
	for k, v := range headers {
		m.resp.Header().Add(k, v)
	}
	m.resp.WriteHeader(status)
	m.resp.Write(body)
	log.Printf("Response success: %v\n", status)
}

func (m *Media) ResponseErr(status int, headers map[string]string) {
	for k, v := range headers {
		m.resp.Header().Add(k, v)
	}
	m.resp.WriteHeader(status)
	log.Printf("Response failed: %v\n", status)
}

func (m *Media) ResponseRedirect(status int, headers map[string]string) {
	for k, v := range headers {
		m.resp.Header().Add(k, v)
	}
	m.resp.WriteHeader(status)
	log.Printf("Response redirect: %v\nLocation: %v\n", status, headers["Location"])
}
