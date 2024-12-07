package test

import (
	"fmt"
	"main/router/utils/endpoint"
	"net/http"
)

func GetTestAPIEndpoint() endpoint.Endpoint {

	APIEndpoint := endpoint.Endpoint{Path: "/test"}

	APIEndpoint.Listen("/", func(v http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/test/")
	})

	APIEndpoint.Listen("/testing1", func(v http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/test/testing1")
	})

	APIEndpoint.Listen("/testing1/test", func(v http.ResponseWriter, r *http.Request) {
		fmt.Println("/api/testing/testing1/test")
	})

	return APIEndpoint
}
