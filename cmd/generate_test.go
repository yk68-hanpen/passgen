package cmd

import (
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	tests := []struct {
		name         string
		length       int
		useSpecial   bool
		useNumbers   bool
		useUppercase bool
		useLowercase bool
		wantLength   int
		wantContains map[string]bool
	}{
		{
			name:         "Default options",
			length:       12,
			useSpecial:   true,
			useNumbers:   true,
			useUppercase: true,
			useLowercase: true,
			wantLength:   12,
			wantContains: map[string]bool{
				"lowercase": true,
				"uppercase": true,
				"numbers":   true,
				"special":   true,
			},
		},
		{
			name:         "Only lowercase",
			length:       8,
			useSpecial:   false,
			useNumbers:   false,
			useUppercase: false,
			useLowercase: true,
			wantLength:   8,
			wantContains: map[string]bool{
				"lowercase": true,
				"uppercase": false,
				"numbers":   false,
				"special":   false,
			},
		},
		{
			name:         "No options selected defaults to lowercase",
			length:       10,
			useSpecial:   false,
			useNumbers:   false,
			useUppercase: false,
			useLowercase: false,
			wantLength:   10,
			wantContains: map[string]bool{
				"lowercase": true,
				"uppercase": false,
				"numbers":   false,
				"special":   false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePassword(tt.length, tt.useSpecial, tt.useNumbers, tt.useUppercase, tt.useLowercase)

			// Check length
			if len(got) != tt.wantLength {
				t.Errorf("generatePassword() length = %v, want %v", len(got), tt.wantLength)
			}

			// Check character types
			hasLowercase := strings.ContainsAny(got, "abcdefghijklmnopqrstuvwxyz")
			hasUppercase := strings.ContainsAny(got, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
			hasNumbers := strings.ContainsAny(got, "0123456789")
			hasSpecial := strings.ContainsAny(got, "!@#$%^&*()-_=+[]{}|;:,.<>?/")

			if hasLowercase != tt.wantContains["lowercase"] {
				t.Errorf("generatePassword() contains lowercase = %v, want %v", hasLowercase, tt.wantContains["lowercase"])
			}
			if hasUppercase != tt.wantContains["uppercase"] {
				t.Errorf("generatePassword() contains uppercase = %v, want %v", hasUppercase, tt.wantContains["uppercase"])
			}
			if hasNumbers != tt.wantContains["numbers"] {
				t.Errorf("generatePassword() contains numbers = %v, want %v", hasNumbers, tt.wantContains["numbers"])
			}
			if hasSpecial != tt.wantContains["special"] {
				t.Errorf("generatePassword() contains special = %v, want %v", hasSpecial, tt.wantContains["special"])
			}
		})
	}
}
