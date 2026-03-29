# flag-constraints

Demonstrates flag constraint methods: **`WithRequired`**, **`WithFlagsRequiredTogether`**, **`WithFlagsOneRequired`**, and **`WithFlagsMutuallyExclusive`**.

## What it shows

Cobra enforces all constraints at `Execute()` time and returns an error before `Run` is
called. All four methods panic at startup if any named flag is not registered on the
command.

| Method | Behaviour |
|--------|-----------|
| `WithRequired(flag)` | Execution fails if the flag is absent |
| `WithFlagsRequiredTogether(flags...)` | All must be provided, or none |
| `WithFlagsOneRequired(flags...)` | At least one must be provided |
| `WithFlagsMutuallyExclusive(flags...)` | At most one may be provided |

## Subcommands

| Subcommand | Constraint demonstrated |
|------------|------------------------|
| `required` | `--output` must always be set |
| `together` | `--user` and `--password` must both be provided or both omitted |
| `one-of` | At least one of `--file` or `--url` must be provided |
| `exclusive` | `--json` and `--yaml` cannot be used together |

## Running

```
cd examples

# Required flag: omitting --output returns an error
go run ./flag-constraints required --output result.txt

# Together: must provide both or neither
go run ./flag-constraints together --user alice --password secret

# One-of: must provide at least one
go run ./flag-constraints one-of --file data.csv

# Exclusive: cannot use --json and --yaml simultaneously
go run ./flag-constraints exclusive --json
```

## Key API

```go
cmd.WithRequired("output")
cmd.WithFlagsRequiredTogether("user", "password")
cmd.WithFlagsOneRequired("file", "url")
cmd.WithFlagsMutuallyExclusive("json", "yaml")
```
