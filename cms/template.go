// template is the home for all Structs used by the cms package
package cms

import (
	"html/template"
	"time"
)

// Tmpl parses templates
var Tmpl = template.Must(template.ParseGlob("templates/*"))

// Page is top level object
type Page struct {
	ID      int
	Title   string
	Content string
	Posts   []*Post
}

// Post is for user posts
type Post struct {
	ID            int
	Title         string
	Content       string
	DatePublished time.Time
	Comments      []*Comment
}

// Comment is for user comments
type Comment struct {
	ID            int
	Author        string
	Content       string
	DatePublished time.Time
}
