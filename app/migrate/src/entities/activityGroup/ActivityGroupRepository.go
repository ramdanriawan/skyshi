package activityGroup

import (
	"gorm.io/gorm"
)

type ActivityGroupRepository interface {
	FindAll() []ActivityGroupModel
	FindOne(id int) ActivityGroupModel
	Save(activityGroup ActivityGroupModel) (*ActivityGroupModel, error)
	Update(activityGroup ActivityGroupModel) (*ActivityGroupModel, error)
	Delete(activityGroup ActivityGroupModel) (*ActivityGroupModel, error)
}

type ActivityGroupRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityGroupRepository(db *gorm.DB) ActivityGroupRepository {
	return &ActivityGroupRepositoryImpl{db}
}

func (ur *ActivityGroupRepositoryImpl) FindAll() []ActivityGroupModel {
	var activityGroups []ActivityGroupModel

	_ = ur.db.Find(&activityGroups)

	return activityGroups

}

func (ur *ActivityGroupRepositoryImpl) FindOne(id int) ActivityGroupModel {
	var activityGroup ActivityGroupModel
	_ = ur.db.Find(&activityGroup, id)

	return activityGroup
}

func (ur *ActivityGroupRepositoryImpl) Save(activityGroup ActivityGroupModel) (*ActivityGroupModel, error) {
	result := ur.db.Save(&activityGroup)

	if result.Error != nil {
		return nil, result.Error
	}

	return &activityGroup, nil
}

func (ur *ActivityGroupRepositoryImpl) Update(activityGroup ActivityGroupModel) (*ActivityGroupModel, error) {
	result := ur.db.Model(&activityGroup).Updates(&activityGroup)

	if result.Error != nil {
		return nil, result.Error
	}

	return &activityGroup, nil
}

func (ur *ActivityGroupRepositoryImpl) Delete(activityGroup ActivityGroupModel) (*ActivityGroupModel, error) {
	result := ur.db.Delete(&activityGroup)

	if result.Error != nil {
		return nil, result.Error
	}

	return &activityGroup, nil
}
