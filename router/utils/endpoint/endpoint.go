package endpoint

import (
	// "main/router"

	"main/router/utils"
	"net/http"
)

type Endpoint struct {
	Path        string
	Connections utils.EndpointLinkedList
}

func (E *Endpoint) Listen(path string, v func(http.ResponseWriter, *http.Request)) {
	E.Connections.Add(path, v)
}
