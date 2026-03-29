# config-introspect

Demonstrates config introspection: **`krait.IsSet`**, **`krait.Unmarshal`**, and **`krait.Sub`**.

## What it shows

- `krait.IsSet(key)` — returns `true` when a key is registered on the current command
  and has a value from any source (CLI, env, config file, or default). A key that is not
  registered at all returns `false`, even if it appears in a config file.
- `krait.Unmarshal(&cfg)` — decodes all named parameters into a typed Go struct using
  `mapstructure` tags.
- `krait.Sub(prefix)` — returns a flat `map[string]interface{}` with the prefix
  stripped, useful for extracting a config subtree.

## Subcommands

| Subcommand | What it demonstrates |
|------------|---------------------|
| `is-set` | `IsSet(host)=true` when `APP_HOST` is set; `IsSet(port)=false` because `port` is not registered on this command |
| `unmarshal` | Decodes all named params into `IntrospectConfig` via `krait.Unmarshal` |
| `sub` | Prints `krait.Sub("database")` as a sorted flat map |

## Files

| File | Purpose |
|------|---------|
| `main.go` | App and three introspection subcommands |
| `config.yaml` | Config file used by `unmarshal` and `sub` |

## Running

```
cd examples

# IsSet: host is registered, port is not
APP_HOST=myhost go run ./config-introspect is-set

# Unmarshal into a struct
go run ./config-introspect unmarshal --config config-introspect/config.yaml

# Sub-tree as a flat map
go run ./config-introspect sub --config config-introspect/config.yaml
```

## Key API

```go
krait.IsSet("host")          // bool
krait.Unmarshal(&cfg)        // decode named params into a struct
krait.Sub("database")        // map[string]interface{} with "database." prefix stripped
```
