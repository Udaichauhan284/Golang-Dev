package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/Udaichauhan284/Golang-Dev/internal/config"
	"github.com/Udaichauhan284/Golang-Dev/internal/types"
	_ "github.com/mattn/go-sqlite3" //this is in work behind the secene so thatswhy i have put the _, it is not working like directly inUse, like database/sql
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	//table create
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTERGER
	)`)

	if err != nil {
		return nil, err
	}

	//this is how i can return the multiple value
	return &Sqlite{
		Db: db,
	}, nil
}

// now i want to implement CreateStudent, to implements here implicity
// now this method attach to this Sqlite struct
func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	//now want to create record in database

	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	//now checking how many rows were inserted
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Rows Affected: ", rowsAffected)

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

//these question ?, help me to prevent the SQL injection attack on website, Values will be prepare later

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT * FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	//that data came from stmt, need to serlialize into struct
	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id, &student.Name, &student.Email, &student.Age)
	if err != nil {
		//if user not found
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("no student found with id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("query error: %w", err)
	}

	return student, nil
}
