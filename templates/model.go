package templates

const Model = `package {{.Package}}

import "time"

type {{.Model}} struct {
	ID         int64 {{.Backtick}}db:"id"{{.Backtick}}
	CreatedAt time.Time  {{.Backtick}}db:"created_at"{{.Backtick}}
	UpdatedAt time.Time  {{.Backtick}}db:"updated_at"{{.Backtick}}
	DeletedAt *time.Time {{.Backtick}}db:"deleted_at"{{.Backtick}}

	// more fields
    // IgnoreField string {{.Backtick}}db:"-"{{.Backtick}} - ignore field
}

// interface model -----------------------------

func ({{.ModelSymb}} {{.Model}}) Table() string {
	return "{{.TableName}}"
}

func ({{.ModelSymb}} {{.Model}}) GetID() int64 {
	return {{.ModelSymb}}.ID
}

func ({{.ModelSymb}} {{.Model}}) GetUpdatedAt() time.Time {
	return {{.ModelSymb}}.UpdatedAt
}

func ({{.ModelSymb}} {{.Model}}) GetCreatedAt() time.Time {
	return {{.ModelSymb}}.CreatedAt
}

func ({{.ModelSymb}} {{.Model}}) GetDeletedAt() *time.Time {
	return {{.ModelSymb}}.DeletedAt
}

// model methods -----------------------------
// more methods …
`