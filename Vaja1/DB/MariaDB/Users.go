package MariaDB

import (
	"errors"
	"fmt"
	"vaja1/DataStructures"
)

// GetUsers function from the DB interface
func (db *MariaDB) GetUsers() (users []DataStructures.User, err error) {
	rows, err := db.database.Query("SELECT user_id, username, email from user")
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

	var user DataStructures.User
	users = make([]DataStructures.User, 0)

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		users = append(users, user)
	}
	return
}

func (db *MariaDB) GetUserById(userId int) (user DataStructures.User, err error) {
	rows, err := db.database.Query("SELECT user_id, username, email from user  WHERE user_id = ? LIMIT 1", userId)
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

	// Move to the next row (the first returned row), if ot does not exist, it means that the response is empty and does not exist
	if !rows.Next() {
		err = errors.New("user does not exist")
		return
	}

	err = rows.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return
}
