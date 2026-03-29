package main

import (
	"fmt"

	krait "github.com/aqua777/krait"
)

var kraitApp = krait.App("app", "Qobra app short", "Qobra app long").
	WithString("appConfig.str1", "string param description", "str1", "APP_STR1").
	WithString("appConfig.str2", "string param description", "str2", "APP_STR2").
	WithString("appConfig.str3", "string param description", "str3", "APP_STR3").
	WithString("appConfig.str4", "string param description", "str4", "APP_STR4").
	WithRun(func(args []string) error {
		fmt.Println("Executing: app Qobra style")
		fmt.Println("appConfig.string:", krait.AsJson())
		fmt.Println("---")
		// for _, key := range []string{"appConfig.str1", "appConfig.str2", "appConfig.str3", "appConfig.str4"} {
		// 	fmt.Println("Is set", key, "?", krait.IsSet(key))
		// }
		return nil
	})

func main() {
	kraitApp.Execute()
}
