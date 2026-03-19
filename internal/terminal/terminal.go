package terminal

import (
	"fmt"

	"github.com/mgreau/zen/internal/ghostty"
	"github.com/mgreau/zen/internal/iterm"
)

// Terminal represents a terminal emulator that can open tabs/windows.
type Terminal interface {
	Name() string
	OpenTab(workDir, command string) error
	OpenTabWithResume(workDir, sessionID, claudeBin, model string) error
	OpenTabWithClaude(workDir, initialPrompt, claudeBin, model string) error
}

// NewTerminal creates a new terminal instance based on the terminal type.
// An empty string defaults to headless mode.
func NewTerminal(terminalType string) (Terminal, error) {
	switch terminalType {
	case "iterm":
		return &ITermTerminal{}, nil
	case "ghostty":
		return &GhosttyTerminal{}, nil
	case "headless", "":
		return &HeadlessTerminal{}, nil
	default:
		return nil, fmt.Errorf("unsupported terminal type: %s (supported: iterm, ghostty, headless)", terminalType)
	}
}

// HeadlessTerminal prints commands instead of opening terminal tabs.
type HeadlessTerminal struct{}

func (t *HeadlessTerminal) Name() string { return "headless" }

func (t *HeadlessTerminal) OpenTab(workDir, command string) error {
	fmt.Printf("\n  cd %s && %s\n", workDir, command)
	return nil
}

func (t *HeadlessTerminal) OpenTabWithResume(workDir, sessionID, claudeBin, model string) error {
	cmd := claudeBin
	if model != "" {
		cmd += fmt.Sprintf(" --model %s", model)
	}
	cmd += fmt.Sprintf(" --resume %s", sessionID)
	fmt.Printf("\n  cd %s && %s\n", workDir, cmd)
	return nil
}

func (t *HeadlessTerminal) OpenTabWithClaude(workDir, initialPrompt, claudeBin, model string) error {
	cmd := claudeBin
	if model != "" {
		cmd += fmt.Sprintf(" --model %s", model)
	}
	cmd += fmt.Sprintf(" %q", initialPrompt)
	fmt.Printf("\n  cd %s && %s\n", workDir, cmd)
	return nil
}

// ITermTerminal wraps the iTerm functions.
type ITermTerminal struct{}

func (t *ITermTerminal) Name() string {
	return "iTerm2"
}

func (t *ITermTerminal) OpenTab(workDir, command string) error {
	return iterm.OpenTab(workDir, command)
}

func (t *ITermTerminal) OpenTabWithResume(workDir, sessionID, claudeBin, model string) error {
	return iterm.OpenTabWithResume(workDir, sessionID, claudeBin, model)
}

func (t *ITermTerminal) OpenTabWithClaude(workDir, initialPrompt, claudeBin, model string) error {
	return iterm.OpenTabWithClaude(workDir, initialPrompt, claudeBin, model)
}

// GhosttyTerminal wraps the Ghostty functions.
type GhosttyTerminal struct{}

func (t *GhosttyTerminal) Name() string {
	return "Ghostty"
}

func (t *GhosttyTerminal) OpenTab(workDir, command string) error {
	return ghostty.OpenTab(workDir, command)
}

func (t *GhosttyTerminal) OpenTabWithResume(workDir, sessionID, claudeBin, model string) error {
	return ghostty.OpenTabWithResume(workDir, sessionID, claudeBin, model)
}

func (t *GhosttyTerminal) OpenTabWithClaude(workDir, initialPrompt, claudeBin, model string) error {
	return ghostty.OpenTabWithClaude(workDir, initialPrompt, claudeBin, model)
}