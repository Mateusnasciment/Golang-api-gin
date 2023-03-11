package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: "1", Title: "Todo 1"},
	{ID: "2", Title: "Todo 2"},
	{ID: "3", Title: "Todo 3"},
}

func main() {
	r := gin.Default()
	r.GET("/todos", GetTodos)
	r.GET("/todos/:id", GetTodo)
	r.Run()
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	for _, item := range todos {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{})
}
