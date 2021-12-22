package infrastructure

import "fmt"

type ZoneAddress struct {
	address []int
}

func NewZoneAddress(zone *Zone) *ZoneAddress {
	za := new(ZoneAddress)
	za.address = []int{zone.ZoneId, 0, len(zone.Systems)}
	return za
}

func (za *ZoneAddress) GetZoneAddress() string {
	return fmt.Sprintf("%03d.%03d.%03d", za.address[0], za.address[1], za.address[2])
}
