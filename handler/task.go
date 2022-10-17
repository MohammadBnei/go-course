package handler

import (
	"fmt"
	"go-course/task"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type taskHandler struct {
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService}
}

func (handler *taskHandler) Store(c *gin.Context) {
	var input task.InputTask
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTask, err := handler.taskService.Store(input)
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Message: "new task successfully added",
		Data:    newTask,
	}

	c.JSON(http.StatusOK, response)
}

func (handler *taskHandler) List(c *gin.Context) {

	tasks, err := handler.taskService.ListAll()
	if err != nil {
		response := Response{
			Success: false,
			Message: "something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := Response{
		Success: true,
		Data:    tasks,
	}

	fmt.Println(tasks)

	c.JSON(http.StatusOK, response)
}
