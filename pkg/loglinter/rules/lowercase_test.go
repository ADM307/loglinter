package rules

import "testing"

func TestCheckLowercase(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"empty", "", false},
		{"lowercase", "hello world", false},
		{"uppercase", "Hello world", true},
		{"digit first", "1st message", false},
		{"symbol first", "!important", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, errMsg := checkLowercase(tt.msg)
			if tt.wantErr && ok {
				t.Errorf("checkLowercase() expected error, got ok")
			}
			if !tt.wantErr && !ok {
				t.Errorf("checkLowercase() unexpected error: %s", errMsg)
			}
		})
	}
}
