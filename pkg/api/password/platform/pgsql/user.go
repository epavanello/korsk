package pgsql

import (
	"github.com/epavanello/gorsk/pkg/models"
	"github.com/go-pg/pg/v9/orm"
)

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u User) View(db orm.DB, id int) (models.User, error) {
	user := models.User{Base: models.Base{ID: id}}
	err := db.Select(&user)
	return user, err
}

// Update updates user's info
func (u User) Update(db orm.DB, user models.User) error {
	return db.Update(&user)
}
