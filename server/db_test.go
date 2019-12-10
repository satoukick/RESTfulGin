package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	logs "github.com/satoukick/webserver/log"
)

func TestMock(t *testing.T) {
	_, mock, err := sqlmock.NewWithDSN("mock_0")
	if err != nil {
		logs.Fatal("Test Mock panic", err)
	}

	db, err := gorm.Open("sqlmock", "mock_0")
	if err != nil {
		logs.Fatal("Test Mock gorm open panic", err)
	}

	defer db.Close()

	r := setupRouter()
	ht := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/todos/", nil)
	r.ServeHTTP(ht, request)
	actual := ht.Body.Bytes()

	rows := sqlmock.NewRows([]string{"id", "title", "completed"}).
		AddRow(1, "record 1", true).
		AddRow(2, "record 2", false)

	mock.ExpectQuery("^SELECT (.+) FROM todo_models$")

}
