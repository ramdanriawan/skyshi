package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	activityGroup "skyshi.com/src/entities/activityGroup"
	todo "skyshi.com/src/entities/todo"
)

type ActivityGroupController struct {
	activityGroupService activityGroup.ActivityGroupService
	todoService          todo.TodoService
	todoRepository       todo.TodoRepository
	todoRepositoryImpl   todo.TodoRepositoryImpl
	ctx                  *gin.Context
}

func NewActivityGroupController(activityGroupService activityGroup.ActivityGroupService, todoService todo.TodoService, todoRepository todo.TodoRepository, todoRepositoryImpl todo.TodoRepositoryImpl, ctx *gin.Context) ActivityGroupController {
	return ActivityGroupController{activityGroupService, todoService, todoRepository, todoRepositoryImpl, ctx}
}

func (uc *ActivityGroupController) Index(ctx *gin.Context) {
	data := uc.activityGroupService.GetAll()

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   data,
	})
}

func (uc *ActivityGroupController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.activityGroupService.GetByID(id)

	if data.ID == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": fmt.Sprintf("Activity with ID %d not found", id),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "Success",
		"message": "Success",
		"data":   data,
	})
}

func (uc *ActivityGroupController) Create(ctx *gin.Context) {
	data, err := uc.activityGroupService.Create(ctx)

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

func (uc *ActivityGroupController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	activityGroupModel := uc.activityGroupService.GetByID(id)

	if activityGroupModel.ID < 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": fmt.Sprintf("Activity with ID %d not found", id),
		})

		ctx.Abort()

		return
	}

	data, err := uc.activityGroupService.Update(ctx)

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
		"data":   data,
	})
}

func (uc *ActivityGroupController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	activityGroupModel := uc.activityGroupService.GetByID(id)

	if activityGroupModel.ID < 1 {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Error",
			"message": fmt.Sprintf("Activity with ID %d not found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.activityGroupService.Delete(ctx)

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
