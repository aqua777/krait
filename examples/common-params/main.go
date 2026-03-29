package main

import (
	"fmt"

	krait "github.com/aqua777/krait"
)

func main() {
	commonParams := krait.NewConfigParams().
		With("common.string", "string", "s", "COMMON_STRING", "This is common string", "default-value", nil).
		With("common.int", "int", "i", "COMMON_INT", "This is common int", "1", nil)

	versionCmd := krait.New("version", "Version command", "Version command description").
		WithParams(commonParams).
		WithRun(func(args []string) error {
			fmt.Println("Executing: version Krait style", krait.AsJson())
			return nil
		})

	configCmd := krait.New("config", "Config command", "Config command").
		WithParams(commonParams).
		WithRun(func(args []string) error {
			fmt.Println("Executing: config Krait style", krait.AsJson())
			return nil
		})

	krait.App("app", "Krait app short", "Krait app long").
		WithCommand(versionCmd).
		WithCommand(configCmd).
		Execute()
}
