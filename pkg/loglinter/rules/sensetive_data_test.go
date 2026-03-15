package rules

import "testing"

func TestCheckSensitiveData(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"safe", "user logged in", false},
		{"contains password", "password=123", true},
		{"contains token", "token: abc", true},
		{"contains secret", "my secret data", true},
		{"contains key", "api_key=xyz", true},
		{"case insensitive", "PASSWORD", true},
		{"substring safe", "passphrase", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, errMsg := checkSensitiveData(tt.msg)
			if tt.wantErr && ok {
				t.Errorf("checkSensitiveData() expected error, got ok")
			}
			if !tt.wantErr && !ok {
				t.Errorf("checkSensitiveData() unexpected error: %s", errMsg)
			}
		})
	}
}
