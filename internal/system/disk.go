package system

import (
	"os/exec"
	"strings"
)

func Disk() string {
	output, err := exec.Command("df", "-h", "/").Output()
	if err != nil {
		return "unknown"
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	if len(lines) < 2 {
		return "unknown"
	}

	fields := strings.Fields(lines[1])

	if len(fields) < 5 {
		return "unknown"
	}

	return fields[4] + " used"
}
