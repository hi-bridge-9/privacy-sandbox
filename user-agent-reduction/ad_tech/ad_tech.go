package ad_tech

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hi-bridge-9/privacy-sandbox/lib/ad_tech"
)

var (
	baseAdTag   = `<a href=%s>%s</a>`
	baseImagTag = `<img src=%s width=%d height=%d>`
	baseResp    = `{"ads": "%s"}`
)

func Handler(w http.ResponseWriter, r *http.Request) {
	at := ad_tech.New(w)
	if r.Method != "GET" {
		at.Response(http.StatusMethodNotAllowed, nil, nil)
		return
	}

	path := strings.Split(r.URL.Path, "/")
	if path[2] == "image" {
		img, err := at.ReadAdImage("./ad_tech/image/pop_wadai_sns.png")
		if err != nil {
			at.Response(http.StatusInternalServerError, nil, nil)
		}

		headers := map[string]string{
			"Content-Type": "image/png",
		}
		at.Response(http.StatusOK, headers, img)
		return
	}

	imgTag := fmt.Sprintf(baseImagTag, "./ad_tech/image", 450, 450)
	adTag := fmt.Sprintf(baseAdTag, "https://www.apple.com/jp/", imgTag)
	resp, err := at.ConvertToJSON(baseResp, adTag)
	if err != nil {
		at.Response(http.StatusInternalServerError, nil, nil)
		return
	}

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	at.Response(http.StatusOK, headers, resp)
}
