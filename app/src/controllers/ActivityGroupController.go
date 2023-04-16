package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *ActivityGroupController) GetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data := uc.activityGroupService.GetByID(id)

	fmt.Println(id)
	fmt.Println("53636345635346346334524525")

	if data.Id == 0 {
		ctx.JSON(404, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Activity with ID %d Not Found", id),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	})
}

func (uc *ActivityGroupController) Create(ctx *gin.Context) {
	data, err := uc.activityGroupService.Create(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Title") {

			ctx.JSON(400, gin.H{
				"status":  "Bad Request",
				"message": "title cannot be null",
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

func (uc *ActivityGroupController) Update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	activityGroupModel := uc.activityGroupService.GetByID(id)

	if activityGroupModel.Id < 1 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Activity with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.activityGroupService.Update(ctx)

	if err != nil {

		ctx.JSON(400, gin.H{
			"status":  "Bad Request",
			"message": err.Error(),
		})

		ctx.Abort()

		return
	}

	ctx.JSON(200, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    uc.activityGroupService.GetByID(id),
	})
}

func (uc *ActivityGroupController) Delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	activityGroupModel := uc.activityGroupService.GetByID(id)
	fmt.Println(id)
	fmt.Println("idsdfojdsfjdlsfjdklsfjkdsjfdiklshjfklidshfkldshfjkdhsfjkhsdfjkhsdjkfghdjksfgdjksgfhjdsf")
	if activityGroupModel.Id < 1 {
		ctx.JSON(404, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Activity with ID %d Not Found", id),
		})

		ctx.Abort()

		return
	}

	_, err := uc.activityGroupService.Delete(ctx)

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
