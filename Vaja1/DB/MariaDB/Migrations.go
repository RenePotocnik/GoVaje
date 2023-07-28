package MariaDB

import (
	"database/sql"
	"vaja1/Logic"
)

/*
Definition for database migration
Whenever we want to change something we add a new object to the array with a higher version
*/
var definitions = []migrationDefinition{
	{
		Version:     1,
		Description: "Initial database creation",
		Script:      createTask,
	},
	{
		Version:        2,
		Description:    "Add user table",
		Script:         createUser,
		PostScriptFunc: addUser,
	},
}

func addUser(db *sql.DB) (err error) {
	username := "demo"
	password := "demo"

	// Encrypt the password
	pass_hash, err := Logic.EncryptPassword(password)
	if err != nil {
		return err
	}

	// Insert the user into the database
	_, err = db.Exec("INSERT INTO user (username, pass_hash) VALUES (?, ?)", username, pass_hash)
	if err != nil {
		return err
	}

	// Get the demo user's id
	rows, err := db.Query("SELECT user_id FROM user WHERE username = ?", username)
	if err != nil {
		return err
	}
	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			return
		}
	}()
	var userId int
	if rows.Next() {
		err = rows.Scan(&userId)
		if err != nil {
			return err
		}
	}

	// Update the task table to reference the demo user
	_, err = db.Exec("UPDATE task SET user_id = ?", userId)
	if err != nil {
		return err
	}

	return nil
}

// Task table creation
// Foreign key added to `user_id` to reference to `user[user_id]`
var createTask = `
CREATE TABLE IF NOT EXISTS task(
    id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, 
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date_added TEXT NOT NULL,
    predicted_date TEXT NOT NULL,
	user_id INT NOT NULL
)
CHARACTER SET 'utf8mb4',
COLLATE 'utf8mb4_unicode_ci'
`

// User table creation
var createUser = `
CREATE TABLE IF NOT EXISTS user(
    user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    username TEXT NOT NULL,
    pass_hash TEXT NOT NULL
)
CHARACTER SET 'utf8mb4',
COLLATE 'utf8mb4_unicode_ci';

ALTER TABLE task
ADD FOREIGN KEY (user_id) REFERENCES user(user_id);
`
