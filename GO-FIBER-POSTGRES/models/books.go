package models

import "gorm.io/gorm" // Importing GORM ORM package to interact with the database


// Books struct defines the schema/model for the 'books' table in the database
type Books struct{
	ID uint `gorm:"primary key;autoIncrement" json:"id"`
	Author *string `json:"author"` //using *string allows the fields to be NULL in the DB
	Title *string `json:"title"`
	Publisher *string `json:"publisher"`
}

// MigrateBooks runs the auto-migration for the Books model
func MigrateBooks(db *gorm.DB) error {
	// Automatically creates or updates the 'books' table to match the Books struct
	err := db.AutoMigrate(&Books{})
	
	// Return any error that occurs during the migration
	return err
}
//AutoMigrate() check if the table exists
// if not,, it creates it
// if yes, it adjust the schema (adds missing columns, etc)