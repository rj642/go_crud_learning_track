package main

import (
	"example.com/m/v2/initializers"
	"example.com/m/v2/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// To create a table, use this command written below
	initializers.DB.AutoMigrate(&models.Post{})
}

/**
run this migrate.go file to create an empty table with the desired model class.
go run migrate.go -> It will create an empty table with schema as post class
**/
