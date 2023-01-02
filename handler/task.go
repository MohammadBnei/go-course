package handler

import (
	"go-course/task"
	"net/http"
	"strconv"

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

func (ts *taskHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

// Store godoc
// @Summary      Stores a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        taskInput   body      task.InputTask  true  "Message body"
// @Success      201  {object}  task.Task
// @Failure      400  {object}	Response
// @Router       /task [post]
func (th *taskHandler) Store(c *gin.Context) {
	// Get json body
	var input task.InputTask
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newTask, err := th.taskService.Store(input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusCreated, newTask)
}

func (th *taskHandler) FetchAll(c *gin.Context) {
	tasks, err := th.taskService.FetchAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Data: tasks,
	})
}

func (th *taskHandler) FetchById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	task, err := th.taskService.FetchById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Data: task,
	})
}

func (th *taskHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	// Get json body
	var input task.InputTask
	err = c.ShouldBindJSON(&input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Cannot extract JSON body",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	uTask, err := th.taskService.Update(id, input)
	if err != nil {
		response := &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := &Response{
		Message: "New task created",
		Data:    uTask,
	}
	c.JSON(http.StatusCreated, response)
}

func (th *taskHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Wrong id parameter",
			Data:    err.Error(),
		})
		return
	}

	err = th.taskService.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{
			Success: false,
			Message: "Something went wrong",
			Data:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &Response{
		Message: "Task successfully deleted",
	})
}
