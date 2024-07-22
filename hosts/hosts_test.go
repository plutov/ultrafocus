package hosts

import (
	"reflect"
	"testing"
)

func TestExtractDomains(t *testing.T) {
	tests := []struct {
		name        string
		data        string
		wantDomains []string
		wantStatus  FocusStatus
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
			h := &Manager{}
			gotDomains, gotStatus, err := h.ExtractDomains(tt.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.ExtractDomains() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(gotDomains, tt.wantDomains) {
				t.Errorf("Manager.ExtractDomains() gotDomains = %v, want %v", gotDomains, tt.wantDomains)
			}
			if gotStatus != tt.wantStatus {
				t.Errorf("Manager.ExtractDomains() gotStatus = %v, want %v", gotStatus, tt.wantStatus)
			}
		})
	}
}
