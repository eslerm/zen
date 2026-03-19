package notify

import (
	"fmt"
	"log"
	"os/exec"
)

// Send sends a desktop notification using the platform-appropriate tool.
// Falls back gracefully: osascript (macOS) -> notify-send (Linux) -> log only.
func Send(title, message, subtitle string) error {
	// Try macOS osascript
	if path, err := exec.LookPath("osascript"); err == nil {
		script := fmt.Sprintf(`display notification %q with title %q`, message, title)
		if subtitle != "" {
			script = fmt.Sprintf(`display notification %q with title %q subtitle %q`, message, title, subtitle)
		}
		return exec.Command(path, "-e", script).Run()
	}

	// Try Linux notify-send
	if path, err := exec.LookPath("notify-send"); err == nil {
		body := message
		if subtitle != "" {
			body = fmt.Sprintf("%s\n%s", subtitle, message)
		}
		return exec.Command(path, title, body).Run()
	}

	// No notification tool available — log and continue
	log.Printf("[notify] %s: %s", title, message)
	return nil
}

// PRReview notifies about a new PR review request.
func PRReview(prNumber int, prTitle, author, repo string) error {
	return Send(
		"New PR Review Request",
		fmt.Sprintf("PR #%d: %s", prNumber, prTitle),
		fmt.Sprintf("by %s in %s", author, repo),
	)
}

// WorktreeReady notifies that a worktree is ready.
func WorktreeReady(prNumber int, worktreePath string) error {
	return Send(
		"Worktree Ready",
		fmt.Sprintf("PR #%d worktree created", prNumber),
		worktreePath,
	)
}

// PRMerged notifies about a PR merge.
func PRMerged(prNumber int, prTitle string) error {
	return Send(
		"PR Merged",
		fmt.Sprintf("PR #%d: %s", prNumber, prTitle),
		"Worktree can be cleaned up",
	)
}

// StaleWorktrees notifies about stale worktrees found.
func StaleWorktrees(count int) error {
	return Send(
		"Stale Worktrees Found",
		fmt.Sprintf("%d worktrees can be cleaned up", count),
		"Run: zen cleanup",
	)
}
