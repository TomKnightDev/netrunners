package tests

import (
	"testing"

	infra "github.com/tomknightdev/netrunners/infrastructure"
)

func TestRam(t *testing.T) {
	var sys1 = infra.System{
		Memory: []infra.Ram{
			{
				RamId:    0,
				RamSize:  1000,
				RamSpeed: 1333,
			},
			{
				RamId:    1,
				RamSize:  1000,
				RamSpeed: 1333,
			}},
	}

	if sys1.GetTotalMemory() != 2000 {
		t.Errorf("Total memory was not %d", sys1.GetTotalMemory())
	}
}

func TestGetSystemInfo(t *testing.T) {
	var sys1 = infra.System{
		Memory: []infra.Ram{
			{
				RamId:    0,
				RamSize:  1000,
				RamSpeed: 1333,
			},
			{
				RamId:    1,
				RamSize:  1000,
				RamSpeed: 1333,
			}},
	}

	if sys1.GetSystemInfo() != "2000" {
		t.Errorf("System info is not correct")
	}
}
