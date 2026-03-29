# default-params

Demonstrates the full set of **named parameter types** using the `WithStringP`, `WithIntP`, etc. family, alongside `krait.AsJson()` and `krait.Current().Params.List()`.

## What it shows

- `WithStringP`, `WithIntP`, `WithBoolP`, `WithFloat64P`, `WithDurationP`,
  `WithStringSliceP`, `WithStringToStringP` — named parameters with a short flag.
- `krait.AsJson()` — pretty-prints all named-parameter values as JSON.
- `krait.Current().Params.List()` — lists all registered parameter names on the current
  command.
- A `USE_COBRA=true` environment variable switches the same binary to a plain Cobra
  implementation (via `examples/config-params/cmd`) for comparison.

## Running

```
cd examples

# Krait-style (default)
go run ./default-params --string hello --int 7 --bool --float64 3.14 --duration 2s

# With a config file
go run ./default-params --config default-params/config.yaml

# Cobra comparison
USE_COBRA=true go run ./default-params --string hello
```

## Key API

```go
krait.App("app", ...).
    WithStringP("appConfig.string", "...", "string", "s", "APP_STRING", "defaultStr").
    WithIntP("appConfig.int", "...", "int", "i", "APP_INT").
    WithBoolP("appConfig.bool", "...", "bool", "b", "APP_BOOL").
    WithFloat64P("appConfig.float64", "...", "float64", "f", "APP_FLOAT64").
    WithDurationP("appConfig.duration", "...", "duration", "d", "APP_DURATION").
    WithStringSliceP("appConfig.stringSlice", "...", "slice", "l", "APP_SLICE").
    WithStringToStringP("appConfig.stringToString", "...", "map", "m", "APP_MAP").
    WithRun(func(args []string) error {
        fmt.Println(krait.AsJson())
        return nil
    })
```
