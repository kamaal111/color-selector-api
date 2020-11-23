package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// User database model
type User struct {
	tableName  struct{}      `pg:"users"`
	ID         int           `pg:"id,pk"`
	Email      string        `pg:"email,unique"`
	CreatedAt  time.Time     `pg:"created_at"`
	UpdatedAt  time.Time     `pg:"updated_at"`
	SavedColor []*SavedColor `pg:"saved_colors,rel:has-many"`
}

// Save : This function inserts a user in to the database
// table of the model
func (user *User) Save(pgDB *pg.DB) error {
	_, insertErr := pgDB.Model(user).Insert()
	return insertErr
}

// GetUserByID : This function returns a user with a specific primarykey
func GetUserByID(pgDB *pg.DB, userID int) (User, error) {
	var user User
	getUserErr := pgDB.Model((*User)(nil)).Where("id = ?", userID).Select(&user)
	return user, getUserErr
}

// GetUserByEmail : This function returns a user with a specific email
func GetUserByEmail(pgDB *pg.DB, email string) (User, error) {
	var user User
	getUserErr := pgDB.Model((*User)(nil)).Where("email = ?", email).Select(&user)
	return user, getUserErr
}

// CreateUsersTable : This function creates a users table,
// if it does not exist
func CreateUsersTable(pgDB *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createTableErr := pgDB.Model((*User)(nil)).CreateTable(options)
	return createTableErr
}
