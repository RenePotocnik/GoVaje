package MariaDB

/*
Definicije za migracije baze.
Kadar želimo nekaj spremenit dodamo novi objekt v array z višjo verzijo
*/
var definitions = []migrationDefinition{
	{
		Version:     1,
		Description: "Initial database creation",
		Script:      createTask,
	},
}

// Initial database creation
// var createDB = `
// CREATE TABLE IF NOT EXISTS user(
//     user_id INT NOT NULL AUTO_INCREMENT,
//     username TEXT NOT NULL,
//     pass_hash TEXT NOT NULL,
//     email TEXT NOT NULL,
//     PRIMARY KEY (user_id)
// )
// CHARACTER SET 'utf8mb4',
// COLLATE 'utf8mb4_unicode_ci'
// `

// Task table creation
var createTask = `
CREATE TABLE IF NOT EXISTS task(
    id INT NOT NULL AUTO_INCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    date_added DATETIME NOT NULL,
    predicted_date DATETIME NOT NULL,
    PRIMARY KEY (id)
)
CHARACTER SET 'utf8mb4',
COLLATE 'utf8mb4_unicode_ci'
    `
