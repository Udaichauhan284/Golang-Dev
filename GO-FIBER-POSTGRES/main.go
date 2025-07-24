package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/udaichauhan/go-fiber-postgres/models"
	"github.com/udaichauhan/go-fiber-postgres/storage"
	"gorm.io/gorm"
);

type Book struct{
	Author string `json:"author"`
	Title string	`json:"title"`
	Publisher string	`json:"publisher"`
}
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	if err := context.BodyParser(&book); err != nil {
		context.Status(http.StatusUnprocessableEntity)
		return context.JSON(&fiber.Map{"message": "request failed"})
	}

	//save to database
	if err := r.DB.Create(&book).Error; err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message" : "could not create book"})
		return err;
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "book has been added",
	})
	return nil;
	
}

func (r *Repository) GetBooks(context *fiber.Ctx) error {
	//create a slice to hold the retrieved book records from the database
	books := &[]models.Books{}

	///Query all record from "books" table and store them in the "books" slice
	err := r.DB.Find(books).Error;

	//return the error if any error you face
	if err != nil {
		context.Status((http.StatusBadRequest)).JSON(&fiber.Map{
			"message" : "could not get books",
		});
		return err;
	}

	//if successfully, return the list of books with a 200 status code
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "books fetch successfully",
		"data" : books,
	});
	return nil;
}

func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	book := models.Books{}
	id := context.Params("id")
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message" : "id cannot be empty",
		})
		return nil;
	}

	err := r.DB.Delete(book, id);
	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message" : "could not delete book",
		})
		return err.Error;
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "books delete successfully",
	})
	return nil;
}

func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id");
	books := &models.Books{}
	if id == ""{
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message" : "id cannot be empty",
		})
		return nil;
	}
	fmt.Println("the ID is: ", id);

	err := r.DB.Where("id = ?", id).First(books).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message" : "could not get the book",
		})
		return err;
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message" : "book id fetched successfully",
		"data" : books,
	});
	return nil;
}

func(r *Repository) SetupRoutes(app *fiber.App){
	api := app.Group("/api");
	api.Post("/create_books", r.CreateBook);
	api.Delete("delete_book/:id", r.DeleteBook);
	api.Get("/get_books/:id", r.GetBookByID);
	api.Get("/books", r.GetBooks);
}

func main(){
	err := godotenv.Load(".env");
	if err != nil {
		log.Fatal(err);
	}

	//now get the config from storage.go
	config := &storage.Config{
		Host: os.Getenv("DB_HOST"),
		Port: os.Getenv("DB_PORT"),
		Password: os.Getenv("DB"),
		User : os.Getenv("DB_USER"),
		SSLMode: os.Getenv("DB_SSLMODE"),
		DbName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config);
	if err != nil {
		log.Fatal("could not load the database");
	}

	migrationErr := models.MigrateBooks(db);
	if migrationErr != nil {
		log.Fatal("could not migrate db");
	}

	r := Repository{
		DB : db,
	}

	app := fiber.New()
	r.SetupRoutes(app);
	app.Listen(":8080");
}