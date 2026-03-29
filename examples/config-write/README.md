# config-write

Demonstrates config file write helpers: **`WriteConfig`**, **`WriteConfigAs`**, and **`SafeWriteConfigAs`**.

## What it shows

- `krait.WriteConfig()` — overwrites the file that was loaded via `--config`.
- `krait.WriteConfigAs(path)` — writes current settings to an explicit path (creates or
  overwrites).
- `krait.SafeWriteConfigAs(path)` — writes to an explicit path only if the file does not
  already exist; returns an error otherwise.
- Why `SafeWriteConfig()` (no path) is not used directly: Viper's `SafeWriteConfig`
  requires a `configPath`/`configName` pair, which krait does not set when `WithConfig`
  is used. `SafeWriteConfigAs` with the loaded path is the correct equivalent.

## Subcommands

| Subcommand | What it does |
|------------|-------------|
| `write` | Overwrites the file passed via `--config` |
| `safe-write` | Writes to the loaded path; errors if the file already exists |
| `write-as` | Writes to `--out` (create or overwrite) |
| `safe-write-as` | Writes to `--out` only if it does not already exist |

## Files

| File | Purpose |
|------|---------|
| `main.go` | Four write subcommands |
| `tmp.yaml` | Scratch file used by `write` and `safe-write` examples |

## Running

```
cd examples

# Overwrite the loaded config file
go run ./config-write write --config config-write/tmp.yaml --message "hello"

# Write to a new path
go run ./config-write write-as --out /tmp/out.yaml --message "hello"

# Safe write (errors if /tmp/out.yaml already exists)
go run ./config-write safe-write-as --out /tmp/new.yaml --message "hello"
```

## Key API

```go
krait.WriteConfig()
krait.WriteConfigAs("/path/to/output.yaml")
krait.SafeWriteConfigAs("/path/to/new.yaml")
```
