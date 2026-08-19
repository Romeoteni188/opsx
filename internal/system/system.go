package system

import "runtime"

type InfoResult struct {
	OS           string `json:"os"`
	Architecture string `json:"architecture"`
	Kernel       string `json:"kernel"`
	CPU          string `json:"cpu"`
	Memory       string `json:"memory"`
	Disk         string `json:"disk"`
}

func Info() InfoResult {
	return InfoResult{
		OS:           OS(),
		Architecture: runtime.GOARCH,
		Kernel:       Kernel(),
		CPU:          CPU(),
		Memory:       Memory(),
		Disk:         Disk(),
	}
}
