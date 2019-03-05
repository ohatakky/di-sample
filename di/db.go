package di

import "database/sql"

func InjectDB() *sql.DB {
	db, err := sql.Open("db", "dsn")
	if err != nil {
		panic(err)
	}
	return db
}
