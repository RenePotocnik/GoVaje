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
	_, err = db.Exec("INSERT INTO user (username, pass_hash) VALUES (?, ?)", username, pass_hash)
	if err != nil {
		return err
	}

	// Change the `user_id` of all the tasks to the ID of the newly created user
	demo_user_id, err := db.Query("SELECT user_id FROM user WHERE username = ? AND pass_hash = ?", username, pass_hash)
	_, err = db.Exec("UPDATE task SET user_id = ?", demo_user_id)

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
    predicted_date TEXT NOT NULL
	user_id INT NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user(user_id)
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
COLLATE 'utf8mb4_unicode_ci'
`
