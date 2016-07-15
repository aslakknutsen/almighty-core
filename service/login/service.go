package login

import "github.com/jinzhu/gorm"

// Service defines the basic operations required to perform a remote oauth login\
// user registration
type Service interface {
	PerformIntialSetup()
}

// DBLoginService defins a basic Login Service
type DBLoginService struct {
	DB *gorm.DB
}


