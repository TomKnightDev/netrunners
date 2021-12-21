package infrastructure

type System struct {
	SystemId   int
	SystemName string
	Processor  []Cpu
	Memory     []Ram
	Storage    []Ssd
	Power      Psu
}

func NewSystem() *System {
	sys := new(System)
	sys.SystemId = 0

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
