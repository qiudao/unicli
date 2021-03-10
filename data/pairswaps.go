package main

// Recent Swaps within a Pair

// Get the last 100 swaps on a pair by fetching Swap events and passing in the pair address.
// You'll often want token information as well

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
				swaps(orderBy: timestamp, orderDirection: desc, where: {pair: "0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"})
				{
					pair {
						token0 {
							symbol
						}
						token1 {
							symbol
						}
					}
					amount0In
					amount0Out
					amount1In
					amount1Out
					amountUSD
					to
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
