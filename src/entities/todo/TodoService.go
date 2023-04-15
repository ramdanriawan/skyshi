package todo

import (
	"fmt"
	"strconv"

	dto "skyshi.com/src/entities/todo/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoService interface {
	GetAll(activityGroupId int) []TodoModel
	GetByID(id int) TodoModel
	Create(ctx *gin.Context) (*TodoModel, error)
	Update(ctx *gin.Context) (*TodoModel, error)
	Delete(ctx *gin.Context) (*TodoModel, error)
}

type TodoServiceImpl struct {
	todoRepository TodoRepository
}

func NewTodoService(todoRepository TodoRepository) TodoService {
	return &TodoServiceImpl{todoRepository}
}

func (us *TodoServiceImpl) GetAll(activityGroupId int) []TodoModel {
	return us.todoRepository.FindAll(activityGroupId)
}

func (us *TodoServiceImpl) GetByID(id int) TodoModel {
	return us.todoRepository.FindOne(id)
}

func (us *TodoServiceImpl) Create(ctx *gin.Context) (*TodoModel, error) {
	var input dto.TodoCreateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	todo := TodoModel{
		ActivityGroupId: input.ActivityGroupId,
		IsActive:     bool(input.IsActive),
		Priority:   input.Priority,
	}
fmt.Println(bool(input.IsActive));
	result, err := us.todoRepository.Save(todo)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *TodoServiceImpl) Update(ctx *gin.Context) (*TodoModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.TodoUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	todo := TodoModel{
		ID:       int64(id),
		ActivityGroupId: input.ActivityGroupId,
		IsActive:     bool(input.IsActive),
		Priority:   input.Priority,
	}

	result, err := us.todoRepository.Update(todo)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *TodoServiceImpl) Delete(ctx *gin.Context) (*TodoModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	todo := TodoModel{
		ID: int64(id),
	}

	result, err := us.todoRepository.Delete(todo)

	if err != nil {
		return nil, err
	}

	return result, nil
}
