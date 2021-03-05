#!/bin/sh

curl -g \
		 -X POST \
		 -H "Content-Type: application/json" \
		 -H "Authorization: Bearer 8c7dbd270cb98e83f9d8d57fb8a2ab7bac9d7501905fb013c69995ebf1b2a719" \
		 -d '{"query":"query{uniswapFactories(first:5) { id pairCount totalVolumeUSD} tokens(first: 5) {id symbol name decimals } }" }' \
			https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2
