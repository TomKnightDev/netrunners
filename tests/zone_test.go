package tests

import (
	"testing"

	infra "github.com/tomknightdev/netrunners/infrastructure"
)

func TestGetSystemsInfo(t *testing.T) {
	// Create zone
	z := infra.NewZone(0, "Zone0")
	z.Systems = append(z.Systems, *infra.NewSystem(0, "Sys0", *infra.NewZoneAddress(z)))

	if len(z.GetSystemsInfo()) != 1 {
		t.Error("Zone count was not 1")
	}
}

func TestZoneGetSystemInfo(t *testing.T) {
	z := infra.NewZone(0, "Zone0")
	z.Systems = append(z.Systems, *infra.NewSystem(0, "Sys0", *infra.NewZoneAddress(z)))

	if z.GetSystemInfo("Sys0") == "System not found" {
		t.Error("System not found")
	}
}
