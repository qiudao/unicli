all: pairswap

factory:
	@curl -g \
		 -X POST \
		 -H "Content-Type: application/json" \
		 -d '{"query":"query{uniswapFactories(first: 5) { id pairCount totalVolumeUSD totalVolumeETH untrackedVolumeUSD totalLiquidityUSD totalLiquidityETH txCount} }" }' \
			https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2 | jq


token:
	go run token.go graphql.go
	
tokenall:
	go run tokenall.go graphql.go

pairswap:
	go run pairswaps.go graphql.go

pairall:
	go run pairall.go graphql.go

swap:
	go run swap.go graphql.go

pairday:
	go run pairday.go graphql.go

CHI:
	./query "{tokens(id: \"0x00cb166e39a77a693146a13536f72f5bef689511\") { id name symbol decimals}}" 

test:
	./query "{uniswapFactories(first: 5) { id pairCount totalVolumeUSD totalVolumeETH untrackedVolumeUSD totalLiquidityUSD totalLiquidityETH txCount}}" | jq
	./query "{tokens(first: 5) { id symbol name decimals}}" | jq
