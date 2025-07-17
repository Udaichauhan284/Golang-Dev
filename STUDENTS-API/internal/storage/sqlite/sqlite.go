package sqlite

import (
	"database/sql"

	"github.com/Udaichauhan284/Golang-Dev/internal/config"
	_ "github.com/mattn/go-sqlite3" //this is in work behind the secene so thatswhy i have put the _, it is not working like directly inUse, like database/sql
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error){
	db, err := sql.Open("sqlite3", cfg.StoragePath);
	if err != nil {
		return nil, err;
	}

	//table create
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTERGER
	)`)

	if err != nil {
		return nil, err;
	}

	//this is how i can return the multiple value 
	return &Sqlite{
		Db : db,
	}, nil
}