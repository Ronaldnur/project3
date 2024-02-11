package database

import (
	"database/sql"
	"fmt"
	"log"
	"project3/infra/config"

	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	appConfig := config.GetAppConfig()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBPassword, appConfig.DBName,
	)

	db, err = sql.Open(appConfig.DBDialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	userTable := `
	CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		full_name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		role VARCHAR(255) NOT NULL,
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now()
	);
	`
	categoryTable := `
	CREATE TABLE IF NOT EXISTS "category" (
		id SERIAL PRIMARY KEY,
		type VARCHAR(255) NOT NULL,
		created_at timestamptz DEFAULT now(),
		updated_at timestamptz DEFAULT now()
	);
	`
	taskTable := `
	CREATE TABLE IF NOT EXISTS "task" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			status BOOLEAN NOT NULL,
			user_id INT NOT NULL,
			category_id INT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT task_user_id_fk
				FOREIGN KEY(user_id)
					REFERENCES "user"(id)
					ON DELETE CASCADE,
			CONSTRAINT task_category_id_fk
				FOREIGN KEY(category_id)
					REFERENCES "category"(id)
					ON DELETE CASCADE
		);
		`

	createTableQueries := fmt.Sprintf("%s %s %s", userTable, categoryTable, taskTable)

	_, err := db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}
