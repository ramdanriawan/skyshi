package todo

type TodoCreateDto struct {
	ID       int64   `json:"id"`
	ActivityGroupId int64   `json:"activity_group_id" validate:"required"`
	IsActive     bool  `json:"is_active"`
	Priority   string  `json:"priority"  validate:"required"`
}
