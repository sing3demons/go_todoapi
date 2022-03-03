package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "sing"})

	r := gin.Default()
	r.GET("/users", func(ctx *gin.Context) {
		var u User
		db.First(&u)
		ctx.JSON(200, u)
	})
	r.Run(":8080")
}
