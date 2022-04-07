package media

import (
	"net/http"

	"github.com/hi-bridge-9/privacy-sandbox/lib/media"
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
		const adtechURL = 'http://localhost/ad_tech'
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

	<p id="user-agent">
    </p>
	<script>
		var userAgentResult = navigator.userAgent;
		document.getElementById("user-agent").innerHTML = userAgentResult;
	</script>
</body>
</html>
`

func Handler(w http.ResponseWriter, r *http.Request) {
	m := media.New(w)
	if r.Method != "GET" {
		m.Response(http.StatusMethodNotAllowed, nil, nil)
		return
	}

	headers := map[string]string{
		"Content-Type": "text/html",
		"Accept-CH":    "Sec-CH-UA-Reduced", // for User Agent Reduction Origin Trials
	}

	m.Response(http.StatusOK, headers, []byte(topPage))
}
