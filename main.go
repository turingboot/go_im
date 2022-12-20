package main

import (
	"go_im/common/initialize"
	"go_im/common/router"
)

func main() {
	initialize.LoadConfig()
	initialize.MysqlConnector()
	router.Router()
}
