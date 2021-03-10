package main

import (
		"fmt"
		"strings"
		"net/http"
		"io/ioutil"
		"context"
		"log"
	   )

func main2() {
	//qs := `{"query":"query{pair(id: "0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"){ id} }"}`
	//qs := `{"query":"query{tokens(first: 1){ id symbol} }"}`
	qs := `{ "query":"query{pair(id: "0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"){id reserve0} }"  }`
	//qs := `{"query":"query{pair(id: "0xa478c2975ab1ea89e8196811f51a7b7ade33eb11"){ id } }"}`
	httpPost(qs)
}

func httpPost(query string) {
	fmt.Println(query)
	resp, err := http.Post("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2",
			"Content-Type: application/json",
			strings.NewReader(query))

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}

func main() {
// create a client (safe to share across requests)
client := NewClient("https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2")

// make a request
req := NewRequest(`    
		query ($key: String!) {
		pair(id:$key) {
		id
		reserve0
		}
		}
		`)


		// set any variables
		req.Var("key", "0xa478c2975ab1ea89e8196811f51a7b7ade33eb11")

		// set header fields
		req.Header.Set("Cache-Control", "no-cache")

		// define a Context for the request
		ctx := context.Background()

		// run it and capture the response
		var responseData map[string]interface{}
		if err := client.Run(ctx, req, &responseData); err != nil {
			log.Fatal(err)
		}

		for k := range responseData {
			fmt.Printf("%s: %s\n", k, responseData[k])
		}
}
