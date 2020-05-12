# User Model Repository Example

```go
package models

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type UserRepository Repository

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db:    db,
		table: User{}.Table(),
	}
}

// FindByID returns a record from the database by ID.
func (r UserRepository) FindByID(id int64) (u *User, err error) {
	u = &User{}
	err = r.db.Get(u, "SELECT * FROM `"+r.table+"` WHERE `id` = ?", id)
	if err == sql.ErrNoRows {
		err = nil
		u = nil
	}
	return
}

// FindAll returns all records from the database.
func (r UserRepository) FindAll() ([]User, error) {
	var users []User
	err := r.db.Select(&users, "SELECT `u`.* FROM `"+r.table+"`")
	return users, err
}

```