package sendmessage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/db/dbcontroller"
	"main/router/utils/endpoint"
	"net/http"
)

type MessageFormat struct {
	Message string `json:"message"`
}

func GetEndpoint() endpoint.Endpoint {
	var MessageEndpoint endpoint.Endpoint = endpoint.Endpoint{Path: "/sendmessage"}

	MessageEndpoint.Listen("", func(v http.ResponseWriter, r *http.Request) {

		body, err := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the body for reuse
		if err != nil {
			http.Error(v, "Failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var payload MessageFormat
		fmt.Println(json.Unmarshal(body, &payload))
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(v, "Invalid JSON format", http.StatusBadRequest)
		}

		fmt.Println(payload.Message)

		AllWebhooks, err := dbcontroller.GetAllWebhooks()
		if err == nil {
			postBody, _ := json.Marshal(map[string]string{
				"message": payload.Message,
			})
			for i := 0; i < len(AllWebhooks); i++ {
				fmt.Println(AllWebhooks[i].Host)
				http.Post(AllWebhooks[i].Host, "application/json", bytes.NewBuffer(postBody))
			}
		}
	})

	return MessageEndpoint
}
