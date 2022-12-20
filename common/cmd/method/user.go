package method

import "gorm.io/gen"

type UserMethod interface {
	// GetUserById
	//
	// where("id=@ID")
	GetUserById(ID int) (gen.T, error)

	// GetUserByName
	//
	// where("username=@username")
	GetUserByName(username string) (gen.T, error)

	// GetUserList
	//
	// sql(SELECT * FROM `user` WHERE  deleted_at is NULL)
	GetUserList() ([]gen.T, error)
}
