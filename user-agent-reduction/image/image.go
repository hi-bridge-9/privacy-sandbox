package image

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

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// main.goから見た画像のパス
	img, err := readFile("./image/pop_wadai_sns.png")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Add("Content-Type", "image/png")
	w.WriteHeader(http.StatusOK)
	w.Write(img)
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
