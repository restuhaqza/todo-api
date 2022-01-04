package activity_group

type CreateActivityGroupInput struct {
	Title string `json:"title"`
	Email string `json:"email"`
}

type UpdateInput struct {
	Title *string `json:"title"`
	Email *string `json:"email"`
}
