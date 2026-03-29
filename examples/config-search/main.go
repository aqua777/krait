package main

import (
	"fmt"
	"os"
	"strings"

	krait "github.com/aqua777/krait"
)

// Hardcoded YAML for the reader subcommand (no file on disk).
const readerConfigYAML = `demo:
  title: "in-memory YAML via WithConfigReader"
  count: 7
`

var demoParams = krait.NewConfigParams().
	With("demo.title", "title", "t", "DEMO_TITLE", "Demo title", "default-title", nil).
	With("demo.count", "count", "n", "DEMO_COUNT", "Demo count", -1, nil)

func printDemoFromConfig() error {
	fmt.Printf("demo.title=%q\n", krait.GetString("demo.title"))
	fmt.Printf("demo.count=%d\n", krait.GetInt("demo.count"))
	return nil
}

func runPrintDemo([]string) error {
	return printDemoFromConfig()
}

func main() {
	search := krait.New("search", "Load config via name/path search", "Finds config.yaml in the current directory using WithConfigName and WithConfigPath (no --config).").
		WithConfigName("config").
		WithConfigPath(".").
		WithParams(demoParams).
		WithRun(runPrintDemo)

	reader := krait.New("reader", "Load config from io.Reader", "Loads YAML from an in-memory string using WithConfigReader.").
		WithConfigReader(strings.NewReader(readerConfigYAML), "yaml").
		WithParams(demoParams).
		WithRun(runPrintDemo)

	err := krait.App("app", "Config source demo", "Demonstrates WithConfigName/WithConfigPath discovery and WithConfigReader.").
		WithCommand(search).
		WithCommand(reader).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
