package activityGroup

import (
	"gorm.io/gorm"
)

type ActivityGroupRepository interface {
	FindAll() []ActivityGroupModel
	FindOne(id int) ActivityGroupModel
	Save(todotransaction ActivityGroupModel) (*ActivityGroupModel, error)
	Update(todotransaction ActivityGroupModel) (*ActivityGroupModel, error)
	Delete(todotransaction ActivityGroupModel) (*ActivityGroupModel, error)
}

type ActivityGroupRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityGroupRepository(db *gorm.DB) ActivityGroupRepository {
	return &ActivityGroupRepositoryImpl{db}
}

func (ur *ActivityGroupRepositoryImpl) FindAll() []ActivityGroupModel {
	var todotransactions []ActivityGroupModel

	_ = ur.db.Find(&todotransactions)

	return todotransactions

}

func (ur *ActivityGroupRepositoryImpl) FindOne(id int) ActivityGroupModel {
	var todotransaction ActivityGroupModel
	_ = ur.db.Find(&todotransaction, id)

	return todotransaction
}

func (ur *ActivityGroupRepositoryImpl) Save(todotransaction ActivityGroupModel) (*ActivityGroupModel, error) {
	result := ur.db.Save(&todotransaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todotransaction, nil
}

func (ur *ActivityGroupRepositoryImpl) Update(todotransaction ActivityGroupModel) (*ActivityGroupModel, error) {
	result := ur.db.Model(&todotransaction).Updates(&todotransaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todotransaction, nil
}

func (ur *ActivityGroupRepositoryImpl) Delete(todotransaction ActivityGroupModel) (*ActivityGroupModel, error) {
	result := ur.db.Delete(&todotransaction)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todotransaction, nil
}
