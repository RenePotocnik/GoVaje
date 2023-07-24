package Logic

import (
	"os"
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

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
