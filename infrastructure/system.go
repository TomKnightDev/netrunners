package infrastructure

import "fmt"

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

	// Defaults
	sys.AddMemory(Ram{RamId: 0, RamSize: 1024, RamSpeed: 1333})
	sys.AddCpu(Cpu{CpuId: 0, CpuModel: CpuModels["12600K"]})

	return sys
}

func (s *System) AddCpu(cpu Cpu) {
	s.Processor = append(s.Processor, cpu)
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

func (s *System) GetCpuInfo() string {
	info := fmt.Sprintf("- Cpu: %s", s.Processor[0].CpuModel.CpuModelName)
	info += fmt.Sprintf("\n- Manufacturer: %s", s.Processor[0].CpuModel.CpuModelManufacturer.String())
	info += fmt.Sprintf("\n- Clock: %f", s.Processor[0].CpuModel.CpuModelClock)
	return info
}

func (s *System) GetSystemInfo() string {
	info := s.GetCpuInfo()
	info += "\n" + fmt.Sprintf("- Ram: %d", s.GetTotalMemory())

	return info
}
