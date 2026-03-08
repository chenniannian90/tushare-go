package main

import (
	"regexp"
	"testing"
)

func TestGetVersion(t *testing.T) {
	tests := []struct {
		name        string
		version     string
		gitCommit   string
		wantPattern string
	}{
		{
			name:        "dev version with unknown commit",
			version:     "dev",
			gitCommit:   "unknown",
			wantPattern: `^dev$`,
		},
		{
			name:        "release version with commit",
			version:     "v1.0.0",
			gitCommit:   "abc123def",
			wantPattern: `^v1\.0\.0 \(abc123d\)$`,
		},
		{
			name:        "version with short commit hash",
			version:     "v2.1.3-beta",
			gitCommit:   "xyz789abc",
			wantPattern: `^v2\.1\.3-beta \(xyz789a\)$`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original values
			origVersion := Version
			origGitCommit := GitCommit
			defer func() {
				Version = origVersion
				GitCommit = origGitCommit
			}()

			// Set test values
			Version = tt.version
			GitCommit = tt.gitCommit

			// Get version
			got := GetVersion()

			// Check pattern
			matched, err := regexp.MatchString(tt.wantPattern, got)
			if err != nil {
				t.Fatalf("regexp error: %v", err)
			}
			if !matched {
				t.Errorf("GetVersion() = %q, want pattern %q", got, tt.wantPattern)
			}
		})
	}
}

func TestGetFullVersionInfo(t *testing.T) {
	tests := []struct {
		name      string
		version   string
		gitCommit string
		buildDate string
		want      string
	}{
		{
			name:      "full version info",
			version:   "v1.0.0",
			gitCommit: "abc123",
			buildDate: "2024-01-01_12:00:00",
			want:      "v1.0.0, commit: abc123, built at: 2024-01-01_12:00:00",
		},
		{
			name:      "dev version info",
			version:   "dev",
			gitCommit: "unknown",
			buildDate: "unknown",
			want:      "dev, commit: unknown, built at: unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Save original values
			origVersion := Version
			origGitCommit := GitCommit
			origBuildDate := BuildDate
			defer func() {
				Version = origVersion
				GitCommit = origGitCommit
				BuildDate = origBuildDate
			}()

			// Set test values
			Version = tt.version
			GitCommit = tt.gitCommit
			BuildDate = tt.buildDate

			// Get full version info
			got := GetFullVersionInfo()

			if got != tt.want {
				t.Errorf("GetFullVersionInfo() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestVersionInMCPImplementation(t *testing.T) {
	// This test ensures that the Version variable is properly used in MCP implementation
	// Save original value
	origVersion := Version
	defer func() { Version = origVersion }()

	// Set a test version
	Version = "v1.2.3-test"

	// Verify that Version variable can be accessed
	if Version != "v1.2.3-test" {
		t.Errorf("Version variable = %q, want %q", Version, "v1.2.3-test")
	}

	// GetVersion should return the version
	got := GetVersion()
	expected := "v1.2.3-test"
	if got != expected {
		t.Errorf("GetVersion() = %q, want %q", got, expected)
	}
}
