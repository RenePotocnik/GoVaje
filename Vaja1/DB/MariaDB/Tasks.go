package MariaDB

import (
	"errors"
	"fmt"
	"vaja1/DataStructures"
)

func (db *MariaDB) GetTasks() (tasks []DataStructures.Task, err error) {

	// Naredimo query na bazo
	// Za stavke, od katerih ne pričakujemo odgovora (UPDATE, INSERT) uporabimo namesto "Query", "Exec"
	rows, err := db.database.Query("SELECT id, title, description, date_added, predicted_date from task")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Zapremo vrstice, ko je funkcija končana
	defer func() {
		err2 := rows.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	// Ustvarimo objekt User in neomejen array tipa User
	var task DataStructures.Task
	tasks = make([]DataStructures.Task, 0)

	// Loop čez vse vrstice
	for rows.Next() {

		// Preberemo podatke vrstice v objekt
		err = rows.Scan(&task.Id, &task.Title, &task.Description, &task.DateAdded, &task.PredictedDate)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Dodamo objekt user na konec arraya users
		tasks = append(tasks, task)

	}
	// Vrnemo podatke. Ne rabimo napisat katero spremenljivko vrnemo saj je ta definirana in inicializirana na vrhu funkcije
	return

}

func (db *MariaDB) GetTaskById(taskId int) (task DataStructures.Task, err error) {

	rows, err := db.database.Query("SELECT id, title, description, date_added, predicted_date from task  WHERE id = ? LIMIT 1", taskId)
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
