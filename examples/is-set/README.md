# is-set

Demonstrates **`WithString`** (no short flag) and **`krait.IsSet`**.

## What it shows

- `WithString(name, desc, flag, envVar, default...)` — registers a named string
  parameter without a short flag alias.
- `krait.IsSet(key)` — returns `true` when the key is registered on the current command
  and has a value from any source (CLI, env var, config file, or default). A key that
  is not registered at all returns `false` regardless of whether it has a default.

## Running

```
cd examples

# str1–str4 are all registered; IsSet returns true for each (default is a value too)
go run ./is-set

# Override one via env var
APP_STR1=custom go run ./is-set
```

## Key API

```go
krait.App("app", ...).
    WithString("appConfig.str1", "...", "str1", "APP_STR1").
    WithString("appConfig.str2", "...", "str2", "APP_STR2").
    WithRun(func(args []string) error {
        fmt.Println(krait.IsSet("appConfig.str1")) // true — registered + has default
        fmt.Println(krait.IsSet("appConfig.other")) // false — not registered
        return nil
    })
```
