package system

import (
	"os"
	"strings"
)

func CPU() string {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return "unknown"
	}

	for _, line := range strings.Split(string(data), "\n") {
		if strings.HasPrefix(line, "model name") {
			parts := strings.SplitN(line, ":", 2)

			if len(parts) == 2 {
				return strings.TrimSpace(parts[1])
			}
		}
	}

	return "unknown"
}
