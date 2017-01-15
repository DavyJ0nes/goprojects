package cms

import (
	"database/sql"
	"fmt"
	// Postgres Driver
	_ "github.com/lib/pq"
)

var (
	store = newDB()
)

// PgStore is an Exported db instance
type PgStore struct {
	DB *sql.DB
}

func newDB() *PgStore {
	PG_HOST := "docker"
	PG_GP := "goprojects"
	connString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", PG_GP, PG_GP, PG_HOST, PG_GP)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		panic(err)
	}

	dbCheck(db)

	return &PgStore{
		DB: db,
	}
}

func dbCheck(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		panic(err)
	}
}

// CreatePage adds new page to DB
func CreatePage(p *Page) (int, error) {
	var id int
	err := store.DB.QueryRow("INSERT INTO pages(title, content) VALUES($1, $2) RETURNING id", p.Title, p.Content).Scan(&id)
	return id, err
}

// GetPage returns the content of a Page from DB
func GetPage(id string) (*Page, error) {
	var p Page
	err := store.DB.QueryRow("SELECT * FROM pages WHERE id = $1", id).Scan(&p.ID, &p.Title, &p.Content)
	return &p, err
}

// GetPages returns All pages from DB
func GetPages() ([]*Page, error) {
	rows, err := store.DB.Query("SELECT * FROM pages")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pages := []*Page{}
	for rows.Next() {
		var p Page
		err = rows.Scan(&p.ID, &p.Title, &p.Content)
		if err != nil {
			return nil, err
		}
		pages = append(pages, &p)
	}
	return pages, nil
}
