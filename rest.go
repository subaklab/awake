package awake

import (
"fmt"
)


type Item struct {}

const (
GET = "GET"
POST = "POST"
PUT = "PUT"
DELETE = "DELETE"
HEAD = "HEAD"
)

type Rest struct {
	mux *http.ServeMux
	muxInitialize bool
}


func NewAPI() *Rest {
	return &Rest{}
}

func (api *Rest) requestHandler(resource interface{} http.HandlerFunc{
	return func(rw http.ResponseWriter, request *http.Request){
		if request.ParseForm() != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return 
		}
		var handler func(url.Values, http.Header) (int, interface{}, http.Header)
	}
	})