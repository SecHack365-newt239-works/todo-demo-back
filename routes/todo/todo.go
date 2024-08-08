package todo

import (
	"fmt"
	"net/http"
	"time"
	"todo-demo-back/utils"

	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID        string `gorm:"primaryKey;size:255;default:uuid_generate_v4()"`
	Label     string
	Timestamp time.Time
	Done      bool
}

type createTodoParam struct {
	Label string `json:"label"`
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
		return c.NoContent(http.StatusOK)
	}
}

func GetTodoByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := utils.GetDBConnection()
		if err != nil {
			fmt.Printf("Error getting DB connection: %v\n", err)
		}
		id := c.Param("id")
		var todo Todo
		db.First(&todo, "id = ?", id)
		return c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodoDone() echo.HandlerFunc {
	return func(c echo.Context) error {
		db, err := utils.GetDBConnection()
		if err != nil {
			fmt.Printf("Error getting DB connection: %v\n", err)
		}
		id := c.Param("id")
		status := c.QueryParam("status")
		if status == "true" {
			db.Model(&Todo{}).Where("id = ?", id).Update("Done", true)
		} else {
			db.Model(&Todo{}).Where("id = ?", id).Update("Done", false)
		}
		return c.NoContent(http.StatusOK)
	}
}
