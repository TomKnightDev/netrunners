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

	if len(vals) != 100 {
		t.Error("Incorrect system list returned")
	}
}

func TestProcessCommandLSHW(t *testing.T) {
	Init()

	quit, vals := ProcessCommand("lshw Sys079")

	if quit == true {
		t.Error("Process command return incorrect value")
	}

	if vals[0] == "System not found" {
		t.Error("System not found")
	}
}
