package main

import (
	"fmt"
	"os"
	"strings"

	krait "github.com/aqua777/krait"
)

// Demonstrates persistent flags, command groups, aliases, version, deprecated
// commands and flags, silence-errors, and pass-through args.

func runServe([]string) error {
	fmt.Printf("serve: verbose=%v\n", krait.GetBool("app.verbose"))
	return nil
}

func runMigrate([]string) error {
	fmt.Printf("migrate: verbose=%v\n", krait.GetBool("app.verbose"))
	return nil
}

func runLegacy([]string) error {
	fmt.Println("legacy: command still runs after deprecation notice on stderr")
	return nil
}

func runExec(args []string) error {
	fmt.Println(strings.Join(args, " "))
	return nil
}

func runFail([]string) error {
	return fmt.Errorf("intentional failure")
}

func main() {
	serve := krait.New("serve", "Start the server", "Runs the HTTP server").
		InGroup("mgmt").
		WithRun(runServe)

	migrate := krait.New("migrate", "Run database migrations", "Applies pending migrations").
		InGroup("mgmt").
		WithAliases("mv", "m").
		WithBoolP("migrate.legacy", "Legacy migration path", "legacy", "l", "MIGRATE_LEGACY", false).
		WithFlagDeprecated("legacy", "default migration path is preferred").
		WithRun(runMigrate)

	legacy := krait.New("legacy", "Legacy entrypoint", "Deprecated compatibility command").
		WithDeprecated("use 'migrate' instead").
		WithRun(runLegacy)

	exec := krait.New("exec", "Pass-through execution", "Disables flag parsing; forwards tokens verbatim").
		WithDisableFlagParsing().
		WithRun(runExec)

	fail := krait.New("fail", "Return an error", "Demonstrates WithSilenceErrors on the root command").
		WithRun(runFail)

	err := krait.App("app", "Command feature demo", "Demonstrates persistent flags, groups, aliases, version, deprecation, silence-errors, and pass-through args").
		WithVersion("1.2.3").
		WithGroup("mgmt", "Management Commands").
		WithPersistentBoolP("app.verbose", "Verbose logging", "verbose", "v", "APP_VERBOSE", false).
		WithSilenceErrors().
		WithCommand(serve).
		WithCommand(migrate).
		WithCommand(legacy).
		WithCommand(exec).
		WithCommand(fail).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
