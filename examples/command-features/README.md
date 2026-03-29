# command-features

Demonstrates command-level features: persistent flags, command groups, aliases, version, deprecation, `WithSilenceErrors`, and `WithDisableFlagParsing`.

## What it shows

| Feature | Method | Behaviour |
|---------|--------|-----------|
| Persistent flags | `WithPersistentBoolP` | Flag inherited by all subcommands |
| Command groups | `WithGroup` / `InGroup` | Subcommands grouped under a label in `--help` |
| Aliases | `WithAliases` | Extra names that invoke the same subcommand |
| Version | `WithVersion` | Adds `--version` flag to the root command |
| Deprecated command | `WithDeprecated` | Command still works but prints a deprecation notice to stderr |
| Deprecated flag | `WithFlagDeprecated` | Flag still works but prints a deprecation notice |
| Silence errors | `WithSilenceErrors` | Cobra does not print the error returned by `Run` |
| Disable flag parsing | `WithDisableFlagParsing` | All tokens (including flag-like strings) passed as positional args |

## Subcommands

| Subcommand | Feature(s) demonstrated |
|------------|------------------------|
| `serve` | Command group (`mgmt`) |
| `migrate` | Command group, aliases (`mv`, `m`), deprecated `--legacy` flag |
| `legacy` | Deprecated command |
| `exec` | `WithDisableFlagParsing` — pass-through args |
| `fail` | Returns an error; `WithSilenceErrors` suppresses the auto-printed message |

## Running

```
cd examples

# Version flag on the root
go run ./command-features --version

# Persistent --verbose inherited by subcommands
go run ./command-features serve --verbose
go run ./command-features migrate --verbose

# Alias
go run ./command-features mv --verbose

# Deprecated command (still runs, prints warning to stderr)
go run ./command-features legacy

# Pass-through: all tokens forwarded verbatim
go run ./command-features exec --any-flag value

# Error silenced: no "Error: ..." line printed
go run ./command-features fail
```

## Key API

```go
root.WithVersion("1.2.3")
root.WithGroup("mgmt", "Management Commands")
root.WithPersistentBoolP("app.verbose", "...", "verbose", "v", "APP_VERBOSE", false)
root.WithSilenceErrors()

sub.InGroup("mgmt")
sub.WithAliases("mv", "m")
sub.WithDeprecated("use 'migrate' instead")
sub.WithFlagDeprecated("legacy", "default migration path is preferred")
sub.WithDisableFlagParsing()
```
