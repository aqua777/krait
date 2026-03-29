package main

import (
	"encoding/json"
	"fmt"
	"os"

	krait "github.com/aqua777/krait"

	"github.com/aqua777/krait/examples/config-params/cmd"
)

func asJson(v any) string {
	json, _ := json.MarshalIndent(v, "", "  ")
	return string(json)
}

var kraitApp = krait.App("app", "Qobra app short", "Qobra app long").
	// WithUsageOnError().
	// WithConfig("", "config", "c", "APP_CONFIG").
	WithStringP("appConfig.string", "string param description", "string", "s", "APP_STRING"). //, "defaultStr").
	WithIntP("appConfig.int", "int param description", "int", "i", "APP_INT").
	WithBoolP("appConfig.bool", "bool param description", "bool", "b", "APP_BOOL").
	WithFloat64P("appConfig.float64", "float64 param description", "float64", "f", "APP_FLOAT64").
	WithDurationP("appConfig.duration", "duration param description", "duration", "d", "APP_DURATION").
	WithStringSliceP("appConfig.stringSlice", "string slice param description", "slice", "l", "APP_SLICE").
	WithStringToStringP("appConfig.stringToString", "string to string param description", "map", "m", "APP_MAP").
	WithRun(func(args []string) error {
		fmt.Println("All Params:", krait.Current().Params.List())
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
