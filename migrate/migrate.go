package main

import (
	"github.com/hemang-go/go-crud/initializers"
	"github.com/hemang-go/go-crud/models"
)

func init() {
	initializers.ConnectToDB()
	initializers.Load_Env()
}
func main() {
	initializers.DB.AutoMigrate(&models.BlogPost{})
}
