//go:build windows
// +build windows

package hosts

import (
	"os"
)

var hostsPath = os.Getenv("SystemRoot") + `\System32\drivers\etc\hosts`
