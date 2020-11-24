package main

import (
	"github.com/bitterlox/tradepanel/client/backend"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"log"
)

// TODO: move Backend declaration here so it doesn't have long import path

type Server struct {
	backend.Backend
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	b, err := backend.NewBackend()
	if err != nil {
		log.Fatal("could not connect: ", err)
	}

	server := &Server{*b}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "go-broker",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(server)
	_ = app.Run()
}
