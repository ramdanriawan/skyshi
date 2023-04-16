package activityGroup

type ActivityGroupCreateDto struct {
	Id    int64  `json:"id"`
	Title string `json:"title" validate:"required,max=100,min=1"`
	Email string `json:"email" validate:"required,email"`
}
