# Migration Tools

Build and run:
```
cd migrate && go build -o "migrate"
migrate user:pass@tcp(localhost:3306)/dbname?parseTime=true&multiStatements=true
```

**Required argument:**
- `--db=…` database connection

**Optional keys:**
- `--up, -u` migration up one step.

- `--down, -d` rollback migration one step.

- `--force=…` force the specified version of the migration. (--force=1 execute current).

Running without additional keys will launch all versions of migrations if necessary.

Keys cannot be set at the same time `--up` and `--down`.

## Naming Migration

Migration up:

```
[version]_[name].up.sql
```

Rollback migration:

```
[version]_[name].down.sql
```

`version` - unsigned integer. Usually as a date and time in the format `YYYYmmddHHii`.
`name` - the name of the migration for convenience.