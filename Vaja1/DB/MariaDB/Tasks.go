package MariaDB

import (
	"errors"
	"fmt"
	"vaja1/DataStructures"
)

func (db *MariaDB) GetTasks(userID int) (tasks []DataStructures.Task, err error) {
	// For statements that do not return a result (such as INSERT, UPDATE, and DELETE), use `Exec` instead of `Query`.
	rows, err := db.database.Query("SELECT id, title, description, date_added, predicted_date from task WHERE user_id = ?", userID)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Close the rows when the function ends
	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	// Create task object and an undefined array of tasks
	var task DataStructures.Task
	tasks = make([]DataStructures.Task, 0)

	// Loop through all rows
	for rows.Next() {
		// Read the data from the row into the task object
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.DateAdded, &task.PredictedDate)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Add the task user to the tasks array
		tasks = append(tasks, task)
	}
	return
}

func (db *MariaDB) GetTaskById(taskId, userId int) (task DataStructures.Task, err error) {

	rows, err := db.database.Query("SELECT id, title, description, date_added, predicted_date from task  WHERE id = ? AND user_id = ? LIMIT 1", taskId, userId)
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

	if rows.Next() {

		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.DateAdded, &task.PredictedDate)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	} else {
		err = errors.New("task not found")
		return
	}

	return
}

func (db *MariaDB) PutTaskById(taskId int, task DataStructures.Task) (err error) {
	_, err = db.database.Exec("UPDATE task SET title = ?, description = ?, predicted_date = ?, user_id = ? WHERE id = ?", task.Title, task.Description, task.PredictedDate, task.UserId, taskId)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

// CreateTask creates a new task in the database
func (db *MariaDB) CreateTask(task DataStructures.Task, userId int) (err error) {
	_, err = db.database.Exec("INSERT INTO task (title, description, date_added, predicted_date, user_id) VALUES (?, ?, ?, ?, ?)", task.Title, task.Description, task.DateAdded, task.PredictedDate, userId)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}

// DeleteTask deletes a task from the database given the task ID
func (db *MariaDB) DeleteTask(taskId int) (err error) {
	_, err = db.database.Exec("DELETE FROM task WHERE id = ?", taskId)
	if err != nil {
		fmt.Println(err.Error())
	}
	return
}
