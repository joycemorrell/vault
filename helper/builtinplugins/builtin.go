package builtinplugins

import (
	"github.com/hashicorp/vault/helper/consts"
	"github.com/hashicorp/vault/plugins/database/cassandra"
	"github.com/hashicorp/vault/plugins/database/hana"
	"github.com/hashicorp/vault/plugins/database/mongodb"
	"github.com/hashicorp/vault/plugins/database/mssql"
	"github.com/hashicorp/vault/plugins/database/mysql"
	"github.com/hashicorp/vault/plugins/database/postgresql"
	"github.com/hashicorp/vault/plugins/helper/database/credsutil"
)

// TODO move to registry
var databasePlugins = map[string]BuiltinFactory{
	// These four databasePlugins all use the same mysql implementation but with
	// different username settings passed by the constructor.
	"mysql-database-plugin":        mysql.New(mysql.MetadataLen, mysql.MetadataLen, mysql.UsernameLen),
	"mysql-aurora-database-plugin": mysql.New(credsutil.NoneLength, mysql.LegacyMetadataLen, mysql.LegacyUsernameLen),
	"mysql-rds-database-plugin":    mysql.New(credsutil.NoneLength, mysql.LegacyMetadataLen, mysql.LegacyUsernameLen),
	"mysql-legacy-database-plugin": mysql.New(credsutil.NoneLength, mysql.LegacyMetadataLen, mysql.LegacyUsernameLen),

	"postgresql-database-plugin": postgresql.New,
	"mssql-database-plugin":      mssql.New,
	"cassandra-database-plugin":  cassandra.New,
	"mongodb-database-plugin":    mongodb.New,
	"hana-database-plugin":       hana.New,
}

// BuiltinFactory is the func signature that should be returned by
// the plugin's New() func.
type BuiltinFactory func() (interface{}, error)

func toBuiltinFactory(ifc interface{}) BuiltinFactory {
	return func() (interface{}, error) {
		return ifc, nil
	}
}

// Get returns the BuiltinFactory func for a particular backend plugin
// from the databasePlugins map.
func Get(name string, pluginType consts.PluginType) (BuiltinFactory, bool) {
	switch pluginType {
	case consts.PluginTypeCredential:
		f, ok := credentialBackends[name]
		return toBuiltinFactory(f), ok
	case consts.PluginTypeSecrets:
		f, ok := logicalBackends[name]
		return toBuiltinFactory(f), ok
	case consts.PluginTypeDatabase:
		f, ok := databasePlugins[name]
		return f, ok
	default:
		return nil, false
	}
}

// TODO should this now include more keys? or be renamed to being a db plugin?
// Keys returns the list of plugin names that are considered builtin databasePlugins.
func Keys() []string {
	keys := make([]string, len(databasePlugins))

	i := 0
	for k := range databasePlugins {
		keys[i] = k
		i++
	}

	return keys
}
