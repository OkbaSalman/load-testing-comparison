package main

import (
	"gitlab.com/go-api/initializers"
	"gitlab.com/go-api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}