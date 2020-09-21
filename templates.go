package main

var tmplModelInterface = `package {{.Package}}

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

type Repository struct {
	db    *sqlx.DB
	table string
}

type SimpleModel interface {
	Table() string
}

type Model interface {
	Table() string
	GetID() int64

	GetUpdatedAt() time.Time
	GetCreatedAt() time.Time
	GetDeletedAt() *time.Time
}

func Create(m Model, db *sqlx.DB) (id int64, err error) {
	var (
		res    sql.Result
		set    string
		values string
	)
	if m.GetID() > 0 {
		return Save(m, db)
	}

	setCreatedAt(m, time.Now())
	if m.GetUpdatedAt().IsZero() {
		setUpdatedAt(m, time.Now())
	}
	if m.GetDeletedAt() != nil && m.GetDeletedAt().IsZero() {
		setDeletedAt(m, nil)
	}

	if set, values, err = fieldsForInsert(m); err != nil {
		return
	}

	res, err = db.NamedExec("INSERT INTO {{.Backtick}}"+m.Table()+"{{.Backtick}} ("+set+") VALUES ("+values+")", m)
	if err == nil {
		id, err = res.LastInsertId()
		setID(m, id)
	}
	return
}

func Save(m Model, db *sqlx.DB) (id int64, err error) {
	if m.GetID() == 0 {
		return Create(m, db)
	}
	setUpdatedAt(m, time.Now())

	var sqlSet string
	if sqlSet, err = fieldsForUpdate(m); err != nil {
		return
	}
	_, err = db.NamedExec("UPDATE {{.Backtick}}"+m.Table()+"{{.Backtick}} SET "+sqlSet+" WHERE id=:id", m)
	if err == nil {
		id = m.GetID()
	}
	return
}

func Update(m Model, db *sqlx.DB) (id int64, err error) {
	if m.GetID() > 0 {
		return Save(m, db)
	}
	err = errors.New("This is a new entry. ")
	return
}

func setID(values Model, id int64) {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("ID").SetInt(id)
}

func setCreatedAt(values Model, t time.Time) {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("CreatedAt").Set(reflect.ValueOf(t))
}

func setUpdatedAt(values Model, t time.Time) {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	v.FieldByName("UpdatedAt").Set(reflect.ValueOf(t))
}

func setDeletedAt(values Model, t *time.Time) {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if t != nil {
		v.FieldByName("DeletedAt").SetPointer(unsafe.Pointer(t))
		return
	}
	v.FieldByName("DeletedAt").SetPointer(nil)
}

func fieldsForInsert(model interface{}) (set string, values string, err error) {
	var fields []string
	if fields, err = tableFields(model); err != nil {
		return
	}

	sqlValues := make([]string, len(fields))
	for i, name := range fields {
		fields[i] = "{{.Backtick}}" + name + "{{.Backtick}}"
		sqlValues[i] = ":" + name
	}

	set = strings.Join(fields, ",")
	values = strings.Join(sqlValues, ",")
	return
}

func fieldsForUpdate(model interface{}) (set string, err error) {
	var fields []string
	if fields, err = tableFields(model); err != nil {
		return
	}

	for i, name := range fields {
		fields[i] = "{{.Backtick}}" + name + "{{.Backtick}}=:" + name
	}
	set = strings.Join(fields, ",")
	return
}

func tableFields(values interface{}) (fields []string, err error) {
	v := reflect.ValueOf(values)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	fields = []string{}
	switch {
	case v.Kind() == reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i).Tag.Get("db")
			if field == "-" {
				continue

			} else if field == "" {
				fields = append(fields, strings.ToLower(v.Type().Field(i).Name))
				continue
			}

			fields = append(fields, field)
		}
		return

	case v.Kind() == reflect.Map:
		fields = make([]string, len(v.MapKeys()))
		for i, k := range v.MapKeys() {
			fields[i] = k.String()
		}
		return
	}

	err = fmt.Errorf("dbTableFields requires a struct or a map, found: %s", v.Kind().String())
	return
}
`

var tmplModel = `package {{.Package}}

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

var tmplRepository = `package {{.Package}}

import (
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
func (r {{.Model}}Repository) FindByID(id int64) ({{.ModelSymb}} *{{.Model}}, err error) {
	{{.ModelSymb}} = &{{.Model}}{}
	err = r.db.Get({{.ModelSymb}}, "SELECT * FROM {{.Backtick}}"+r.table+"{{.Backtick}} WHERE {{.Backtick}}id{{.Backtick}} = ?", id)
	if err == sql.ErrNoRows {
		err = nil
		{{.ModelSymb}} = nil
	}
	return
}

// FindAll returns all records from the database.
func (r {{.Model}}Repository) FindAll() ([]*{{.Model}}, error) {
	var {{.SliceModelName}} []*{{.Model}}
	err := r.db.Select(&{{.SliceModelName}}, "SELECT {{.Backtick}}{{.ModelSymb}}{{.Backtick}}.* FROM {{.Backtick}}"+r.table+"{{.Backtick}}")
	return {{.SliceModelName}}, err
}
`

var tmplStringArray = `package {{.Package}}

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type {{.Struct}} []string

func ({{.StructSymb}} {{.Struct}}) String() string {
	b, _ := json.Marshal({{.StructSymb}})
	return string(b)
}

func ({{.StructSymb}} *{{.Struct}}) Scan(val interface{}) (err error) {
	switch v := val.(type) {
	case []byte:
		return json.Unmarshal(v, &{{.StructSymb}})

	case string:
		return json.Unmarshal([]byte(v), &{{.StructSymb}})

	default:
		return fmt.Errorf("Unsupported type: %T. ", v)
	}
}

func ({{.StructSymb}} {{.Struct}}) Value() (driver.Value, error) {
	return json.Marshal({{.StructSymb}})
}
`
