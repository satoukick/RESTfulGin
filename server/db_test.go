package server

import (
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
}
