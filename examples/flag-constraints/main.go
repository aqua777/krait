package main

import (
	"fmt"
	"os"
	"strings"

	krait "github.com/aqua777/krait"
)

// Demonstrates required flags, required-together, one-required, and mutually-exclusive
// flag groups via krait's fluent API.

func runRequired([]string) error {
	fmt.Printf("required: ok output=%q\n", krait.GetString("required.output"))
	return nil
}

func runTogether([]string) error {
	fmt.Printf("together: ok user=%q password=<redacted>\n", krait.GetString("together.user"))
	return nil
}

func runOneOf([]string) error {
	var parts []string
	if f := krait.GetString("oneof.file"); f != "" {
		parts = append(parts, fmt.Sprintf("file=%q", f))
	}
	if u := krait.GetString("oneof.url"); u != "" {
		parts = append(parts, fmt.Sprintf("url=%q", u))
	}
	fmt.Printf("one-of: ok %s\n", strings.Join(parts, " "))
	return nil
}

func runExclusive([]string) error {
	fmt.Printf("exclusive: ok json=%v yaml=%v\n", krait.GetBool("exclusive.json"), krait.GetBool("exclusive.yaml"))
	return nil
}

func main() {
	required := krait.New("required", "Required flag", "Fails unless --output is set").
		WithStringP("required.output", "Output path", "output", "o", "FLAG_REQUIRED_OUTPUT", "").
		WithRequired("output").
		WithRun(runRequired)

	together := krait.New("together", "Flags required together", "User and password must both be set, or neither").
		WithStringP("together.user", "User name", "user", "u", "FLAG_TOGETHER_USER", "").
		WithStringP("together.password", "Password", "password", "p", "FLAG_TOGETHER_PASSWORD", "").
		WithFlagsRequiredTogether("user", "password").
		WithRun(runTogether)

	oneOf := krait.New("one-of", "At least one of two flags", "Provide --file or --url (or both)").
		WithStringP("oneof.file", "Local file path", "file", "f", "FLAG_ONEOF_FILE", "").
		WithStringP("oneof.url", "Remote URL", "url", "", "FLAG_ONEOF_URL", "").
		WithFlagsOneRequired("file", "url").
		WithRun(runOneOf)

	exclusive := krait.New("exclusive", "Mutually exclusive flags", "Use --json or --yaml, not both").
		WithBoolP("exclusive.json", "JSON output", "json", "j", "FLAG_EXCLUSIVE_JSON", false).
		WithBoolP("exclusive.yaml", "YAML output", "yaml", "y", "FLAG_EXCLUSIVE_YAML", false).
		WithFlagsMutuallyExclusive("json", "yaml").
		WithRun(runExclusive)

	err := krait.App("app", "Flag constraints demo", "Demonstrates WithRequired, WithFlagsRequiredTogether, WithFlagsOneRequired, and WithFlagsMutuallyExclusive").
		WithCommand(required).
		WithCommand(together).
		WithCommand(oneOf).
		WithCommand(exclusive).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
