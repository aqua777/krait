# config-hotreload

Demonstrates **`WithConfigWatcher`** — live config file reloading for long-running daemons.

## What it shows

- `cmd.WithConfigWatcher(fn)` wires Viper's `WatchConfig`/`OnConfigChange` so the callback
  fires whenever the config file changes on disk.
- The callback runs in a background goroutine (spawned by fsnotify). Shared state accessed
  from the callback must be synchronised; the example uses a mutex to suppress duplicate
  watcher events for the same value.
- `krait.GetString` inside the callback reads the freshly-reloaded value without any
  additional setup.

## Files

| File | Purpose |
|------|---------|
| `main.go` | App and daemon subcommand wired with `WithConfigWatcher` |
| `config.yaml` | Initial config file read by the daemon |

## Running

```
cd examples
go run ./config-hotreload daemon --config config-hotreload/config.yaml
```

Expected output (the demo writes a new value to the file after 80 ms, then restores it):

```
initial: "hello from config"
config reloaded: demo.message="updated on disk"
```

## Key API

```go
cmd.WithConfigWatcher(func() {
    msg := krait.GetString("demo.message")
    // react to the new value
})
```
