# config-vars

Demonstrates **var-bound parameters** (`WithStringVarP`, `WithIntVarP`, etc.) and **`WithDebug`** / **`IsDebug`**.

## What it shows

- `WithStringVarP(&ptr, ...)`, `WithIntVarP`, `WithBoolVarP`, `WithFloat64VarP`,
  `WithDurationVarP`, `WithStringSliceVarP`, `WithStringToStringVarP` — writes the
  resolved value directly into a typed pointer before `Run` is called. No getter call
  needed inside `Run`.
- `WithDebug(flag, shortFlag, envVar)` — registers a boolean debug flag. Readable inside
  the run lifecycle via `krait.IsDebug()`.
- `cmd.AllSettingsAsJson()` — all named-param (Viper-backed) values as JSON.
- `cmd.ArgsSettingsAsJson()` — all var-bound param values as JSON.

## Running

```
cd examples

# Basic run — all var-bound params at their defaults
go run ./config-vars

# Override a few values
go run ./config-vars --string hello --int 99 --bool

# With a config file
go run ./config-vars --config config-vars/config.yaml

# Debug mode prints both Viper and argsViper snapshots
go run ./config-vars --debug
```

## Key API

```go
var host string
var port int

krait.App("app", ...).
    WithStringVarP(&host, "Host", "host", "h", "APP_HOST", "localhost").
    WithIntVarP(&port, "Port", "port", "p", "APP_PORT", 8080).
    WithDebug("debug", "", "APP_DEBUG").
    WithRun(func(args []string) error {
        // host and port are already populated here
        if krait.IsDebug() {
            fmt.Println(krait.Current().AllSettingsAsJson())
        }
        return nil
    })
```
