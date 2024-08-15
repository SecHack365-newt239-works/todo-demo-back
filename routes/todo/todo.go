package todo

import (
	"fmt"
	"net/http"
	"time"
	"todo-demo-back/utils"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        int `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	Label     string
	Timestamp time.Time
	Done      bool
}

type createTodoParam struct {
	Label     string    `json:"label"`
	Timestamp time.Time `json:"timestamp"`
	Done      bool      `json:"done"`
}

type responseParam struct {
	ID int `json:"id"`
}

func CreateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		createTodoParam := createTodoParam{}
		if err := c.Bind(&createTodoParam); err != nil {
			return err
		}
		db, err := utils.GetDBConnection()
		if err != nil {
			fmt.Printf("Error getting DB connection: %v\n", err)
		}
		newTodo := Todo{Label: createTodoParam.Label, Timestamp: time.Now(), Done: false}
		result := db.Create(&newTodo)
		if result.Error != nil {
			fmt.Printf("Error creating user: %v\n", result.Error)
		}
		return c.JSON(http.StatusOK, responseParam{ID: newTodo.ID})
	}
}

func GetTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := utils.GetDBConnection()
		if err != nil {
			fmt.Printf("Error getting DB connection: %v\n", err)
		}
		var todos []Todo
		db.Where("done = false").Find(&todos)
		fmt.Printf("todos: %v\n", todos)
		return c.JSON(http.StatusOK, todos)
	}
}

func DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := utils.GetDBConnection()
		if err != nil {
			fmt.Printf("Error getting DB connection: %v\n", err)
		}
		id := c.Param("id")
		db.Model(&Todo{}).Where("id = ?", id).Update("done", true)

		return c.NoContent(http.StatusOK)
	}
}
