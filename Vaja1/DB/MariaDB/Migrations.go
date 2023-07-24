package MariaDB

/*
Definicije za migracije baze.
Kadar želimo nekaj spremenit dodamo novi objekt v array z višjo verzijo
*/
var definitions = []migrationDefinition{
	{
		Version:     1,
		Description: "Initial database creation",
		Script:      createDB,
	},
	{
		Version:     1.1,
		Description: "Add some field",
		Script:      "UPDATE TABLE users ADD phone_number TEXT NOT NULL;",
	},
}

// Začetna verzija baze
var createDB = `

CREATE TABLE IF NOT EXISTS user(
    user_id INT NOT NULL AUTO_INCREMENT,
    username TEXT NOT NULL,
    pass_hash TEXT NOT NULL,
    email TEXT NOT NULL,
    PRIMARY KEY (user_id)
)
CHARACTER SET 'utf8mb4',
COLLATE 'utf8mb4_unicode_ci'
`
