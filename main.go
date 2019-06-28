package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/satoukick/webserver/logs"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=uqee.com sslmode=false")
	if err != nil {
		logs.Logger.Error("can not connect to database")
	}
	db.AutoMigrate(&todoModel{})
}

type todoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

type transformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"Completed"`
}

func main() {
	router := gin.Default()

	v1 := router.Group("api/v1/todos")
	{
		v1.POST("/", createTodo)
		// v1.GET("/", fetchAllTodo)
		// v1.GET("/:id", fetchSingleTodo)
		// v1.PUT("/:id", updateTodo)
		// v1.DELETE("/:id", deleteTodo)
	}
	router.Run()
}

func createTodo(c *gin.Context) {
	completed, _ := strconv.Atoi(c.PostForm("completed"))
	todo := todoModel{
		Title:     c.PostForm("title"),
		Completed: completed,
	}
	db.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "Todo item created successfully!",
		"resourceId": todo.ID,
	})
}
