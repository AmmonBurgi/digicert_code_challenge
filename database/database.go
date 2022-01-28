package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "mysecretpassword"
	dbname   = "postgres"
)

type Book struct {
	Id     int
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}

var DB *sql.DB
var err error

func InitDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = DB.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected.")

	_, _ = DB.Exec(`CREATE TABLE books (
		id SERIAL PRIMARY KEY,
		isbn varchar(255) NOT NULL,
		title varchar(255) NOT NULL,
		author varchar(255) NOT NULL,
		price decimal(5,2) NOT NULL
	);`)
}

func GetAllBooks() ([]Book, error) {
	var books []Book

	rows, err := DB.Query("select * from books")

	if err != nil {
		return books, err
	}

	defer rows.Close()

	for rows.Next() {
		var book Book

		if err := rows.Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Price); err != nil {
			return books, err
		}

		books = append(books, book)
	}

	return books, nil
}

func CreateBook(book Book) error {
	statement := fmt.Sprintf(`INSERT INTO books (isbn, title, author, price) VALUES ('%s', '%s', '%s', %f)`, book.Isbn, book.Title, book.Author, book.Price)
	_, err = DB.Exec(statement)
	return err
}

func DeleteBookById(i int) error {
	statement := fmt.Sprintf(`DELETE FROM books WHERE id = %d`, i)

	_, err = DB.Exec(statement)
	return err
}

func GetBookById(i int) (Book, error) {
	var book Book

	if err := DB.QueryRow("SELECT * FROM books WHERE id = $1", i).Scan(&book.Id, &book.Isbn, &book.Title, &book.Author, &book.Price); err != nil {
		return book, err
	}

	return book, nil
}

func UpdateBookById(i int, book Book) error {
	statement := fmt.Sprintf(`UPDATE books SET isbn = '%s', title = '%s', author = '%s', price = %f WHERE id = %d`, book.Isbn, book.Title, book.Author, book.Price, i)
	_, err = DB.Exec(statement)
	return err
}
