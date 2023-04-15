package todo

type TodoUpdateDto struct {
	ID       int64   `json:"id" validate:"required"`
	ActivityGroupId int64   `json:"activity_group_id" validate:"required"`
	IsActive     bool  `json:"is_active"`
	Priority   string  `json:"priority"  validate:"required"`
}
