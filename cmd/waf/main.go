package main

import (
	"fmt"
	"net/url"
	"os"

	"wafy/internal/proxy"
)

func main() {
	googleUrl, _ := url.Parse("https://google.com")

	reverseProxy := proxy.NewReverseProxy(":8080", *googleUrl)

	if err := reverseProxy.RunHTTP(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}
