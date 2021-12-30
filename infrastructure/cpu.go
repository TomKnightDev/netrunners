package infrastructure

type CpuModel struct {
	CpuModelId           int
	CpuModelName         string
	CpuModelClock        float32
	CpuModelTdp          int
	CpuModelManufacturer CpuManufacturer
}

var CpuModels = make(map[string]CpuModel)

func init() {
	CpuModels["12600K"] = CpuModel{
		CpuModelId:           0,
		CpuModelName:         "12600K",
		CpuModelClock:        3.7,
		CpuModelTdp:          125,
		CpuModelManufacturer: intel,
	}
	CpuModels["5950X"] = CpuModel{
		CpuModelId:           1,
		CpuModelName:         "5950X",
		CpuModelClock:        3.4,
		CpuModelTdp:          105,
		CpuModelManufacturer: amd,
	}
}

type CpuManufacturer int

const (
	intel CpuManufacturer = iota
	amd
)

func (cm CpuManufacturer) String() string {
	return []string{"intel", "amd"}[cm]
}
