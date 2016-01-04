package auth

// IAuth interface
type IAuth interface {
	Signin() error
}

// Tables for package auth's
func Tables() []interface{} {
	return []interface{}{
		&AuthLocal{},
		&AuthCode{},
		&OAuth{},
		&AuthAPI{},

		&User{},
		&Role{},
		&Permission{},
	}
}

// CreateTables for package auth's
func CreateTables() {
	DB.SingularTable(true)
	DB.AutoMigrate(Tables()...)
	_createRelations()
}
func _createRelations() {
	DB.SingularTable(true)
	// Add foreign key
	// 1st param : foreignkey field
	// 2nd param : destination table(id)
	// 3rd param : ONDELETE
	// 4th param : ONUPDATE
	DB.Model(&AuthLocal{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")
	DB.Model(&AuthCode{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")
	DB.Model(&OAuth{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")
	DB.Model(&AuthAPI{}).AddForeignKey("user_id", "user(id)", "RESTRICT", "RESTRICT")
}
