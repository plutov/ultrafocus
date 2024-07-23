package hosts

import (
	"bufio"
	"io"
	"os"
	"slices"
	"strings"
)

type FocusStatus string

const (
	defaultIpAddress             = "127.0.0.1"
	FocusStatusOn    FocusStatus = "on"
	FocusStatusOff   FocusStatus = "off"
	CommentStart                 = "#ultrafocus:start"
	CommentEnd                   = "#ultrafocus:end"
	CommentStatusOn              = "#ultrafocus:on"
	CommentStatusOff             = "#ultrafocus:off"
)

func ExtractDomainsFromHostsFile() ([]string, FocusStatus, error) {
	domains := []string{}
	status := FocusStatusOff
	f, err := os.OpenFile(hostsPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return domains, status, err
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return domains, status, err
	}

	var extractErr error
	domains, status, extractErr = extractDomainsFromData(string(data))
	if extractErr != nil {
		return domains, status, extractErr
	}

	return domains, status, nil
}

func CleanDomainsList(domains []string) []string {
	uniqueDomains := []string{}
	for _, domain := range domains {
		domain = strings.TrimSpace(strings.ToLower(domain))
		if domain != "" && !slices.Contains(uniqueDomains, domain) {
			uniqueDomains = append(uniqueDomains, domain)
		}
	}

	return uniqueDomains
}

func WriteDomainsToHostsFile(domains []string, status FocusStatus) error {
	f, err := os.OpenFile(hostsPath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	newData, updateErr := updateHostsData(string(data), domains, status)
	if updateErr != nil {
		return updateErr
	}

	if _, err := f.Seek(0, 0); err != nil {
		return err
	}
	if _, err := f.WriteString(newData); err != nil {
		return err
	}
	if err := f.Truncate(int64(len(newData))); err != nil {
		return err
	}

	return nil
}

func extractDomainsFromData(data string) ([]string, FocusStatus, error) {
	domains := []string{}
	inComment := false
	status := FocusStatusOff

	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if line == CommentStart {
			inComment = true
			continue
		} else if line == CommentEnd {
			inComment = false
			break
		}

		if inComment {
			if line == CommentStatusOn {
				status = FocusStatusOn
				continue
			}
			if line == CommentStatusOff {
				status = FocusStatusOff
				continue
			}

			uncommentedLine := strings.Replace(line, "#", "", 1)
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

func updateHostsData(originalData string, domains []string, status FocusStatus) (string, error) {
	inComment := false
	newData := ""
	scanner := bufio.NewScanner(strings.NewReader(originalData))
	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if line == CommentStart {
			inComment = true
			continue
		} else if line == CommentEnd {
			inComment = false
			continue
		}

		if !inComment {
			newData += line + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		return originalData, err
	}

	newData += CommentStart + "\n"
	if status == FocusStatusOn {
		newData += CommentStatusOn + "\n"
	} else {
		newData += CommentStatusOff + "\n"
	}
	for _, d := range domains {
		if status == FocusStatusOn {
			newData += defaultIpAddress + " " + d + "\n"
		} else {
			newData += "#" + defaultIpAddress + " " + d + "\n"
		}
	}
	newData += CommentEnd + "\n"

	return newData, nil
}
