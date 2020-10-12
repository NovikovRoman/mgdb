# Example Array Structure For A JSON Columns

```go
package models

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
)

type StringArray []string

func (s StringArray) String() string {
    b, _ := json.Marshal(s)
    return string(b)
}

func (s *StringArray) Scan(val interface{}) (err error) {
    switch v := val.(type) {
    case []byte:
        return json.Unmarshal(v, &s)

    case string:
        return json.Unmarshal([]byte(v), &s)

    default:
        return fmt.Errorf("Unsupported type: %T. ", v)
    }
}

func (s StringArray) Value() (driver.Value, error) {
    return json.Marshal(s)
}

```