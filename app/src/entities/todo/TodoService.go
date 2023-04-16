package todo

import (
	"fmt"
	// "io/ioutil"
	"strconv"

	dto "skyshi.com/src/entities/todo/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type TodoService interface {
	GetAll(activityGroupId int) []TodoModel
	GetByID(id int64) TodoModel
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

func (us *TodoServiceImpl) GetByID(id int64) TodoModel {
	return us.todoRepository.FindOne(id)
}

func (us *TodoServiceImpl) Create(ctx *gin.Context) (*TodoModel, error) {
	var input dto.TodoCreateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {

		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println(input)

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}
	// fmt.Println(input.IsActive);
	// fmt.Println("input.IsActive");
	todo := TodoModel{
		Priority:        input.Priority,
		ActivityGroupId: input.ActivityGroupId,
		Title:           input.Title,
		IsActive:        true,
		// IsActive:        bool(input.IsActive),
	}

	if input.Priority == "" {
		input.Priority = "very-high"

		todo = TodoModel{
			Priority:        input.Priority,
			ActivityGroupId: input.ActivityGroupId,
			Title:           input.Title,
			IsActive:        true,
			// IsActive:        bool(input.IsActive),
		}
	}

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
		Id:    int64(id),
		Title: input.Title,
		// ActivityGroupId: input.ActivityGroupId,
		IsActive: bool(input.IsActive),
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
		Id: int64(id),
	}

	result, err := us.todoRepository.Delete(todo)

	if err != nil {
		return nil, err
	}

	return result, nil
}
