package controllers

import (
	"fmt"
	"net/http"
	"strconv"

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
	id, err := strconv.ParseInt(ctx.Query("activity_group_id"), 32, 8);
fmt.Println(id);
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "Success",
			"message": err.Error(),
		})

		ctx.Abort()

		return;
	}

	data := uc.todoService.GetAll(int(id))

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   data,
	})
}

func (uc *TodoController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.todoService.GetByID(id)

	if data.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": fmt.Sprintf("Todo Items with ID %d not found", id),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   data,
	})
}

func (uc *TodoController) Create(ctx *gin.Context) {
	data, err := uc.todoService.Create(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   data,
	})
}

func (uc *TodoController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todoModel := uc.todoService.GetByID(id)

	if todoModel.ID < 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": fmt.Sprintf("Todo Items with ID %d not found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.todoService.Update(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"message":   err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   uc.todoService.GetByID(id),
	})
}

func (uc *TodoController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	todoModel := uc.todoService.GetByID(id)

	if todoModel.ID < 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": fmt.Sprintf("Todo Items with ID %d not found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.todoService.Delete(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "Error",
			"message":   err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   "Berhasil menghapus data!",
	})
}
