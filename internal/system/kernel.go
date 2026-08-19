package system

import "golang.org/x/sys/unix"

func Kernel() string {
	var uts unix.Utsname

	if err := unix.Uname(&uts); err != nil {
		return "unknown"
	}

	return unix.ByteSliceToString(uts.Release[:])
}
