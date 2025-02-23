<h1 align="center">
ultrafocus
</h1>

<h1 align="center">
Focus is key. Unlock your potential.
</h1>

A dead simple CLI tool to block distracting websites and boost productivity. Customize your blacklist, focus on tasks, and reclaim your time.

<p align="center" width="100%">
<img src="https://github.com/plutov/ultrafocus/blob/main/screenshots/domains.png" hspace="10" height="200px">
<img src="https://github.com/plutov/ultrafocus/blob/main/screenshots/status.png" hspace="10" height="200px">
</p>

## Installation

Download the latest binary from the [releases page](https://github.com/plutov/ultrafocus/releases/latest) or use `go install`:

```bash
go install github.com/plutov/ultrafocus@latest
```

## Usage

ultrafocus needs `sudo` to modify `/etc/hosts` file. It won't affect your existing configuration, the changes made by ultrafocus are separated by `#ultrafocus:start` and `#ultrafocus:end` comments.

```bash
sudo ultrafocus
```

ultrafocus also runs a server on `127.0.0.1:80` where all the requests are redirected to. This page shows a random motivational message.

<img src="https://github.com/plutov/ultrafocus/blob/main/screenshots/result.png" height="125px">

## Supported platforms

- macOS
- Linux
- Windows

## Default blacklist

You can find the default blacklist in [hosts/domains.go](https://github.com/plutov/ultrafocus/blob/main/hosts/domains.go).

## Run tests

```bash
go test -v -race ./...
```
