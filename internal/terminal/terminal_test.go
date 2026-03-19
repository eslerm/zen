package terminal

import (
	"testing"
)

func TestNewTerminal(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		wantName    string
		wantErr     bool
	}{
		{"iterm explicit", "iterm", "iTerm2", false},
		{"ghostty", "ghostty", "Ghostty", false},
		{"headless explicit", "headless", "headless", false},
		{"empty defaults to headless", "", "headless", false},
		{"invalid terminal", "invalid", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			term, err := NewTerminal(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTerminal(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if err != nil {
				return
			}
			if got := term.Name(); got != tt.wantName {
				t.Errorf("NewTerminal(%q).Name() = %q, want %q", tt.input, got, tt.wantName)
			}
		})
	}
}

func TestITermTerminalName(t *testing.T) {
	term := &ITermTerminal{}
	if got := term.Name(); got != "iTerm2" {
		t.Errorf("ITermTerminal.Name() = %q, want %q", got, "iTerm2")
	}
}

func TestGhosttyTerminalName(t *testing.T) {
	term := &GhosttyTerminal{}
	if got := term.Name(); got != "Ghostty" {
		t.Errorf("GhosttyTerminal.Name() = %q, want %q", got, "Ghostty")
	}
}

func TestHeadlessTerminalName(t *testing.T) {
	term := &HeadlessTerminal{}
	if got := term.Name(); got != "headless" {
		t.Errorf("HeadlessTerminal.Name() = %q, want %q", got, "headless")
	}
}

func TestHeadlessTerminalOpenTab(t *testing.T) {
	term := &HeadlessTerminal{}
	if err := term.OpenTab("/tmp/test", "echo hello"); err != nil {
		t.Errorf("HeadlessTerminal.OpenTab() error = %v", err)
	}
}

func TestHeadlessTerminalOpenTabWithResume(t *testing.T) {
	term := &HeadlessTerminal{}
	if err := term.OpenTabWithResume("/tmp/test", "session-123", "claude", "sonnet"); err != nil {
		t.Errorf("HeadlessTerminal.OpenTabWithResume() error = %v", err)
	}
}

func TestHeadlessTerminalOpenTabWithClaude(t *testing.T) {
	term := &HeadlessTerminal{}
	if err := term.OpenTabWithClaude("/tmp/test", "/review-pr", "claude", ""); err != nil {
		t.Errorf("HeadlessTerminal.OpenTabWithClaude() error = %v", err)
	}
}
