package models

import (
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// SavedColor database model
type SavedColor struct {
	tableName struct{}  `pg:"saved_colors"`
	ID        int       `pg:"id,pk"`
	Name      string    `pg:"name,unique"`
	Hex       string    `pg:"hex"`
	CreatedAt time.Time `pg:"created_at"`
	UpdatedAt time.Time `pg:"updated_at"`
	UserID    int       `pg:"user_id"`
	User      *User     `pg:"user,rel:has-one"`
}

// Save : This function inserts a saved color in to the database
// table of the model
func (savedColor *SavedColor) Save(pgDB *pg.DB) error {
	_, insertErr := pgDB.Model(savedColor).Insert()
	return insertErr
}

// CreateSavedColorsTable : This function creates a saved colors table,
// if it does not exist
func CreateSavedColorsTable(pgDB *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createTableErr := pgDB.Model((*SavedColor)(nil)).CreateTable(options)
	return createTableErr
}
