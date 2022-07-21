// Package conf provides config for each plugin
package conf

type PluginConfNS string

const (
	Sensitive  PluginConfNS = "veinmind-sensitive"
	WeakPass   PluginConfNS = "veinmind-weakpass"
	FileFilter PluginConfNS = "veinmind-file-filter"
)
