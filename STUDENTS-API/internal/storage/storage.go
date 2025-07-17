package storage

//this is Storage interface which have, createstudent function, which return the id and error
type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
}