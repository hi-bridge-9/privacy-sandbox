package ad_tech

import (
	"log"
	"net/http"
)

var ad = `
`

type image struct {
	src    string
	width  int
	height int
}

type adModel struct {
	href  string
	id    string
	image image
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(topPage))
}

// func genAds() string {

// 	id, _ := uuid.NewRandom()
// 	return &adModel{
// 		href: "https://trust-token.d2c-ts1.com/click?lp=" + lp,
// 		id:   id.String(),
// 		image: image{
// 			src:    "https://js.d2c-ts1.com/image/pop_wadai_nosyouhin.png",
// 			width:  450,
// 			height: 450,
// 		},
// 	}
// 	return
// }
