package todo

import (

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll(activityGroupId int) []TodoModel
	FindOne(id int64) TodoModel
	Save(todo TodoModel) (*TodoModel, error)
	Update(todo TodoModel) (*TodoModel, error)
	Delete(todo TodoModel) (*TodoModel, error)
}

type TodoRepositoryImpl struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &TodoRepositoryImpl{db}
}

func (ur *TodoRepositoryImpl) FindAll(activityGroupId int) []TodoModel {
	var todos []TodoModel

	_ = ur.db.Where("activity_group_id = ?", activityGroupId).Find(&todos)

	return todos

}

func (ur *TodoRepositoryImpl) FindOne(id int64) TodoModel {
	var todo TodoModel
	_ = ur.db.Find(&todo, id)
	
	return todo
}

func (ur *TodoRepositoryImpl) Save(todo TodoModel) (*TodoModel, error) {
	result := ur.db.Save(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}

func (ur *TodoRepositoryImpl) Update(todo TodoModel) (*TodoModel, error) {
	

	result := ur.db.Model(&todo).Updates(&todo)

	if result.Error != nil {
		
		return nil, result.Error
	}

	return &todo, nil
}

func (ur *TodoRepositoryImpl) Delete(todo TodoModel) (*TodoModel, error) {
	result := ur.db.Delete(&todo)

	if result.Error != nil {
		return nil, result.Error
	}

	return &todo, nil
}
