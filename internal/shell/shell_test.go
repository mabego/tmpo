package shell

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLegacyWindowsCMD(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Run("returns false on non-Windows platforms", func(t *testing.T) {
			result := IsLegacyWindowsCMD()
			assert.False(t, result)
		})
		return
	}

	tests := []struct {
		name     string
		prompt   string
		term     string
		expected bool
	}{
		{
			name:     "legacy CMD with PROMPT set and no TERM",
			prompt:   "$P$G",
			term:     "",
			expected: true,
		},
		{
			name:     "not CMD because PROMPT is not set",
			prompt:   "",
			term:     "",
			expected: false,
		},
		{
			name:     "Unix-like shell overrides PROMPT",
			prompt:   "$P$G",
			term:     "xterm-256color",
			expected: false,
		},
		{
			name:     "neither PROMPT nor TERM set",
			prompt:   "",
			term:     "xterm-256color",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("PROMPT", tt.prompt)
			t.Setenv("TERM", tt.term)

			result := IsLegacyWindowsCMD()
			assert.Equal(t, tt.expected, result)
		})
	}
}
