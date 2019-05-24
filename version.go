package main

import "fmt"

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

// Version is
func Version() string {
	return fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
}
