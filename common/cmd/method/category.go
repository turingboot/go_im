package method

import "gorm.io/gen"

type CategoryMethod interface {

	// GetCategoryById
	//
	// where("id=@ID")
	GetCategoryById(ID int) (gen.T, error)

	// GetCategoryByName
	//
	// where("name=@name")
	GetCategoryByName(name string) (gen.T, error)

	// GetCategoryList
	//
	// sql(SELECT * FROM `category`)
	GetCategoryList() ([]gen.T, error)
}
