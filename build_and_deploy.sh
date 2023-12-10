#!/usr/bin/env bash

source .env

env GOOS=linux GOARCH=arm64 go build -o ydns-updater .
scp ydns-updater $SCP_DEPLOY_DEST