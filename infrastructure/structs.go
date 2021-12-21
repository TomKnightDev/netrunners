package infrastructure

type Cpu struct {
	CpuId    int
	CpuType  CpuType
	CpuSpeed float64
}

type CpuType struct {
	CpuTypeId    int
	CpuTypeName  string
	CpuTypeSpeed float64
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
