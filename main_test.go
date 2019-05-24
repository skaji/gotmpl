package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestBasic(t *testing.T) {
	stdout := new(bytes.Buffer)

	exit := run(stdout, os.Stderr, []string{"me", "example/text.txt"})
	if exit != 0 {
		t.Fatal("unexpected")
	}

	whoami := os.Getenv("USER")
	if stdout.String() != fmt.Sprintf("I'm %s\n", whoami) {
		t.Fatal("unexpected")
	}
}
