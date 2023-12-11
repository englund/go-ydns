# YDNS Updater

This project is a YDNS updater written in Go.

## Build

```sh
$ go build -o ydns-updater
```

## Usage

**Retrieve current IP**

```sh
$ ydns-updater ip
```

**Update an YDNS record with current IP**

```sh
$ ydns-updater update -H example.com [-H example2.com]
```
