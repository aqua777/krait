package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	krait "github.com/aqua777/krait"
)

// prefixParams uses empty environmentVarName so AutomaticEnv + WithEnvPrefix maps
// MYAPP_HOST→host and MYAPP_PORT→port without repeating the prefix per flag.
var prefixParams = krait.NewConfigParams().
	With("host", "host", "", "", "HTTP listen host", "localhost", nil).
	With("port", "port", "", "", "HTTP listen port", 9000, nil)

func runPrefix([]string) error {
	fmt.Printf("host=%s\n", krait.GetString("host"))
	return nil
}

func runTypes([]string) error {
	fmt.Printf("int_slice=%v\n", krait.GetIntSlice("ids"))
	fmt.Printf("time=%s\n", krait.GetTime("created_at").Format(time.RFC3339))
	fmt.Printf("size_bytes=%d\n", krait.GetSizeInBytes("max_size"))
	m := krait.GetStringMapStringSlice("headers")
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("headers[%s]=%v\n", k, m[k])
	}
	return nil
}

func main() {
	// WithEnvPrefix must be on the command that owns the Viper used at Run (subcommands
	// each have their own Viper; the root app does not inherit env prefix to children).
	prefix := krait.New("prefix", "Env prefix", "Resolves MYAPP_HOST and MYAPP_PORT via WithEnvPrefix; omit per-flag env prefix in declarations.").
		WithEnvPrefix("MYAPP").
		WithParams(prefixParams).
		WithRun(runPrefix)

	types := krait.New("types", "Getter-only value types", "Reads GetIntSlice, GetTime, GetSizeInBytes, and GetStringMapStringSlice from the config file (no matching With* for these types).").
		WithEnvPrefix("MYAPP").
		WithConfig("", "config", "c", "MYAPP_CONFIG").
		WithRun(runTypes)

	err := krait.App("app", "Env prefix and getter-only types", "Demonstrates WithEnvPrefix and package-level getters for config-only value shapes.").
		WithCommand(prefix).
		WithCommand(types).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
