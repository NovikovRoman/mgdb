# Example Structure For A JSON Columns

```go
package models

import (
    "bytes"
    "database/sql/driver"
    "encoding/json"
    "errors"
    "fmt"
)

type SimpleStruct struct {
	DoNotSaveToDB string `db:"-"`
	Field1        int    `db:"field1"`
	Field2        string `db:"field2"`
	Field3        bool   `db:"field3"`
}

func (s SimpleStruct) String() string {
    b, _ := json.Marshal(s)
    return string(b)
}

func (s *SimpleStruct) Scan(val interface{}) (err error) {
    switch v := val.(type) {
    case []byte:
        if bytes.Compare(v, []byte("[]")) == 0 {
            return
        }
        err = json.Unmarshal(v, s)
        return

    case string:
        if v == "[]" {
            return
        }
        err = json.Unmarshal([]byte(v), s)
        return

    default:
        return errors.New(fmt.Sprintf("Unsupported type: %T", v))
    }
}

func (s SimpleStruct) Value() (driver.Value, error) {
    return json.Marshal(s)
}

func (s *SimpleStruct) ConvertValue() (string, error) {
    b, err := json.Marshal(s)
    if err != nil {
        return "[]", err
    }
    return string(b), nil
}
```