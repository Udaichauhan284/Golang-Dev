package types

//the validator which i have install, using here in struct
type Student struct {
	Id int64 `json:"id"` //making this bigger, because it save in database, by storage.go
	Name string `json:"name" validate:"required"`
	Email string `josn:"email" validate:"required"`
	Age int `json:"age" validate:"required"`
}