package templates

const Interface = `package {{.Package}}

import (
    "context"
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

func Create(ctx context.Context, db *sqlx.DB, m Model) (id int64, err error) {
    var (
        res    sql.Result
        set    string
        values string
    )
    if m.GetID() > 0 {
        return Save(ctx, db, m)
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
	
	query := "INSERT INTO {{.Backtick}}"+m.Table()+"{{.Backtick}} ("+set+") VALUES ("+values+")"
    if ctx == nil {
        res, err = db.NamedExec(query, m)
    } else {
        res, err = db.NamedExecContext(ctx, query, m)
    }

    if err == nil {
        id, err = res.LastInsertId()
        setID(m, id)
    }
    return
}

func Save(ctx context.Context, db *sqlx.DB, m Model) (id int64, err error) {
    if m.GetID() == 0 {
        return Create(ctx, db, m)
    }
    setUpdatedAt(m, time.Now())

    var sqlSet string
    if sqlSet, err = fieldsForUpdate(m); err != nil {
        return
    }

    query := "UPDATE {{.Backtick}}"+m.Table()+"{{.Backtick}} SET "+sqlSet+" WHERE id=:id"
    if ctx == nil {
        _, err = db.NamedExec(query, m)
    } else {
        _, err = db.NamedExecContext(ctx, query, m)
    }

    if err == nil {
        id = m.GetID()
    }
    return
}

func Update(ctx context.Context, db *sqlx.DB, m Model) (id int64, err error) {
    if m.GetID() > 0 {
        return Save(ctx, db, m)
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
