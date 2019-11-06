#!/usr/bin/env bash

set -x

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
echo "gkicoin binary dir:" "$DIR"
pushd "$DIR" >/dev/null

COMMIT=$(git rev-parse HEAD)
BRANCH=$(git rev-parse --abbrev-ref HEAD)
GOLDFLAGS="-X main.Commit=${COMMIT} -X main.Branch=${BRANCH}"

go run -ldflags "${GOLDFLAGS}" cmd/gkicoin/gkicoin.go \
    -enable-wallet-api=true \
	-enable-unversioned-api=true \
    -disable-csrf=false \
    -rpc-interface=false \
    -log-level=debug \
    $@

popd >/dev/null
