package infrastructure

import "fmt"

type Zone struct {
	ZoneId   int
	ZoneName string
	Systems  []System
}

func NewZone(zoneId int, zoneName string) *Zone {
	z := new(Zone)
	z.ZoneId = zoneId
	z.Systems = []System{}
	z.ZoneName = zoneName
	return z
}

func (z *Zone) GetSystemsInfo() []string {
	var systemsInfo = []string{}

	for _, sys := range z.Systems {
		systemsInfo = append(systemsInfo,
			fmt.Sprintf("%s - %v", sys.SystemName, sys.ZoneAddress.GetZoneAddress()))
	}

	return systemsInfo
}

func (z *Zone) GetSystemInfo(sysName string) string {
	for _, sys := range z.Systems {
		if sys.SystemName == sysName {
			return fmt.Sprintf("%s - %v\n%s", sys.SystemName, sys.ZoneAddress.GetZoneAddress(), sys.GetSystemInfo())
		}
	}

	return fmt.Sprint("System not found")
}
