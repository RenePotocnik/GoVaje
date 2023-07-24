package MariaDB

import (
	"database/sql"
	"errors"
	"fmt"
	"sort"
)

type migrator struct {
	database    *sql.DB
	definitions []migrationDefinition
}

type migrationDefinition struct {
	Version        float64
	Description    string
	Script         string
	PreScriptFunc  migratorFunc
	PostScriptFunc migratorFunc
}

type migratorFunc func(*sql.DB) error

func (m *migrator) Migrate() (err error) {

	if len(m.definitions) == 0 {
		return errors.New("no definitions for migrator")
	}

	sort.Slice(m.definitions, func(i, j int) bool {
		return m.definitions[i].Version < m.definitions[j].Version
	})

	_, err = m.database.Exec(createVersionsTable)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	version, err := m.getVersion()

	for i := range m.definitions {

		if m.definitions[i].Version > version {

			if m.definitions[i].PreScriptFunc != nil {
				err = m.definitions[i].PreScriptFunc(m.database)
				if err != nil {
					err = errors.New("error pre migrating database at: " + m.definitions[i].Description + " - Error: " + err.Error())
					fmt.Println(err.Error())
					return err
				}

			}

			_, err = m.database.Exec(m.definitions[i].Script)
			if err != nil {
				err = errors.New("error migrating database at: " + m.definitions[i].Description + " - Error: " + err.Error())
				fmt.Println(err.Error())
				return err
			}

			if m.definitions[i].PostScriptFunc != nil {
				err = m.definitions[i].PostScriptFunc(m.database)
				if err != nil {
					err = errors.New("error post migrating database at: " + m.definitions[i].Description + " - Error: " + err.Error())
					fmt.Println(err.Error())
					return err
				}

			}

			err = m.setVersion(m.definitions[i].Version)
			if err != nil {
				err = errors.New("error setting database version at: " + m.definitions[i].Description + " - Error: " + err.Error())
				fmt.Println(err.Error())
				return err
			}
		}

	}

	return nil

}

func (m *migrator) getVersion() (version float64, err error) {

	statement, err := m.database.Prepare("SELECT version FROM db_version ORDER BY version DESC LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := statement.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	rows, err := statement.Query()
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

	if !rows.Next() {
		version = 0
		return
	}

	err = rows.Scan(&version)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	return

}

func (m *migrator) setVersion(version float64) (err error) {

	statement, err := m.database.Prepare("INSERT INTO db_version (version) VALUES (?)")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer func() {
		err2 := statement.Close()
		if err2 != nil {
			fmt.Println(err2.Error())
		}
	}()

	_, err = statement.Exec(version)
	if err != nil {

		fmt.Println(err.Error())

	}

	return

}

var createVersionsTable = `
CREATE TABLE IF NOT EXISTS db_version(
    version_id INT NOT NULL AUTO_INCREMENT,
    version FLOAT NOT NULL,
    PRIMARY KEY (version_id)
);
`
