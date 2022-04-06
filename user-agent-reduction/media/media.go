package media

import (
	"log"
	"net/http"
)

var topPage = `
	<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>publisher</title>
    </style>
</head>
<body>
    <h1>UserAgent-Reduction メディアサイト</h1>
    <div id="ad">
    </div>
    <script>
		const adtechDomain = 'localhost'
		const adtechOrigin = 'http://' + adtechDomain
		const adtechURL = adtechOrigin + '/ad_tech'

		fetch(adtechURL)
			.then(
				response => response.json()
			)
			.then(
				data => {
					if (data != null && data != "" && data != undefined) {
						var target = document.getElementById("ad");
						if (data.ads != null && data.ads != "" && data.ads != undefined) {
							target.innerHTML = data.ads;
						}
					}
					console.log(data);
				}
			)
			.catch(
				error => console.log(error)
			);
    </script>
</body>
</html>
`

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Invalid request method: %v\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Add("Content-Type", "text/html")
	w.Header().Add("Accept-CH", "Sec-CH-UA-Reduced")   // for User Agent Reduction
	w.Header().Add("Critical-CH", "Sec-CH-UA-Reduced") // for User Agent Reduction
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(topPage))
}
