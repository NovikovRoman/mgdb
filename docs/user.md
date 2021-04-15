# User Model Example

```go
package models

import "time"

type User struct {
	ID        int64      `db:"id"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`

	// more fields …
	// IgnoreField string {{.Backtick}}db:"-"{{.Backtick}} - ignore field
}

// interface model -----------------------------

func (u User) Table() string {
	return "users"
}

func (u User) GetID() int64 {
	return u.ID
}

func (u User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u User) GetDeletedAt() *time.Time {
	return u.DeletedAt
}

// model methods -----------------------------
// more methods …

```