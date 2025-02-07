package main

import "github.com/okppop/url-shortener/app"

func main() {
	a := &app.App{}
	a.Init("./config.yaml")
	a.Start()
}
