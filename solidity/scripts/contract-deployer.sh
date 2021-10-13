#!/bin/bash

set -a # automatically export all variables
source .env
set +a

npx ts-node \
contract-deployer.ts \
--cosmos-node="$COSMOS_NODE" \
--eth-node="$RINKEBY_URL" \
--eth-privkey="$PRIVATE_KEY" \
--contract=Gravity.json \
--test-mode=false