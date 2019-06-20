// +build !windows

package configdir

import (
	"github.com/kardianos/osext"
	"os"
	"path/filepath"
)

func configDir(appName string) string {
	configPath, ok := os.LookupEnv("HOME")
	if !ok {
		return filepath.Join(Executable(), ".config", appName)
	}
	return filepath.Join(configPath, ".config", appName)
}

func Executable() string {
	executablePath, err := osext.Executable()
	if err != nil {
		executablePath, _ = os.Executable()
	}

	linkedExecutablePath, err := filepath.EvalSymlinks(executablePath)
	if err != nil {
		return executablePath
	}
	return linkedExecutablePath
}
