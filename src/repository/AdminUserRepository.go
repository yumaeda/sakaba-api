package repository

import (
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
)

// AdminUserRepository is responsible for reading from and writing to DB Table `admin_users`.
type AdminUserRepository struct{}

// GetAdminUserByEmail returns an admin user specfied by email.
func (c AdminUserRepository) GetAdminUserByEmail(email string) model.AdminUser {
	db, closer, err := infrastructure.ConnectToTiDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	adminUser := model.AdminUser{}
	db.Table("admin_users").
		Select("id", "email", "password").
		Where("email = ?", email).
		Scan(&adminUser)

	return adminUser
}
