# config-search

Demonstrates alternative config loading strategies: **`WithConfigName`** / **`WithConfigPath`** discovery and **`WithConfigReader`**.

## What it shows

- `cmd.WithConfigName(name)` + `cmd.WithConfigPath(dir)` — discovers a config file by
  name and directory without requiring the caller to pass `--config`. Viper searches the
  given paths for a file matching the name (any supported extension).
- `cmd.WithConfigReader(r, format)` — loads config from an `io.Reader` (e.g. an
  in-memory string). No file on disk is needed. Hot-reload (`WithConfigWatcher`) has no
  effect when a reader is used.

## Subcommands

| Subcommand | Config source |
|------------|--------------|
| `search` | Discovers `config.yaml` in the current directory via `WithConfigName`/`WithConfigPath` |
| `reader` | Reads hardcoded YAML from an `strings.NewReader` via `WithConfigReader` |

## Files

| File | Purpose |
|------|---------|
| `main.go` | App with `search` and `reader` subcommands |
| `config.yaml` | File found by the `search` subcommand (must be in the working directory) |

## Running

```
cd examples/config-search

# Discovery: reads config.yaml from the current directory
go run . search

# Reader: loads inline YAML, no file needed
go run . reader
```

## Key API

```go
cmd.WithConfigName("config")
cmd.WithConfigPath(".")

cmd.WithConfigReader(strings.NewReader(yamlString), "yaml")
```
