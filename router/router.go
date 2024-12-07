package router

import (
	"main/router/utils"
	"main/router/utils/endpoint"
	"net/http"
)

type Router struct {
	Path       string
	Middleware utils.LinkedList
}

func (R *Router) Listen(v func(http.ResponseWriter, *http.Request), middleware ...func(http.ResponseWriter, *http.Request) bool) {
	http.HandleFunc(R.Path, func(rw http.ResponseWriter, r *http.Request) {
		if middleware[0] != nil {
			v(rw, r)
			return
		} else {
			for i := 0; i < len(middleware); i++ {
				if !middleware[i](rw, r) {
					return
				}
			}
			v(rw, r)
		}
	})
}

func (R *Router) Use(E *endpoint.Endpoint) {
	curConnection := E.Connections.Head

	for curConnection != nil {
		http.HandleFunc(R.Path+E.Path+curConnection.Path, curConnection.Value)

		if curConnection.Next != nil {
			curConnection = curConnection.Next
		} else {
			return
		}
	}
}
