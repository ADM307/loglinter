package rules

import "testing"

func TestCheckSpecialSymbols(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"normal", "hello world", false},
		{"exclamation", "hello!", true},
		{"question", "hello?", true},
		{"colon", "hello: world", true},
		{"semicolon", "hello;", true},
		{"ellipsis", "hello...", true},
		{"emoji", "hello 😀", true},
		{"multiple emoji", "😀😁", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, errMsg := checkSpecialSymbols(tt.msg)
			if tt.wantErr && ok {
				t.Errorf("checkSpecialSymbols() expected error, got ok")
			}
			if !tt.wantErr && !ok {
				t.Errorf("checkSpecialSymbols() unexpected error: %s", errMsg)
			}
		})
	}
}
