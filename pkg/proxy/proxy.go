package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Proxy struct{}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	sUrl, err := url.Parse("http://localhost:8081")
	if err != nil {
		fmt.Fprint(rw, err)
	}

	r.Host = sUrl.Host
	r.URL.Host = sUrl.Host
	r.URL.Scheme = sUrl.Scheme
	r.RequestURI = ""

	client := &http.Client{}

	res, err := client.Do(r)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(rw, err)
		return
	}

	for key, values := range res.Header {
		for _, value := range values {
			rw.Header().Set(key, value)
		}
	}

	rw.WriteHeader(res.StatusCode)
	io.Copy(rw, res.Body)
}
