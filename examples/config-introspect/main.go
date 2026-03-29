package main

import (
	"fmt"
	"os"
	"sort"

	krait "github.com/aqua777/krait"
)

// IntrospectConfig matches every named parameter registered in introspectParams for Unmarshal.
type IntrospectConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	App  struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"app"`
	Database struct {
		Host    string `mapstructure:"host"`
		Port    int    `mapstructure:"port"`
		SSLMode string `mapstructure:"sslmode"`
	} `mapstructure:"database"`
}

func withSharedHost(p *krait.ConfigParams) *krait.ConfigParams {
	return p.With("host", "host", "", "APP_HOST", "HTTP listen host", "127.0.0.1", nil)
}

var introspectParams = withSharedHost(krait.NewConfigParams()).
	With("port", "port", "p", "APP_PORT", "HTTP listen port", 8080, nil).
	With("app.name", "app-name", "", "APP_NAME", "Application display name", "demo", nil).
	With("database.host", "db-host", "", "DB_HOST", "Database host", "localhost", nil).
	With("database.port", "db-port", "", "DB_PORT", "Database port", 5432, nil).
	With("database.sslmode", "db-sslmode", "", "DB_SSLMODE", "Database SSL mode", "disable", nil)

// isSet registers only host. IsSet("port") is false because port is not bound on this command—
// an unset key—not "explicit vs default" (if port were registered, default-only would still yield IsSet=true).
var isSetParams = withSharedHost(krait.NewConfigParams())

func runIsSet([]string) error {
	fmt.Printf("host: IsSet=%v\n", krait.IsSet("host"))
	fmt.Printf("port: IsSet=%v\n", krait.IsSet("port"))
	return nil
}

func runUnmarshal([]string) error {
	var cfg IntrospectConfig
	if err := krait.Unmarshal(&cfg); err != nil {
		return err
	}
	fmt.Printf("%+v\n", cfg)
	return nil
}

func runSub([]string) error {
	m := krait.Sub("database")
	if m == nil {
		return fmt.Errorf("Sub(%q) returned nil (no database subtree)", "database")
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%s: %v\n", k, m[k])
	}
	return nil
}

func main() {
	isSet := krait.New("is-set", "IsSet: env vs absent key",
		"With APP_HOST set, IsSet(host)=true. Only host is registered here—port is not—so IsSet(port)=false (unset key), not 'value left at default' (registered defaults still report IsSet=true).").
		WithParams(isSetParams).
		WithRun(runIsSet)

	unmarshal := krait.New("unmarshal", "Unmarshal into struct", "Decodes all named parameters into IntrospectConfig via krait.Unmarshal.").
		WithConfig("", "config", "c", "APP_CONFIG").
		WithParams(introspectParams).
		WithRun(runUnmarshal)

	sub := krait.New("sub", "Sub-tree as flat map", "Prints krait.Sub(\"database\") as a prefix-stripped flat map.").
		WithConfig("", "config", "c", "APP_CONFIG").
		WithParams(introspectParams).
		WithRun(runSub)

	err := krait.App("app", "Config introspection demo", "Demonstrates krait.IsSet, krait.Unmarshal, and krait.Sub.").
		WithCommand(isSet).
		WithCommand(unmarshal).
		WithCommand(sub).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
