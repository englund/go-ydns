# YDNS Updater

This project is a YDNS updater written in Go.

## Build

```sh
$ make build
```

## Usage

**Retrieve current IP**

```sh
$ bin/ydns-updater ip
```

**Update an YDNS record with current IP**

```sh
$ bin/ydns-updater update -H example.com [-H example2.com]
```
