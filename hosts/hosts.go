package hosts

import (
	"os"
)

type Manager struct {
	hostsFile *os.File
}

func (h *Manager) Init() error {
	var err error
	h.hostsFile, err = os.OpenFile("/etc/hosts", os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	return nil
}
