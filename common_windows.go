package configdir

import (
	"github.com/kardianos/osext"
	"os"
	"path/filepath"
)

func configDir(appName string) string {
	configPath, ok := os.LookupEnv("APPDATA")
	if !ok {
		return filepath.Join(Executable(), appName)
	}
	return filepath.Join(configPath, appName)
}

func Executable() string {
	e, err := osext.Executable()
	if err != nil {
		e, _ = os.Executable()
	}

	return e
}
