package infrastructure

type System struct {
	SystemId    int
	SystemName  string
	Processor   []Cpu
	Memory      []Ram
	Storage     []Ssd
	Power       Psu
	ZoneAddress ZoneAddress
}

func NewSystem(systemId int, systemName string, zoneAddress ZoneAddress) *System {
	sys := new(System)
	sys.SystemId = systemId
	sys.SystemName = systemName
	sys.ZoneAddress = zoneAddress

	return sys
}

func (s *System) AddMemory(mem Ram) {
	s.Memory = append(s.Memory, mem)
}

func (s *System) GetTotalMemory() int {
	total := 0
	for _, mem := range s.Memory {
		total += mem.RamSize
	}

	return total
}
