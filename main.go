package main

import "github.com/okppop/url-shortener/app"

func main() {
	a := &app.App{}

	err := a.Init("./config.yaml")
	if err != nil {
		panic(err)
	}

	a.Start()
}
