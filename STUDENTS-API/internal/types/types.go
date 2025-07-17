package types

//the validator which i have install, using here in struct
type Student struct {
	Id int64 //making this bigger, because it save in database, by storage.go
	Name string `validate:"required"`
	Email string `validate:"required"`
	Age int `validate:"required"`
}