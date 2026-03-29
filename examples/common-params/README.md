# common-params

Demonstrates shared parameter sets: **`NewConfigParams`** and **`WithParams`**.

## What it shows

- `krait.NewConfigParams()` creates a reusable parameter set.
- `.With(name, flag, shortFlag, envVar, desc, default, ...)` adds a parameter to the set.
- `cmd.WithParams(params)` applies the entire set to a command — registers all flags,
  binds all env vars, and sets all defaults in one call.
- The same `ConfigParams` instance can be applied to multiple subcommands so shared
  flags are declared exactly once.

## Running

```
cd examples
go run ./common-params version --string hello --int 42
go run ./common-params config --string world
```

Both subcommands accept `--string` and `--int` because both call `WithParams(commonParams)`.

## Key API

```go
commonParams := krait.NewConfigParams().
    With("common.string", "string", "s", "COMMON_STRING", "A shared string", "default-value", nil).
    With("common.int",    "int",    "i", "COMMON_INT",    "A shared int",    "1",             nil)

krait.New("version", ...).WithParams(commonParams).WithRun(...)
krait.New("config",  ...).WithParams(commonParams).WithRun(...)
```
