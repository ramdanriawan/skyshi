package todo

import "time"

type TodoModel struct {
	Id              int64     `json:"id" gorm:"primaryKey;auto_increment:true;index"`
	Title           string    `json:"title" gorm:"type:varchar(100)"`
	ActivityGroupId int64     `json:"activity_group_id" gorm:"type:bigint(20)"`
	IsActive        bool      `json:"is_active" gorm:"type:int(11)"`
	Priority        string    `json:"priority" gorm:"type:enum"`
	CreatedAt       time.Time `json:"createdAt" gorm:"type:timestamps"`
}

func (TodoModel) TableName() string {
	return "todos"
}
