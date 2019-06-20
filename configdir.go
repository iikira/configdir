// Package configdir 获取配置文件路径
package configdir

// ConfigDir 返回配置文件路径
func ConfigDir(appName string) string {
	return configDir(appName)
}
