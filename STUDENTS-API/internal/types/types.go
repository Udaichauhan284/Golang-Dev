package types

//the validator which i have install, using here in struct
type Student struct {
	Id int 
	Name string `validate:"required"`
	Email string `validate:"required"`
	Age int `validate:"required"`
}