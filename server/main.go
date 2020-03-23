package main

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

func main() {
	h := &app.Handler{
		Title:  "akhr2tag",
		Author: "p1gd0g",
		Icon:   app.Icon{Default: "/web/icon.jpg"},
		RawHeaders: []string{`<!-- The core Firebase JS SDK is always required and must be listed first -->
		<script src="https://www.gstatic.com/firebasejs/7.12.0/firebase-app.js"></script>
		
		<!-- TODO: Add SDKs for Firebase products that you want to use
			 https://firebase.google.com/docs/web/setup#available-libraries -->
		<script src="https://www.gstatic.com/firebasejs/7.12.0/firebase-analytics.js"></script>
		
		<script>
		  // Your web app's Firebase configuration
		  var firebaseConfig = {
			apiKey: "AIzaSyAtUtdV-7cUFbwmp3oq7rqNQi3AxjeT1-s",
			authDomain: "quickstart-1554221638006.firebaseapp.com",
			databaseURL: "https://quickstart-1554221638006.firebaseio.com",
			projectId: "quickstart-1554221638006",
			storageBucket: "quickstart-1554221638006.appspot.com",
			messagingSenderId: "792826622257",
			appId: "1:792826622257:web:f106c15b737d1dfb5d8010",
			measurementId: "G-GVLCS2STV1"
		  };
		  // Initialize Firebase
		  firebase.initializeApp(firebaseConfig);
		  firebase.analytics();
		</script>`},
	}

	if err := http.ListenAndServeTLS(":443", "p1gd0g_com.crt", "key.pem", h); err != nil {
		panic(err)
	}
}
