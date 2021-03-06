package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	// create a client (safe to share across requests)
	client := NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")

	// make a request
	req := NewRequest(`    
				query {
				pairs(first: 1000, orderBy: reserveUSD, orderDirection: desc) {
				 id
				}
				}
		`)

	// set header fields
	req.Header.Set("Cache-Control", "no-cache")

	// define a Context for the request
	ctx := context.Background()

	// run it and capture the response
	var responseData map[string]interface{}
	if err := client.Run(ctx, req, &responseData); err != nil {
		log.Fatal(err)
	}

	for _, v := range responseData {
		value := v.([]interface{})
		for i, u := range value {
			fmt.Println(i, u)
		}
	}
}
