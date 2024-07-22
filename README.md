# ultrafocus

CLI tool to block distracting websites and boost productivity. Customize your blacklist, focus on tasks, and reclaim your time.

## Installation

```bash
go install  github.com/plutov/ultrafocus
```

## Usage

ultrafocus needs sudo to modify `/etc/hosts` file. It won't affect your existing configuration, the changes made my ultrafocus are separated by `# ultrafocus` comment.

```bash
sudo ultrafocus
```
