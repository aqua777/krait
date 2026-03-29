# config-params

A full-featured app demonstrating **config file support**, **all named parameter types**, and **nested subcommands** — with an optional Cobra comparison mode.

## What it shows

- `WithConfig(default, flag, shortFlag, envVar)` — config file path follows source
  priority: CLI flag > env var > default.
- All named parameter types with short flags: `WithStringP`, `WithIntP`, `WithBoolP`,
  `WithFloat64P`, `WithDurationP`, `WithStringSliceP`, `WithStringToStringP`.
- Nested subcommand tree: `app version`, `app config show`, `app config set`,
  `app config reset`.
- `krait.AsJson()` — prints all resolved param values as pretty JSON.
- `USE_COBRA=true` switches to the raw Cobra implementation in `cmd/root.go` for
  a side-by-side comparison of the boilerplate krait eliminates.

## Files

| File | Purpose |
|------|---------|
| `main.go` | Krait app + `USE_COBRA` switcher |
| `cmd/root.go` | Plain Cobra equivalent (no Viper) |
| `cmd/root.example` | Example config file |

## Running

```
cd examples

# Krait — all defaults
go run ./config-params

# Pass flags
go run ./config-params --string hello --int 7 --bool --duration 10s

# With config file
go run ./config-params --config config-params/cmd/root.example

# Subcommands
go run ./config-params version
go run ./config-params config show

# Cobra comparison
USE_COBRA=true go run ./config-params --string hello
```

## Key API

```go
krait.App("app", ...).
    WithConfig("", "config", "c", "APP_CONFIG").
    WithStringP("appConfig.string", "...", "string", "s", "APP_STRING", "defaultStr").
    WithIntP("appConfig.int", "...", "int", "i", "APP_INT", 123).
    // ... more param types ...
    WithCommand(krait.New("version", ...).WithRun(...)).
    WithCommand(krait.New("config", ...).
        WithCommand(krait.New("show", ...).WithRun(...))).
    WithRun(func(args []string) error {
        fmt.Println(krait.AsJson())
        return nil
    })
```
