package main

import (
	"fmt"
	"golang-training/database"
	"golang-training/routes"

	"github.com/gofiber/fiber/v2"
	m "golang-training/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"golang_test",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Dogs{})
}

func main() {
	app := fiber.New()
	initDatabase()
	routes.InetRoutes(app)
	app.Listen(":3000")

}
