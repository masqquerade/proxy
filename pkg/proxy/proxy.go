package proxy

import "net/http"

type Proxy struct{}

func (p *Proxy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

}
