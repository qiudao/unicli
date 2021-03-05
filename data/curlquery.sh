#!/bin/sh

if [ $# -ne 1 ]; then 
 echo "Usage: $0 <query-json>"
 exit
fi

echo "{"query":"query$1" }" 
exit
curl -g \
		 -X POST \
		 -H "Content-Type: application/json" \
		 -d "{"query":"query$1" }" \
			https://api.thegraph.com/subgraphs/name/uniswap/uniswap-v2
