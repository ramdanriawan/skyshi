package route

import (
	activityGroup "skyshi.com/src/entities/activityGroup"
	todo "skyshi.com/src/entities/todo"

	controllers "skyshi.com/src/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	ctx                *gin.Context
	todoRepositoryImpl todo.TodoRepositoryImpl
)

func Api(router *gin.Engine, db *gorm.DB) {
	todoRepository := todo.NewTodoRepository(db)
	todoService := todo.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService, ctx)

	activityGroupRepository := activityGroup.NewActivityGroupRepository(db)
	activityGroupService := activityGroup.NewActivityGroupService(activityGroupRepository)
	activityGroupController := controllers.NewActivityGroupController(activityGroupService, todoService, todoRepository, todoRepositoryImpl, ctx)

	v1 := router.Group("/")
	{

		v1.GET("/activity-groups", activityGroupController.Index)
		v1.POST("/activity-groups", activityGroupController.Create)
		v1.GET("/activity-groups/:id", activityGroupController.GetByID)
		v1.PATCH("/activity-groups/:id", activityGroupController.Update)
		v1.DELETE("/activity-groups/:id", activityGroupController.Delete)

		v1.GET("/todo-items", todoController.Index)
		v1.POST("/todo-items", todoController.Create)
		v1.GET("/todo-items/:id", todoController.GetByID)
		v1.PATCH("/todo-items/:id", todoController.Update)
		v1.DELETE("/todo-items/:id", todoController.Delete)
	}
}
