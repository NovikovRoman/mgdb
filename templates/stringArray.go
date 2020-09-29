package templates

const StringArray = `package {{.Package}}

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
