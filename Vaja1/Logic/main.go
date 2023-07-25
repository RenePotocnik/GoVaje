package Logic

import (
	"vaja1/DB"
)

type Controller struct {
	db DB.DB
}

func NewController(db DB.DB) *Controller {
	return &Controller{
		db: db,
	}
}
