#!/bin/bash
# the script run inside the container for all-up-test.sh
NODES=$1

# Prepare the contracts for later deployment
pushd /peggy/solidity/
HUSKY_SKIP_INSTALL=1 npm install
npm run typechain

bash /peggy/tests/container-scripts/setup-validators.sh $NODES

bash /peggy/tests/container-scripts/run-testnet.sh $NODES &

bash /peggy/tests/container-scripts/integration-tests.sh $NODES