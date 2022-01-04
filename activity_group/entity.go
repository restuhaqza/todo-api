package activity_group

type ActivityGroup struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt *string `json:"deleted_at"`
}
