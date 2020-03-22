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
	}

	if err := http.ListenAndServeTLS(":443", "p1gd0g_com.crt", "key.pem", h); err != nil {
		panic(err)
	}
}
