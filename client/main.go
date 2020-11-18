package main

import (
  "github.com/bitterlox/tradepanel/client/backend"
  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

// TODO: move Backend declaration here so it doesn't have long import path

type Server struct {
  backend.Backend
}


func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  server := &Server{*backend.NewBackend()}

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1024,
    Height: 768,
    Title:  "go-broker",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(server)
  app.Run()
}