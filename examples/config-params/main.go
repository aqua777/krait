package main

import (
	"fmt"
	"os"
	"time"

	krait "github.com/aqua777/krait"

	"github.com/aqua777/krait/examples/config-params/cmd"
)

var kraitApp = krait.App("app", "Qobra app short", "Qobra app long").
	// WithUsageOnError().
	WithConfig("", "config", "c", "APP_CONFIG").
	WithStringP("appConfig.string", "string param description", "string", "s", "APP_STRING", "defaultStr").
	WithIntP("appConfig.int", "int param description", "int", "i", "APP_INT", 123).
	WithBoolP("appConfig.bool", "bool param description", "bool", "b", "APP_BOOL", false).
	WithFloat64P("appConfig.float64", "float64 param description", "float64", "f", "APP_FLOAT64", 123.456).
	WithDurationP("appConfig.duration", "duration param description", "duration", "d", "APP_DURATION", time.Second*5).
	WithStringSliceP("appConfig.stringSlice", "string slice param description", "slice", "l", "APP_SLICE", []string{"default1", "default2"}).
	WithStringToStringP("appConfig.stringToString", "string to string param description", "map", "m", "APP_MAP", map[string]string{"key1": "value1", "key2": "value2"}).
	WithCommand(krait.New("version", "version short", "version long").
		WithRun(func(args []string) error {
			fmt.Println("Executing: version Qobra style")
			return nil
		})).
	WithCommand(krait.New("config", "config short", "config long").
		WithCommand(krait.New("show", "show short", "show long").
			WithRun(func(args []string) error {
				fmt.Println("Executing: config show Qobra style")
				return nil
			})).
		WithCommand(krait.New("set", "set short", "set long").
			WithRun(func(args []string) error {
				fmt.Println("Executing: config set Qobra style")
				return nil
			})).
		WithCommand(krait.New("reset", "reset short", "reset long").
			WithRun(func(args []string) error {
				fmt.Println("Executing: config reset Qobra style")
				return nil
			}))).
	WithRun(func(args []string) error {
		fmt.Println("Executing: app Qobra style")
		fmt.Println("All settings:", krait.AsJson())
		return nil
	})

func main() {
	if os.Getenv("USE_COBRA") == "true" {
		cmd.Execute()
	} else {
		kraitApp.Execute()
	}
}
