#!/bin/sh

curl -g \
		 -X POST \
		 -H "Content-Type: application/json" \
		 -d '{"query":"query{uniswapFactories(first: 5) { id pairCount totalVolumeUSD totalVolumeETH untrackedVolumeUSD totalLiquidityUSD totalLiquidityETH txCount} }" }' \
			https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2 | jq
