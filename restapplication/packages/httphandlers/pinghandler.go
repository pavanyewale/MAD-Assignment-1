package httphandlers

import (
	"fmt"
	"log"
	"net/http"
	mthdroutr "pavan/MAD-Assignment-1/restapplication/packages/mthdrouter"
	"pavan/MAD-Assignment-1/restapplication/packages/resputl"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("serveHTTP")
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (p *PingHandler) Get(r *http.Request) resputl.SrvcRes {
	s := r.URL.Query()
	key, ok := s["key"]
	if ok {
		fmt.Println(key[0])
	}
	return resputl.Response200OK("OK")
}
