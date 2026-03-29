package main

import (
	"fmt"
	"os"

	krait "github.com/aqua777/krait"
)

var demoParams = krait.NewConfigParams().
	With("demo.message", "message", "m", "DEMO_MESSAGE", "Sample value persisted by WriteConfig helpers", "default-msg", nil)

func newConfigFileSubcommand(name, short, long string, run func([]string) error) *krait.Command {
	return krait.New(name, short, long).
		WithConfig("", "config", "c", "APP_CONFIG").
		WithParams(demoParams).
		WithRun(run)
}

func newOutPathSubcommand(name, short, long string, run func([]string) error) *krait.Command {
	return krait.New(name, short, long).
		WithParams(demoParams).
		WithStringP("app.out", "Output file path", "out", "o", "CONFIG_WRITE_OUT").
		WithRun(run)
}

func runWrite([]string) error {
	if err := krait.WriteConfig(); err != nil {
		return err
	}
	fmt.Println("WriteConfig: wrote current settings to the file from --config")
	return nil
}

func runSafeWrite([]string) error {
	path := krait.Current().ConfigFile
	if path == "" {
		return fmt.Errorf("pass --config (file must already exist so krait can load it)")
	}
	// Viper's SafeWriteConfig() ignores SetConfigFile and only writes to
	// configPaths[0]/configName.ext; WithConfig/--config uses SetConfigFile only, so
	// SafeWriteConfig() fails with "missing configuration for 'configPath'". Use
	// SafeWriteConfigAs for the explicit loaded path (same semantics as SafeWrite).
	if err := krait.SafeWriteConfigAs(path); err != nil {
		return err
	}
	fmt.Println("SafeWriteConfigAs: wrote new file at the loaded path")
	return nil
}

func requireOutPath() (string, error) {
	out := krait.GetString("app.out")
	if out == "" {
		return "", fmt.Errorf("required flag %q not set", "out")
	}
	return out, nil
}

func runWriteAs([]string) error {
	out, err := requireOutPath()
	if err != nil {
		return err
	}
	if err := krait.WriteConfigAs(out); err != nil {
		return err
	}
	fmt.Printf("WriteConfigAs: wrote to %s\n", out)
	return nil
}

func runSafeWriteAs([]string) error {
	out, err := requireOutPath()
	if err != nil {
		return err
	}
	if err := krait.SafeWriteConfigAs(out); err != nil {
		return err
	}
	fmt.Printf("SafeWriteConfigAs: wrote new file at %s\n", out)
	return nil
}

func main() {
	write := newConfigFileSubcommand("write", "Overwrite config at loaded path",
		"Load via --config (create the file first, e.g. touch tmp.yaml). Calls krait.WriteConfig.", runWrite)

	safeWrite := newConfigFileSubcommand("safe-write", "Safe write to loaded path",
		"Load via --config. Uses krait.SafeWriteConfigAs(loaded path): errors if the file already exists.", runSafeWrite)

	writeAs := newOutPathSubcommand("write-as", "Write to explicit path",
		"Calls krait.WriteConfigAs with --out (create or overwrite).", runWriteAs)

	safeWriteAs := newOutPathSubcommand("safe-write-as", "Write to path only if missing",
		"Calls krait.SafeWriteConfigAs; errors if --out already exists.", runSafeWriteAs)

	err := krait.App("app", "Config write demo", "Demonstrates WriteConfig, SafeWriteConfigAs (loaded path), WriteConfigAs, and SafeWriteConfigAs.").
		WithCommand(write).
		WithCommand(safeWrite).
		WithCommand(writeAs).
		WithCommand(safeWriteAs).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
