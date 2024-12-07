package getwebhooks

import (
	"encoding/json"
	"fmt"
	"main/db/dbcontroller"
	"main/router/utils/endpoint"
	"net/http"
)

func GetEndpoint() endpoint.Endpoint {
	var GetWebhooksEndpoint endpoint.Endpoint = endpoint.Endpoint{Path: "/getwebhooks"}
	GetWebhooksEndpoint.Listen("", func(w http.ResponseWriter, r *http.Request) {

		// send webhooks to client in json format
		entries, err := dbcontroller.GetAllWebhooksSanitary()
		if err != nil {
			http.Error(w, "Failed to get webhooks", http.StatusBadRequest)
			return
		}
		// json.NewEncoder(w).Encode(entries)
		jData, err := json.Marshal(entries)
		if err != nil {
			http.Error(w, "Failed to encode json", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Println(entries)
		// send json object backto client as response
		w.Write(jData)

	})

	return GetWebhooksEndpoint
}
