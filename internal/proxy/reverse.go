package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

const defaultPattern = "/"

type Reverse struct {
	destination url.URL
	host        string
}

func NewReverseProxy(host string, destination url.URL) *Reverse {
	return &Reverse{
		destination: destination,
		host:        host,
	}
}

func (rp *Reverse) RunHTTP() error {
	rProxy := httputil.NewSingleHostReverseProxy(&rp.destination)

	http.Handle(defaultPattern, makeHandle(rProxy))

	return http.ListenAndServe(rp.host, nil)
}

func makeHandle(p *httputil.ReverseProxy) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		w.Header().Set("X-Ben", "Rad")
		p.ServeHTTP(w, r)
	}
}
