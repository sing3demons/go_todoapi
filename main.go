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

type UserHandler struct {
	db *gorm.DB
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	db.Create(&User{Name: "sing"})

	userHandler := UserHandler{db: db}

	r := gin.Default()
	r.GET("/users",userHandler.User)
	r.Run(":8080")
}

func (h *UserHandler) User(ctx *gin.Context) {
	var u User
	h.db.First(&u)
	ctx.JSON(200, u)
}
