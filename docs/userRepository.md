# User Model Repository Example

```go
package models

import (
    "context"
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
func (r UserRepository) FindByID(ctx context.Context, id int64) (u *User, err error) {
	u = &User{}

    query := "SELECT * FROM `"+r.table+"` WHERE `id` = ?"
    if ctx == nil {
        err = r.db.Get(u, query, id)
    } else {
        err = r.db.GetContext(ctx, u, query, id)
    }

	if err == sql.ErrNoRows {
		err = nil
		u = nil
	}
	return
}

// FindAll returns all records from the database.
func (r UserRepository) FindAll(ctx context.Context) (users []*User, err error) {
	users = []*User{}

    query := "SELECT `u`.* FROM `"+r.table+"`"
    if ctx == nil {
        err = r.db.Select(&users, query)
    } else {
        err = r.db.SelectContext(ctx, &users, query)
    }

	return
}

```