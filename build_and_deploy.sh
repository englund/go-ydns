#!/usr/bin/env bash

env GOOS=linux GOARCH=arm64 go build .
scp ydns root@192.168.1.1:/usr/local/bin/ydns-updater