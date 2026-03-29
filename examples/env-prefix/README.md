# env-prefix

Demonstrates **`WithEnvPrefix`** and getter-only value types that have no corresponding CLI flag.

## What it shows

- `cmd.WithEnvPrefix("MYAPP")` enables Viper's `AutomaticEnv` with a prefix so that
  `MYAPP_HOST` resolves to the `host` key — no per-flag env var name needed.
- Getter-only types: `GetIntSlice`, `GetTime`, `GetSizeInBytes`, and
  `GetStringMapStringSlice` read values that can only come from a config file or
  environment variable (no matching `With*` flag type exists for them).

## Subcommands

| Subcommand | What it does |
|------------|-------------|
| `prefix` | Reads `host` and `port` resolved via `MYAPP_HOST` / `MYAPP_PORT` |
| `types` | Reads `ids` (int slice), `created_at` (time), `max_size` (bytes), and `headers` (string map of string slices) from a config file |

## Files

| File | Purpose |
|------|---------|
| `main.go` | App definition with `prefix` and `types` subcommands |
| `config.yaml` | Config file used by the `types` subcommand |

## Running

```
cd examples

# env prefix
MYAPP_HOST=myserver go run ./env-prefix prefix

# getter-only types
go run ./env-prefix types --config env-prefix/config.yaml
```

## Key API

```go
cmd.WithEnvPrefix("MYAPP")

// In Run:
krait.GetIntSlice("ids")
krait.GetTime("created_at")
krait.GetSizeInBytes("max_size")
krait.GetStringMapStringSlice("headers")
```
