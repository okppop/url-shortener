package main

import (
	"fmt"

	"github.com/okppop/url-shortener/conf"
)

func main() {
	a, err := conf.Read("./assets/config.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", a)
	fmt.Println(a.HTTPServer.GetListenAddress())
	fmt.Println(a.Postgresql.GetDSN())
}
