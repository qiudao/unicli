package main

import (
		"fmt"
		"strings"
		"net/http"
		"os"
		"io/ioutil"
	   )

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <query-string>\n", os.Args[0])
		return
	}
	qs := fmt.Sprintf("{\"query\":\"query%s \"}", os.Args[1])
	httpPost(qs)
}

func httpPost(query string) {
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
