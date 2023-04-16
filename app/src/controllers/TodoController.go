package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	todo "skyshi.com/src/entities/todo"
)

type TodoController struct {
	todoService todo.TodoService
	ctx         *gin.Context
}

func NewTodoController(todoService todo.TodoService, ctx *gin.Context) TodoController {
	return TodoController{todoService, ctx}
}

func (uc *TodoController) Index(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Query("activity_group_id"), 32, 8)
	fmt.Println(id)

	type DayAndTime struct {
	}

	days := []*DayAndTime{}

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": err.Error(),
			"data":    days,
		})

		ctx.Abort()

		return
	}

	data := uc.todoService.GetAll(int(id))

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *TodoController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.todoService.GetByID(int64(id))

	type DayAndTime struct {
	}

	days := []*DayAndTime{}

	if data.Id == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Todo with ID %d Not Found", id),
			"data":    days,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *TodoController) Create(ctx *gin.Context) {

	data, err := uc.todoService.Create(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "Title") {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "title cannot be null",
			})

		} else if strings.Contains(err.Error(), "ActivityGroupId") {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "activity_group_id cannot be null",
			})
		} else {
			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": err.Error(),
			})
		}

		ctx.Abort()

		return
	}

	ctx.JSON(201, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *TodoController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todoModel := uc.todoService.GetByID(int64(id))

	if todoModel.Id < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Todo with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.todoService.Update(ctx)

	if err != nil {

		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    uc.todoService.GetByID(int64(id)),
	})
}

func (uc *TodoController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todoModel := uc.todoService.GetByID(int64(id))

	if todoModel.Id < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Todo with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.todoService.Delete(ctx)

	if err != nil {
		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	type ResponseData struct {
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    ResponseData{},
	})
}
