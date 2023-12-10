#!/usr/bin/env bash

source .env

env GOOS=linux GOARCH=arm64 go build .
scp ydns $SCP_DEPLOY_DEST