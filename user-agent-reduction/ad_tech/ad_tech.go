package ad_tech

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

var (
	baseAdTag   = `<a href=%s id=%s>%s</a>`
	baseImagTag = `<img src=%s width=%d height=%d>`
	baseResp    = `{"ads": "%s"}`
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	resp := makeResp()

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(resp))
}

func makeResp() string {
	id, _ := uuid.NewRandom()
	return fmt.Sprintf(baseResp, genAds(id.String()))
}

func genAds(id string) string {
	imgTag := fmt.Sprintf(baseImagTag, "./image", 450, 450)
	return fmt.Sprintf(baseAdTag, "https://www.apple.com/jp/", id, imgTag)
}
