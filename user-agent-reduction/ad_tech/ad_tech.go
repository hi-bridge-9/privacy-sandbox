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
	"strings"

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

	if len(strings.Split(r.Host, "/")) < 1 {
		img, err := readFile("./ad_tech/image/pop_wadai_sns.png")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.Header().Add("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(img)
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

func readFile(fp string) ([]byte, error) {
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
