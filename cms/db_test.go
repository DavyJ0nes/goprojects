package cms

import (
	"reflect"
	"strconv"
	"testing"
)

var p *Page

func Test_CreatePage(t *testing.T) {
	p = &Page{
		Title:   "test",
		Content: "test",
	}

	id, err := CreatePage(p)
	if err != nil {
		t.Error("Failed to create page: %s\n", p)
	}
	p.ID = id
}

func Test_GetPage(t *testing.T) {
	page, err := GetPage(strconv.Itoa(p.ID))
	if err != nil {
		t.Error("failed to get page: %s\n", err.Error())
	}

	if page.ID != p.ID {
		t.Error("page IDs do not match.\nGot:      %d\nExpected: %d\n", page.ID, p.ID)
	}

	if reflect.DeepEqual(page, p) != true {
		t.Error("Pages do not match.\nGot:       %+v\nExpected: %+v\n", page.ID, p.ID)
	}

}
