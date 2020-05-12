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
func (r UserRepository) FindByID(id int64) (d *User, err error) {
	d = &User{}
	err = r.db.Get(d, "SELECT * FROM `"+r.table+"` WHERE `id` = ?", id)
	if err == sql.ErrNoRows {
		err = nil
		d = nil
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