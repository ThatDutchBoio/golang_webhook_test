package main

import (
	"fmt"
	"main/api/getwebhooks"
	"main/api/registerwebhook"
	"main/api/sendmessage"
	"main/api/test"
	"main/apikeys"
	"main/db/dbcontroller"
	"main/router"
	"net/http"
)

func main() {
	fmt.Println(apikeys.GenerateApiKey())
	entries, err := dbcontroller.GetAllWebhooks()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(entries)
	}
	// HomePage

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/public", http.StatusFound)
		fmt.Fprint(w, "Home page")

	})
	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	})

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	// test API
	TestEndpoint := test.GetTestAPIEndpoint()
	MessageEndpoint := sendmessage.GetEndpoint()
	RegisterWebhookEndpoint := registerwebhook.GetEndPoint()
	GetWebhooksEndpoint := getwebhooks.GetEndpoint()
	Router2 := router.Router{Path: "/api"}
	Router2.Use(&TestEndpoint)
	Router2.Use(&RegisterWebhookEndpoint)
	Router2.Use(&MessageEndpoint)
	Router2.Use(&GetWebhooksEndpoint)

	// router.Testing()
	http.ListenAndServe(":8000", nil)
}
