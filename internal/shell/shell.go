package shell

import (
	"os"
	"runtime"
)

func IsLegacyWindowsCMD() bool {
	if runtime.GOOS != "windows" {
		return false
	}

	isCMD := os.Getenv("PROMPT") != ""
	isUnixLike := os.Getenv("TERM") != ""

	return isCMD && !isUnixLike
}
