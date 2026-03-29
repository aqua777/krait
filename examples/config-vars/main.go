package main

import (
	"fmt"
	"time"

	krait "github.com/aqua777/krait"
)

var (
	stringVar         string
	intVar            int
	boolVar           bool
	float64Var        float64
	durationVar       time.Duration
	stringSliceVar    []string
	stringToStringVar map[string]string

	mappings = map[string]any{
		"stringVar":         &stringVar,
		"intVar":            &intVar,
		"boolVar":           &boolVar,
		"float64Var":        &float64Var,
		"durationVar":       &durationVar,
		"stringSliceVar":    &stringSliceVar,
		"stringToStringVar": &stringToStringVar,
	}
)

var kraitApp = krait.App("app", "Qobra app short", "Qobra app long").
	// WithUsageOnError().
	WithConfig("", "config", "c", "APP_CONFIG").
	WithDebug("debug", "", "APP_DEBUG").
	WithStringVarP(&stringVar, "string param description", "string", "s", "APP_STRING", "defaultStr").
	WithIntVarP(&intVar, "int param description", "int", "i", "APP_INT", 123).
	WithBoolVarP(&boolVar, "bool param description", "bool", "b", "APP_BOOL", false).
	WithFloat64VarP(&float64Var, "float64 param description", "float64", "f", "APP_FLOAT64", 123.456).
	WithDurationVarP(&durationVar, "duration param description", "duration", "d", "APP_DURATION", time.Second*5).
	WithStringSliceVarP(&stringSliceVar, "string slice param description", "slice", "l", "APP_SLICE", []string{"default1", "default2"}).
	WithStringToStringVarP(&stringToStringVar, "string to string param description", "map", "m", "APP_MAP", map[string]string{"key1": "value1", "key2": "value2"}).
	WithRun(func(args []string) error {
		fmt.Println("Executing: app Qobra style", "---")
		if krait.IsDebug() {
			fmt.Println("All settings from viper:", krait.Current().AllSettingsAsJson(), "---")
			fmt.Println("All settings from argsViper:", krait.Current().ArgsSettingsAsJson(), "---")
		}

		for key, value := range mappings {
			fmt.Printf("%s: ", key)
			switch v := value.(type) {
			case *string:
				fmt.Printf("%s\n", *v)
			case *int:
				fmt.Printf("%d\n", *v)
			case *bool:
				fmt.Printf("%t\n", *v)
			case *float64:
				fmt.Printf("%f\n", *v)
			case *time.Duration:
				fmt.Printf("%s\n", *v)
			case *[]string:
				fmt.Printf("%v\n", *v)
			case *map[string]string:
				fmt.Printf("%v\n", *v)
			}
		}
		return nil
	})

func main() {
	kraitApp.Execute()
}
