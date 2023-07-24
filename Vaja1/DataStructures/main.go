package DataStructures

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Todo struct {
	Id            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	DateAdded     string `json:"date_added"`
	PredictedDate string `json:"predicted_date"`
}
