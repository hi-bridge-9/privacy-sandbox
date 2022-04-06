package main

import (
	"log"
	"net/http"

	"github.com/hi-bridge-9/privacy-sandbox/user-agent-reduction/ad_tech"
	"github.com/hi-bridge-9/privacy-sandbox/user-agent-reduction/media"
	"github.com/hi-bridge-9/privacy-sandbox/util"
)

var (
	handlerFuncMap = map[string]func(w http.ResponseWriter, r *http.Request){
		"/media":   media.Handler,
		"/ad_tech": ad_tech.Handler,
		"/image": ad_tech.Handler,
	}
)

func main() {
	s := util.NewWebServer(handlerFuncMap)
	if err := s.Start("80"); err != nil {
		log.Fatal(err)
	}
}
