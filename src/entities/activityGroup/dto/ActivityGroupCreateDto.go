package activityGroup

type ActivityGroupCreateDto struct {
	Title  string   `json:"title" validate:"required,max=100,min=1"`
	Email    string   `json:"email" validate:"required,email"`
}
