package DataStructures

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	PassHash string `json:"pass_hash" binding:"required"`
}

type Task struct {
	Id            int    `json:"id"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
	DateAdded     string `json:"date_added" binding:"required"`
	PredictedDate string `json:"predicted_date" binding:"required"`
	UserId        int    `json:"user_id"`
}
