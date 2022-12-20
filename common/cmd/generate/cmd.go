package main

import (
	"go_im/common/global"
	"go_im/common/initialize"
	"gorm.io/gen"
)

// generate code
func main() {
	// specify the output directory (default: "./query")
	// ### if you want to query without context constrain, set mode gen.WithoutContext ###
	g := gen.NewGenerator(gen.Config{
		//OutPath: "D:\\workspace\\golang\\PATH\\myBlog\\common\\cmd/dal/query",
		OutPath: "dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery,
		//if you want the nullable field generation property to be pointer type, set FieldNullable true
		/* FieldNullable: true,*/
		//if you want to generate index tags from database, set FieldWithIndexTag true
		/* FieldWithIndexTag: true,*/
		//if you want to generate type tags from database, set FieldWithTypeTag true
		/* FieldWithTypeTag: true,*/
		//if you need unit tests for query code, set WithUnitTest true
		//WithUnitTest: true,
	})

	// reuse the database connection in Project or create a connection here
	// if you want to use GenerateModel/GenerateModelAs, UseDB is necessary or it will panic
	// db, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))

	initialize.LoadConfig()
	initialize.MysqlConnector()
	g.UseDB(global.Db)

	// apply basic crud api on structs or table models which is specified by table name with function
	// GenerateModel/GenerateModelAs. And generator will generate table models' code when calling Excute.

	g.ApplyBasic(
		//g.GenerateModelAs("user", "User"),
		//g.GenerateModelAs("category", "Category"),
		g.GenerateModel("community"),
		g.GenerateModel("contact"),
		g.GenerateModel("user"),
	)
	//apply diy interfaces on structs or table models
	//g.ApplyInterface(
	//	func(method method.UserMethod) {},
	//	g.GenerateModelAs("user", "User"))
	//g.ApplyInterface(
	//	func(method method.CategoryMethod) {},
	//	g.GenerateModelAs("category", "Category"))

	// execute the action of code generation
	g.Execute()

}
