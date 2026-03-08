package main

import "fmt"

// Version information
var (
	// Version is the current version of the MCP server
	// It will be set automatically during build using ldflags
	Version = "dev"

	// GitCommit is the git commit hash
	// It will be set automatically during build using ldflags
	GitCommit = "unknown"

	// BuildDate is the date when the binary was built
	// It will be set automatically during build using ldflags
	BuildDate = "unknown"
)

// GetVersion returns the full version string
func GetVersion() string {
	if GitCommit != "unknown" {
		return fmt.Sprintf("%s (%s)", Version, GitCommit[:7])
	}
	return Version
}

// GetFullVersionInfo returns complete version information
func GetFullVersionInfo() string {
	return fmt.Sprintf("%s, commit: %s, built at: %s", Version, GitCommit, BuildDate)
}
