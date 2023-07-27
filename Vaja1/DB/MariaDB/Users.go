package MariaDB

import (
	"errors"
	"fmt"
	"vaja1/DataStructures"
)

// GetUserByUsername gets the user struct from the database given a username
func (db *MariaDB) GetUserByUsername(username string) (user DataStructures.User, err error) {
	rows, err := db.database.Query("SELECT user_id, username, pass_hash from user WHERE username = ? LIMIT 1", username)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	// Move to the next row (the first returned row), if it does not exist, it means that the response is empty and does not exist
	if !rows.Next() {
		err = errors.New("user does not exist")
		return
	}

	// Get the Queried data from the row into the user struct
	err = rows.Scan(&user.Id, &user.Username, &user.PassHash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}

func (db *MariaDB) Login(user DataStructures.User) (err error) {
	// Get the user from the database
	userFromDB, err := db.GetUserByUsername(user.Username)
	if err != nil {
		return
	}

	// Check if the password is correct
	if userFromDB.PassHash != user.PassHash {
		err = errors.New("incorrect password")
		return
	}
	return
}
