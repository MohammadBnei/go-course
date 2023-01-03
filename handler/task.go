package handler

import (
	"context"
	"go-course/gen"
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
	gen.UnimplementedTaskServiceServer
	taskService task.Service
}

func NewTaskHandler(taskService task.Service) *taskHandler {
	return &taskHandler{taskService: taskService}
}

func (ts *taskHandler) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}

// TaskStore godoc
// @Summary      Save a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        Task   body      task.InputTask  true "Task input"
// @Success      201  {object}  task.Task
// @Failure      400  {object}  Response
// @Router       /task [post]
func (th *taskHandler) Store(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	newTask, err := th.taskService.Store(task.InputTask{
		Name:        req.GetName(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, err
	}

	return &gen.Response{
		Task: &gen.Task{
			Id:          int32(newTask.ID),
			Name:        newTask.Name,
			Description: newTask.Description,
		},
	}, nil
}

// TaskShow godoc
// @Summary      Show all tasks
// @Tags         task
// @Accept       json
// @Produce      json
// @Success      200  {array}  task.Task
// @Failure      400  {object}  Response
// @Router       /task [get]
func (th *taskHandler) FetchAll(ctx context.Context, req *gen.Request) (*gen.Response, error) {
	tasks, err := th.taskService.FetchAll()
	if err != nil {
		return nil, err
	}

	res := &gen.Response{}
	for _, v := range tasks {
		res.Tasks = append(res.Tasks, &gen.Task{
			Id:          int32(v.ID),
			Name:        v.Name,
			Description: v.Description,
		})
	}

	return res, nil
}

// TaskShow godoc
// @Summary      Show a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  task.Task
// @Failure      400  {object}  Response
// @Router       /task/{id} [get]
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

	c.JSON(http.StatusOK, task)
}

// TaskUpdate godoc
// @Summary      Update a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Param        Task   body      task.InputTask  true "Task input"
// @Success      200  {object}  task.Task
// @Failure      400  {object}  Response
// @Router       /task/{id} [put]
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

	c.JSON(http.StatusOK, uTask)
}

// TaskDelete godoc
// @Summary      Delete a task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Task ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Router       /task/{id} [delete]
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
		Success: true,
		Message: "Task successfully deleted",
	})
}
