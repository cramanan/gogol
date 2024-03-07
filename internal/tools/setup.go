package tools

import "runtime"

func OS() string {
	return runtime.GOOS
}
