package infrastructure

type Cpu struct {
	CpuId    int
	CpuModel CpuModel
	Temp     float32
	Age      float32
}

type Ssd struct {
	SsdId   int
	SsdSize float64
}

type Ram struct {
	RamId    int
	RamSize  int
	RamSpeed float64
}

type Psu struct {
	PsuId int
}
