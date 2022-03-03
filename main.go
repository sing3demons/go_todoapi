package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sing3demons/todo/todo"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&todo.Todo{})

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	handler := todo.NewTodoHandler(db)
	r.POST("/todos", handler.NewTask)
	r.Run(":8080")
}
