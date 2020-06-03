package main

import (
	"github.com/maxence-charriere/go-app/v6/pkg/app"
)

const (
	icon = "https://raw.githubusercontent.com/marlonflying/picturesRepo/master/attractions-192.png"
)

func main() {
	app.Route("/", &hello{})
	app.Route("/city", &city{})
	app.Run()
}
