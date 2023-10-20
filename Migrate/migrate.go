package main

import (
	"JSON_CRUD_API/initalizers"
	models "JSON_CRUD_API/model"
)

func init() {
	initalizers.LoadEnvvariables()
	initalizers.ConnectToDB()
}

func main() {
	initalizers.DB.AutoMigrate(&models.Post{})
}
