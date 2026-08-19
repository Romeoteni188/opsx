package doctor

import (
	"runtime"

	"golang.org/x/sys/unix"
)

func CheckSystem() Category {
	checks := []CheckResult{
		checkOperatingSystem(),
		checkArchitecture(),
		checkKernel(),
	}

	return Category{
		Name:   "System",
		Checks: checks,
	}
}

func checkOperatingSystem() CheckResult {
	if runtime.GOOS == "linux" {
		return CheckResult{
			Name:    "Operating System",
			Passed:  true,
			Message: "Linux",
		}
	}

	return CheckResult{
		Name:    "Operating System",
		Passed:  false,
		Message: runtime.GOOS,
	}
}

func checkArchitecture() CheckResult {
	return CheckResult{
		Name:    "Architecture",
		Passed:  true,
		Message: runtime.GOARCH,
	}
}

func checkKernel() CheckResult {
	var uts unix.Utsname

	if err := unix.Uname(&uts); err != nil {
		return CheckResult{
			Name:    "Kernel",
			Passed:  false,
			Message: err.Error(),
		}
	}

	return CheckResult{
		Name:    "Kernel",
		Passed:  true,
		Message: unix.ByteSliceToString(uts.Release[:]),
	}
}
