package consts

type PluginType int

const (
	PluginTypeCredential PluginType = iota
	PluginTypeSecrets
	PluginTypeDatabase
)
