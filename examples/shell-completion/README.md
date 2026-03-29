# shell-completion

Demonstrates shell completion helpers: **`WithValidArgs`**, **`WithValidArgsFunction`**, and **`WithFlagCompletion`**.

## What it shows

- `cmd.WithValidArgs(args...)` — static list of valid positional argument completions.
- `cmd.WithValidArgsFunction(fn)` — dynamic positional completions via a callback;
  takes precedence over `WithValidArgs` when both are set.
- `cmd.WithFlagCompletion(flag, fn)` — dynamic completion for a named flag's value.
- Cobra generates the `completion` subcommand automatically; these methods control
  what candidates are offered.

## Subcommands

| Subcommand | Completion type |
|------------|----------------|
| `serve` | Static args: `start`, `stop`, `restart` via `WithValidArgs` |
| `deploy` | Dynamic args: `staging`, `production`, `canary` via `WithValidArgsFunction` |
| `run` | Dynamic `--env` flag values: `dev`, `staging`, `prod` via `WithFlagCompletion` |

## Running

```
cd examples
go run ./shell-completion serve start
go run ./shell-completion deploy staging
go run ./shell-completion run --env prod
```

To inspect completions (requires Cobra's `completion` subcommand):

```
go run ./shell-completion completion bash
```

## Key API

```go
cmd.WithValidArgs("start", "stop", "restart")

cmd.WithValidArgsFunction(func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
    return []string{"staging", "production"}, cobra.ShellCompDirectiveNoFileComp
})

cmd.WithFlagCompletion("env", func(args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
    return []string{"dev", "staging", "prod"}, cobra.ShellCompDirectiveNoFileComp
})
```
