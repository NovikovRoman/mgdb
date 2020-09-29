package templates

const MigrateReadme = `# Migration Tools

Build and run:
{{.Backtick}}{{.Backtick}}{{.Backtick}}
cd migrate && go build -o "migrate"
migrate user:pass@tcp(localhost:3306)/dbname?parseTime=true&multiStatements=true
{{.Backtick}}{{.Backtick}}{{.Backtick}}

**Required argument:**
- {{.Backtick}}--db=…{{.Backtick}} database connection

**Optional keys:**
- {{.Backtick}}--up, -u{{.Backtick}} migration up one step.

- {{.Backtick}}--down, -d{{.Backtick}} rollback migration one step.

- {{.Backtick}}--force=…{{.Backtick}} force the specified version of the migration. (--force=1 execute current).

Running without additional keys will launch all versions of migrations if necessary.

Keys cannot be set at the same time {{.Backtick}}--up{{.Backtick}} and {{.Backtick}}--down{{.Backtick}}.

## Naming Migration

Migration up:

{{.Backtick}}{{.Backtick}}{{.Backtick}}
[version]_[name].up.sql
{{.Backtick}}{{.Backtick}}{{.Backtick}}

Rollback migration:

{{.Backtick}}{{.Backtick}}{{.Backtick}}
[version]_[name].down.sql
{{.Backtick}}{{.Backtick}}{{.Backtick}}

{{.Backtick}}version{{.Backtick}} - unsigned integer. Usually as a date and time in the format {{.Backtick}}YYYYmmddHHii{{.Backtick}}.
{{.Backtick}}name{{.Backtick}} - the name of the migration for convenience.`

const MigrateMain = `package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/mysql"
    _ "github.com/golang-migrate/migrate/v4/source/file"
    log "github.com/sirupsen/logrus"
    "gopkg.in/alecthomas/kingpin.v2"
)

var (
    dbConn = kingpin.Arg("db", "Подключение к БД").Required().String()
    down   = kingpin.Flag("down", "Миграция назад на 1 шаг").Short('d').
        Default("false").Bool()
    up = kingpin.Flag("up", "Миграция вперед на 1 шаг").Short('u').
        Default("false").Bool()
    force = kingpin.Arg("force", "Принудительно выполнить version").
        Default("").Uint()
)

func main() {
    kingpin.HelpFlag.Short('h')
    kingpin.Parse()

    if *down && *up {
        log.Fatalln("Выберите одно направление миграции (--up или --down)")
    }

    db, err := sql.Open("mysql", *dbConn)
    if err != nil {
        log.Fatalln(err)
    }

    defer func() {
        if err := db.Close(); err != nil {
            log.Fatal(err)
        }
    }()

    driver, err := mysql.WithInstance(db, &mysql.Config{
        MigrationsTable: "schema_migrations",
        DatabaseName:    "migrations",
    })
    if err != nil {
        log.Fatalln(err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "migrations", driver)
    if err != nil {
        log.Fatalln(err)
    }

    if *force != 0 {
        version := *force
        if *force > 1 {
            if version, _, err = m.Version(); err != nil {
                log.Fatalln(err)
            }
        }

        if err = m.Force(int(version)); err != nil {
            log.Fatalln(err)
        }
    }

    if *down {
        err = m.Steps(-1)

    } else if *up {
        err = m.Steps(1)

    } else {
        err = m.Up()
    }

    if err != nil {
        log.Fatalln(err)
    }
}
`

const MigrateUpSql = `CREATE TABLE IF NOT EXISTS {{.Backtick}}proxy{{.Backtick}}
(
    {{.Backtick}}id{{.Backtick}}         bigint(20)                              NOT NULL AUTO_INCREMENT PRIMARY KEY,
    {{.Backtick}}host{{.Backtick}}       varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL UNIQUE,
    {{.Backtick}}success{{.Backtick}}    boolean                                 NOT NULL,

    {{.Backtick}}created_at{{.Backtick}} datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    {{.Backtick}}updated_at{{.Backtick}} datetime                                NOT NULL DEFAULT CURRENT_TIMESTAMP,
    {{.Backtick}}deleted_at{{.Backtick}} datetime                                         DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;`

const MigrateDownSql = `DROP TABLE IF EXISTS {{.Backtick}}proxy{{.Backtick}};`
