package main

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
)

// Mocked repository for testing
type repository struct {
	db *sql.DB
}

// Repository interface
type Repository interface{}

// NewRepository creates a new repository
func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func TestNewRepository(t *testing.T) {
	// create sql.DB and mock to be used
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	// we expect our function to call db.Ping on the database
	mock.ExpectPing()

	// call our function with the mock database
	repo := NewRepository(db)

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NotNil(t, repo, "Expected repository not to be nil")
}
