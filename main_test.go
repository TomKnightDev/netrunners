package main

import (
	"testing"
)

func TestProcessCommandLS(t *testing.T) {
	Init()
	quit, vals := ProcessCommand("ls")

	if quit == true {
		t.Error("Process command return incorrect value")
	}

	if len(vals) != 3 {
		t.Error("Incorrect system list returned")
	}
}
