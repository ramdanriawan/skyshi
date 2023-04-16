package activityGroup

import (
	// "fmt"
	// "fmt"
	"strconv"

	dto "skyshi.com/src/entities/activityGroup/dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ActivityGroupService interface {
	GetAll() []ActivityGroupModel
	GetByID(id int) ActivityGroupModel
	Create(ctx *gin.Context) (*ActivityGroupModel, error)
	Update(ctx *gin.Context) (*ActivityGroupModel, error)
	Delete(ctx *gin.Context) (*ActivityGroupModel, error)
}

type ActivityGroupServiceImpl struct {
	activityGroupRepository ActivityGroupRepository
}

func NewActivityGroupService(todotransactionRepository ActivityGroupRepository) ActivityGroupService {
	return &ActivityGroupServiceImpl{todotransactionRepository}
}

func (us *ActivityGroupServiceImpl) GetAll() []ActivityGroupModel {
	return us.activityGroupRepository.FindAll()
}

func (us *ActivityGroupServiceImpl) GetByID(id int) ActivityGroupModel {
	return us.activityGroupRepository.FindOne(id)
}

func (us *ActivityGroupServiceImpl) Create(ctx *gin.Context) (*ActivityGroupModel, error) {
	var input dto.ActivityGroupCreateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	activityGroup := ActivityGroupModel{
		Email: input.Email,
		Title: input.Title,
	}

	result, err := us.activityGroupRepository.Save(activityGroup)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *ActivityGroupServiceImpl) Update(ctx *gin.Context) (*ActivityGroupModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var input dto.ActivityGroupUpdateDto

	if err := ctx.ShouldBindJSON(&input); err != nil {
		return nil, err
	}

	validate := validator.New()

	err := validate.Struct(input)

	if err != nil {
		return nil, err
	}

	activityGroup := ActivityGroupModel{
		Id:    int64(id),
		Title: input.Title,
		Email: input.Email,
	}

	result, err := us.activityGroupRepository.Update(activityGroup)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (us *ActivityGroupServiceImpl) Delete(ctx *gin.Context) (*ActivityGroupModel, error) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	todotransaction := ActivityGroupModel{
		Id: int64(id),
	}

	result, err := us.activityGroupRepository.Delete(todotransaction)

	if err != nil {
		return nil, err
	}

	return result, nil
}
