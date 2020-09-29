package templates

const Repository = `package {{.Package}}

import (
    "context"
    "database/sql"
    "github.com/jmoiron/sqlx"
)

type {{.Model}}Repository Repository

func New{{.Model}}Repository(db *sqlx.DB) *{{.Model}}Repository {
    return &{{.Model}}Repository{
        db:    db,
        table: {{.Model}}{}.Table(),
    }
}

// FindByID returns a record from the database by ID.
func (r {{.Model}}Repository) FindByID(ctx context.Context, id int64) ({{.ModelSymb}} *{{.Model}}, err error) {
    {{.ModelSymb}} = &{{.Model}}{}
	query := "SELECT * FROM {{.Backtick}}"+r.table+"{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}} = ?"

    if ctx == nil {
        err = r.db.Get({{.ModelSymb}}, query, id)
    } else {
        err = r.db.GetContext(ctx, {{.ModelSymb}}, query, id)
    }

    if err == sql.ErrNoRows {
        err = nil
        {{.ModelSymb}} = nil
    }
    return
}

// FindAll returns all records from the database.
func (r {{.Model}}Repository) FindAll(ctx context.Context) ({{.SliceModelName}} []*{{.Model}},err error) {
    {{.SliceModelName}} = []*{{.Model}}{}
    query := "SELECT {{.Backtick}}{{.ModelSymb}}{{.Backtick}}.* FROM {{.Backtick}}"+r.table+"{{.Backtick}}"

    if ctx == nil {
        err = r.db.Select(&{{.SliceModelName}}, query)
    } else {
        err = r.db.SelectContext(ctx, &{{.SliceModelName}}, query)
    }
    return
}
`
