package main

import (
	"net/http"
	"os"

	"github.com/maxence-charriere/go-app/v6/pkg/app"
	"github.com/maxence-charriere/go-app/v6/pkg/log"
)

func main() {
//	port := os.Getenv("PORT")
	port := os.Args[1]
	if port == "" {
		port = "7777"
	}

	version := os.Getenv("GAE_VERSION")

	log.Info("starting server").
		T("port", port).
		T("version", version)

	h := &app.Handler{
		Title:  "Attractions Tips",
		Author: "Marlon",
		Styles: []string{
			"/web/hello.css",
			"/web/city.css",
		},
		CacheableResources: []string{
			"/web/space.jpg",
			"/web/beijing.jpg",
			"/web/paris.jpg",
			"/web/sf.jpg",
		},
		Version: version,
	}

	err := http.ListenAndServe(":"+port, h)
	if err != nil {
		log.Error("server crashed").T("error", err)
	}
}
