package main

import (
	"fmt"
	"os"
	"sync"
	"time"

	krait "github.com/aqua777/krait"
)

// daemonParams registers the key loaded from config.yaml for the daemon command.
var daemonParams = krait.NewConfigParams().
	With("demo.message", "message", "m", "DEMO_MESSAGE", "Demo string observed across reloads", "default-msg", nil)

var (
	reloadPrintMu sync.Mutex
	lastPrinted   string
)

// onConfigReloaded runs when the watched config file changes (Viper reloads, then this fires).
//
// Synchronisation: krait invokes this from fsnotify’s background goroutine, not from Run.
// If the callback updated shared application state, you would need a mutex or channel
// hand-off with the main goroutine; this demo uses a mutex only to coalesce duplicate
// watcher callbacks for the same reloaded value.
func onConfigReloaded() {
	msg := krait.GetString("demo.message")
	reloadPrintMu.Lock()
	defer reloadPrintMu.Unlock()
	if msg == lastPrinted {
		return
	}
	lastPrinted = msg
	fmt.Printf("config reloaded: demo.message=%q\n", msg)
}

func runDaemon([]string) error {
	path := krait.Current().ConfigFile
	if path == "" {
		return fmt.Errorf("no config file path (use --config)")
	}

	original, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	fmt.Printf("initial: %q\n", krait.GetString("demo.message"))

	// Non-interactive demo: trigger the same filesystem event a manual edit would, so the
	// watcher runs without leaving the process running indefinitely.
	go func() {
		time.Sleep(80 * time.Millisecond)
		err := os.WriteFile(path, []byte(`demo:
  message: "updated on disk"
`), 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "demo write: %v\n", err)
		}
	}()

	time.Sleep(3 * time.Second)

	if err := os.WriteFile(path, original, 0644); err != nil {
		return fmt.Errorf("restore config: %w", err)
	}
	return nil
}

func main() {
	daemon := krait.New("daemon", "Watch config file for changes", "Loads config, prints the initial value, then prints again when the file changes (demo triggers a write after startup).").
		WithConfig("", "config", "c", "APP_CONFIG").
		WithParams(daemonParams).
		WithConfigWatcher(onConfigReloaded).
		WithRun(runDaemon)

	err := krait.App("app", "Config hot-reload demo", "Demonstrates WithConfigWatcher: Viper reloads the file, then the callback runs.").
		WithCommand(daemon).
		Execute()
	if err != nil {
		os.Exit(1)
	}
}
