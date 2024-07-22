package hosts

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strings"
)

type Manager struct {
	hostsFile *os.File
	Status    FocusStatus
	Domains   []string
}

type FocusStatus string

const (
	ipAddress                    = "127.0.0.1"
	FocusStatusOn    FocusStatus = "on"
	FocusStatusOff   FocusStatus = "off"
	CommentStart                 = "#ultrafocus:start"
	CommentEnd                   = "#ultrafocus:end"
	CommentStatusOn              = "#ultrafocus:on"
	CommentStatusOff             = "#ultrafocus:off"
)

func (h *Manager) Init() error {
	h.Status = FocusStatusOff
	f, err := os.OpenFile(hostsPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var extractErr error
	h.Domains, h.Status, extractErr = h.ExtractDomains(string(data))
	if extractErr != nil {
		return extractErr
	}

	return nil
}

func (h *Manager) ExtractDomains(data string) ([]string, FocusStatus, error) {
	domains := []string{}
	inComment := false
	status := FocusStatusOff

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := scanner.Text()
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine == CommentStart {
			inComment = true
			continue
		} else if trimmedLine == CommentEnd {
			inComment = false
			break
		}

		if inComment {
			if trimmedLine == CommentStatusOn {
				status = FocusStatusOn
				continue
			}
			if trimmedLine == CommentStatusOff {
				status = FocusStatusOff
				continue
			}

			uncommentedLine := strings.Replace(trimmedLine, "#", "", 1)
			fields := strings.Fields(uncommentedLine)
			if len(fields) > 1 {
				if !slices.Contains(domains, fields[1]) {
					domains = append(domains, fields[1])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return domains, status, err
	}

	return domains, status, nil
}
