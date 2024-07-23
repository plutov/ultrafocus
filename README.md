# ultrafocus

A dead simple CLI tool to block distracting websites and boost productivity. Customize your blacklist, focus on tasks, and reclaim your time.

<p align="center" width="100%">
    <img src="https://github.com/plutov/ultrafocus/blob/main/screenshots/domains.png" hspace="10" height="50px">
    <img src="https://github.com/plutov/ultrafocus/blob/main/screenshots/status.png" hspace="10" height="50px">
</p>

## Installation

Download the latest binary from the [releases page](https://github.com/plutov/ultrafocus/releases) or use `go install`:

```bash
go install github.com/plutov/ultrafocus
```

## Usage

ultrafocus needs `sudo` to modify `/etc/hosts` file. It won't affect your existing configuration, the changes made by ultrafocus are separated by `#ultrafocus:start/end` comments.

```bash
sudo ultrafocus
```

## Supported platforms

- macOS
- Linux
- Windows

## Run tests

```bash
go test -v -race ./...
```
