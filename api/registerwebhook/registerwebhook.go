package registerwebhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"main/apikeys"
	"main/db/dbcontroller"
	"main/router/utils/endpoint"
	"main/utils/hashing"
	"net/http"
)

type RegisterBody struct {
	Credentials string `json:"credentials"`
	Host        string `json:"host"`
}

func GetEndPoint() endpoint.Endpoint {

	APIEndpoint := endpoint.Endpoint{Path: "/registerwebhook"}

	// register webhook
	APIEndpoint.Listen("", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(body)) // Reset the body for reuse
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var payload RegisterBody
		fmt.Println(json.Unmarshal(body, &payload))
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		fmt.Println(payload.Credentials)
		fmt.Println(payload.Host)
		nKey := apikeys.GenerateApiKey()
		hashedCredentials := hashing.SHA256(payload.Credentials)
		dbcontroller.RegisterWebhook(nKey, hashedCredentials, payload.Host)

		fmt.Fprintf(w, "Successfully registered webhook with ID: %s", nKey)
	})

	return APIEndpoint
}
