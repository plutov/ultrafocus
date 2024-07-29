package hosts

import (
	"reflect"
	"testing"
)

func TestExtractDomainsFromData(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		wantStatus  FocusStatus
		wantDomains []string
		wantErr     bool
	}{
		{
			name:        "no domains",
			data:        "",
			wantDomains: []string{},
			wantStatus:  FocusStatusOff,
			wantErr:     false,
		},
		{
			name:        "no ultrafocus domains",
			data:        "127.0.0.1 example.com",
			wantDomains: []string{},
			wantStatus:  FocusStatusOff,
			wantErr:     false,
		},
		{
			name: "one unique domain and status on",
			data: `#ultrafocus:start
			#ultrafocus:on
			127.0.0.1 example.com
			127.0.0.1 example.com
			#ultrafocus:end`,
			wantDomains: []string{"example.com"},
			wantStatus:  FocusStatusOn,
			wantErr:     false,
		},
		{
			name: "no domains and status off",
			data: `#ultrafocus:start
			#ultrafocus:end`,
			wantDomains: []string{},
			wantStatus:  FocusStatusOff,
			wantErr:     false,
		},
		{
			name: "commented domains",
			data: `#ultrafocus:start
			# 127.0.0.1 example.com
			#ultrafocus:end`,
			wantDomains: []string{"example.com"},
			wantStatus:  FocusStatusOff,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDomains, gotStatus, err := extractDomainsFromData(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractDomainsFromData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotDomains, tt.wantDomains) {
				t.Errorf("extractDomainsFromData() gotDomains = %v, want %v", gotDomains, tt.wantDomains)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("extractDomainsFromData() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}

func TestUpdateHostsData(t *testing.T) {
	tests := []struct {
		name         string
		originalData string
		status       FocusStatus
		wantNewData  string
		domains      []string
		wantErr      bool
	}{
		{
			name:         "no domains, empty original data",
			originalData: "",
			domains:      []string{},
			status:       FocusStatusOff,
			wantNewData:  "#ultrafocus:start\n#ultrafocus:off\n#ultrafocus:end\n",
		},
		{
			name:         "one domain, status: on, some original data",
			originalData: "127.0.0.1 example.com",
			domains:      []string{"facebook.com"},
			status:       FocusStatusOn,
			wantNewData:  "127.0.0.1 example.com\n#ultrafocus:start\n#ultrafocus:on\n127.0.0.1 facebook.com\n#ultrafocus:end\n",
		},
		{
			name:         "one domain, status: off, original data before and after",
			originalData: "127.0.0.1 example.com\n#ultrafocus:start\n#ultrafocus:on\n127.0.0.1 facebook.com\n#ultrafocus:end\n127.0.0.1 example.com\n",
			domains:      []string{"facebook.com"},
			status:       FocusStatusOn,
			wantNewData:  "127.0.0.1 example.com\n127.0.0.1 example.com\n#ultrafocus:start\n#ultrafocus:on\n127.0.0.1 facebook.com\n#ultrafocus:end\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNewData, err := updateHostsData(tt.originalData, tt.domains, tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("updateHostsData() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gotNewData != tt.wantNewData {
				t.Errorf("updateHostsData() gotNewData = %v, want %v", gotNewData, tt.wantNewData)
			}
		})
	}
}
