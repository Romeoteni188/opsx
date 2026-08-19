package system

import "runtime"

func OS() string {
	return runtime.GOOS
}
