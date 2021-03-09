package main

import (
		"net/http"
	   )

// Client is a client for interacting with GraphQL API.
type Client struct {
	endpoint		string
	httpClient		*http.Client
	useMultipartForm bool
	closeReq		bool
	Log	func(s string)
}

