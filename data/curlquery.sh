#!/bin/sh

curl -g \
		 -X POST \
		 -H "Content-Type: application/json" \
		 -d '{"query":"query{uniswapFactories(first:5) { id pairCount totalVolumeUSD} tokens(first: 5) {id symbol name decimals } }" }' \
			https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2
