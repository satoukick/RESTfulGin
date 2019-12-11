package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	logs "github.com/satoukick/webserver/log"
	"github.com/stretchr/testify/assert"
)

func assertJSON(actual []byte, expectedData interface{}, t *testing.T) {
	expected, err := json.Marshal(expectedData)
	if err != nil {
		t.Fatalf("json marshal error %s", err)
	}

	assert.Equal(t, expected, actual, "should be equal")
}

func TestMockFetchAll(t *testing.T) {
	_, mock, err := sqlmock.NewWithDSN("mock_0")
	if err != nil {
		logs.Fatal("Test Mock panic", err)
	}

	db, err = gorm.Open("sqlmock", "mock_0")
	if err != nil {
		logs.Fatal("Test Mock gorm open panic", err)
	}

	defer db.Close()

	r := setupRouter()
	ht := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/api/v1/todos/", nil)

	rows := sqlmock.NewRows([]string{"id", "title", "completed"}).
		AddRow(1, "record 1", 1).
		AddRow(2, "record 2", 1)

	regEx := "^SELECT (.+) FROM \"todo_models\" (.+)"
	mock.ExpectQuery(regEx).WillReturnRows(rows)

	r.ServeHTTP(ht, request)
	actual := ht.Body.Bytes()

	if ht.Code != http.StatusOK {
		t.Fatalf("wrong code for http request expected %d, actual %d", http.StatusOK, ht.Code)
	}

	expected := struct {
		Data   []transformedTodo `json:"data"`
		Status int               `json:"status"`
	}{
		Status: http.StatusOK,
		Data: []transformedTodo{
			{ID: 1, Title: "record 1", Completed: true},
			{ID: 2, Title: "record 2", Completed: true},
		},
	}
	assertJSON(actual, expected, t)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Expectations were not met %s", err)
	}
}
