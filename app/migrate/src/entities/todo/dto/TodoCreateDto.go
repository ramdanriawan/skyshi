package todo

type TodoCreateDto struct {
	Id              int64  `json:"id"`
	Title           string `json:"title"  validate:"required"`
	ActivityGroupId int64  `json:"activity_group_id" validate:"required"`
	IsActive        bool   `json:"is_active"`
	Priority        string `json:"priority"`
}
