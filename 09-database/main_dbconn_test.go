package main

import (
	"testing"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	_ "github.com/go-sql-driver/mysql"
	"github.com/DATA-DOG/go-sqlmock"
)

// TestdbConn function would test dbConn function
func TestdbConn(t *testing.T) {

	// a test case for dbConn function
	tests := []struct {
		name    string
		dbDriver string
		dbUser string
		dbPass string
		dbName string
		wantErr bool
	}{
		{
			name:    "Test scenario where all database parameters are correct",
			dbDriver: "mysql",
			dbUser: "root",
			dbPass: "12345",
			dbName: "gotest",
			wantErr: false,
		},
		{
			name:    "Test scenario where database parameters are incorrect",
			dbDriver: "mysql",
			dbUser: "wrongUser",
			dbPass: "wrongPassword",
			dbName: "wrongDB",
			wantErr: true,
		},
	}

	// run the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			
			// Create sqlmock database connection
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			// Before we actually call the function, we anticipate that these queries are going to be called
			mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}))

			// call the function
			dbConn()

			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "12345"
	dbName := "gotest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
