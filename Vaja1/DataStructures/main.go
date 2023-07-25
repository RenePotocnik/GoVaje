package DataStructures

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type Task struct {
	Id            int    `json:"id"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
	DateAdded     string `json:"date_added" binding:"required"`
	PredictedDate string `json:"predicted_date" binding:"required"`
}
