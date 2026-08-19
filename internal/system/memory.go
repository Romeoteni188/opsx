package system

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Memory() string {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return "unknown"
	}
	defer func() {
		_ = file.Close()
	}()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) >= 2 && fields[0] == "MemTotal:" {
			kb, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return "unknown"
			}

			gb := float64(kb) / 1024 / 1024

			return strconv.FormatFloat(gb, 'f', 1, 64) + " GB"
		}
	}

	return "unknown"
}
