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
				pairDayDatas(first: 100, orderBy: date, orderDirection: asc, 
					where: {
						pairAddress:"0xa478c2975ab1ea89e8196811f51a7b7ade33eb11", 
						date_gt: 1592505859
						} ) {
					date
					dailyVolumeToken0
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
