package templates

const MigrateReadme = `# Migration Tools

Build and run:
{{.Backtick}}{{.Backtick}}{{.Backtick}}
cd migrate && go build -o "migrate"
migrate user:pass@tcp(localhost:3306)/dbname?parseTime=true&multiStatements=true
{{.Backtick}}{{.Backtick}}{{.Backtick}}

**Optional keys:**
- {{.Backtick}}-u{{.Backtick}} migration up. Default one step.

- {{.Backtick}}-d{{.Backtick}} rollback migration. Default one step.

- {{.Backtick}}-sn{{.Backtick}} - {{.Backtick}}n{{.Backtick}} migration steps.

- {{.Backtick}}-f{{.Backtick}} force the current version of the migration to be installed.

Running without optional keys will execute all required versions of the migration.

You cannot specify the keys {{.Backtick}}-u{{.Backtick}} and {{.Backtick}}-d{{.Backtick}} at the same time. Only {{.Backtick}}-u{{.Backtick}} will be executed.

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
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	// Подключение к БД
	dbConn = kingpin.Arg("db", "Подключение к БД.").String()
	up     = kingpin.Flag("up", "Миграция вперед.").Short('u').Bool()
	down   = kingpin.Flag("down", "Миграция назад.").Short('d').Bool()
	step   = kingpin.Flag("step", "Количество шагов.").Short('s').Default("1").Int()
	force  = kingpin.Flag("force", "Принудительно выполнить текущую version.").Bool()
)

func main() {
	var (
		err     error
		db      *sql.DB
		driver  database.Driver
		m       *migrate.Migrate
		version uint
	)

	kingpin.HelpFlag.Short('h')
	kingpin.Parse()

	if *dbConn == "" {
		log.Fatalln("db variable not set.")
	}

	if db, err = sql.Open("mysql", *dbConn); err != nil {
		log.Fatalln(err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	driver, err = mysql.WithInstance(db, &mysql.Config{
		MigrationsTable: "schema_migrations",
		DatabaseName:    "migrations",
	})
	if err != nil {
		log.Fatalln(err)
	}

	m, err = migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"migrations", driver)
	if err != nil {
		log.Fatalln(err)
	}

	if *force {
		version, _, err = m.Version()
		if err != nil {
			log.Fatalln(err)
		}
		if err = m.Force(int(version)); err != nil {
			log.Fatalln(err)
		}
	}

	if *up {
		err = m.Steps(*step)

	} else if *down {
		err = m.Steps(-*step)

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
