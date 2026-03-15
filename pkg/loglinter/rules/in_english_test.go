package rules

import "testing"

func TestCheckInEnglish(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"english only", "hello world", false},
		{"with cyrillic", "привет world", true},
		{"with digits", "hello 123", false},
		{"with symbols", "hello!", false},
		{"mixed letters", "hello мир", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, errMsg := checkInEnglish(tt.msg)
			if tt.wantErr && ok {
				t.Errorf("checkInEnglish() expected error, got ok")
			}
			if !tt.wantErr && !ok {
				t.Errorf("checkInEnglish() unexpected error: %s", errMsg)
			}
		})
	}
}
