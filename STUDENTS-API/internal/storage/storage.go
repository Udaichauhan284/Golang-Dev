package storage

import (
	"github.com/Udaichauhan284/Golang-Dev/internal/types"
)

//this is Storage interface which have, createstudent function, which return the id and error
type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)

	//creating same for getting student by id
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
	UpdateStudent(id int64, name string, email string, age int) error 
	DeleteStudent(id int64) error
}