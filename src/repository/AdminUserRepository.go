package repository

import (
	"gorm.io/gorm"
	"sakaba.link/api/src/model"
)

// AdminUserRepository is responsible for reading from and writing to DB Table `admin_users`.
type AdminUserRepository struct {
	DB *gorm.DB
}

// GetAdminUserByEmail returns an admin user specfied by email.
func (c AdminUserRepository) GetAdminUserByEmail(email string) model.AdminUser {
	adminUser := model.AdminUser{}
	c.DB.Table("admin_users").
		Select("id", "email", "password").
		Where("email = ?", email).
		Scan(&adminUser)

	return adminUser
}
