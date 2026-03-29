package main

import (
	"fmt"
	"os"
	"strings"

	krait "github.com/aqua777/krait"
	"github.com/spf13/cobra"
)

// Demonstrates WithValidArgs, WithValidArgsFunction, and WithFlagCompletion for shell completion.

func filterPrefix(candidates []string, toComplete string) []string {
	var out []string
	for _, c := range candidates {
		if strings.HasPrefix(c, toComplete) {
			out = append(out, c)
		}
	}
	return out
}

func completeDeployArgs(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	candidates := []string{"staging", "production", "canary"}
	return filterPrefix(candidates, toComplete), cobra.ShellCompDirectiveNoFileComp
}

func completeRunEnvFlag(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	candidates := []string{"dev", "staging", "prod"}
	return filterPrefix(candidates, toComplete), cobra.ShellCompDirectiveNoFileComp
}

func runServe(args []string) error {
	fmt.Println("serve:", strings.Join(args, " "))
	return nil
}

func runDeploy(args []string) error {
	fmt.Println("deploy:", strings.Join(args, " "))
	return nil
}

func runRun([]string) error {
	fmt.Printf("run: env=%s\n", krait.GetString("run.env"))
	return nil
}

func main() {
	serve := krait.New("serve", "Control the server", "Static positional completions via WithValidArgs.").
		WithValidArgs("start", "stop", "restart").
		WithRun(runServe)

	deploy := krait.New("deploy", "Deploy an environment", "Dynamic positional completions via WithValidArgsFunction.").
		WithValidArgsFunction(completeDeployArgs).
		WithRun(runDeploy)

	runCmd := krait.New("run", "Run a job", "Flag value completions via WithFlagCompletion on --env.").
		WithStringP("run.env", "Environment name", "env", "e", "RUN_ENV", "dev").
		WithFlagCompletion("env", completeRunEnvFlag).
		WithRun(runRun)

	err := krait.App("app", "Shell completion demo", "Demonstrates WithValidArgs, WithValidArgsFunction, and WithFlagCompletion.").
		WithCommand(serve).
		WithCommand(deploy).
		WithCommand(runCmd).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
