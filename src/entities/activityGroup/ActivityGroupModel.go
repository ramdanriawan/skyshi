package activityGroup

import "time"

type ActivityGroupModel struct {
	ID        int64  `json:"id" gorm:"primaryKey;auto_increment:true;index"`
	Title     string `json:"title" gorm:"type:varchar(100)"`
	Email     string `json:"email" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp"`
}

func (ActivityGroupModel) TableName() string {
	return "activity_groups"
}
