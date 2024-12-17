package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)




func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}
func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name          string
		db            *sql.DB
		expectedError error
	}{
		{
			name:          "TestFindNonExisting",
			db:            nil,
			expectedError: sql.ErrNoRows,
		},
		{
			name:          "TestFindWithInvalidID",
			db:            nil,
			expectedError: sql.ErrNoRows,
		},
		{
			name:          "TestAddWithInvalidDate",
			db:            nil,
			expectedError: sql.ErrNoRows,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			mock.ExpectQuery("^SELECT (.+) FROM tasks WHERE id=?").
				WithArgs(tc.db).
				WillReturnError(tc.expectedError)

			repo := NewRepository(db)
			task, err := repo.Find(tc.db)

			t.Logf("Running scenario: %s", tc.name)

			if err != nil {

				if tc.expectedError != nil {

					if assert.Error(t, err) {
						assert.Equal(t, tc.expectedError, err)
					}
				} else {
					t.Errorf("Error was not expected: %s", err)
				}
			} else {

				if tc.expectedError != nil {
					t.Errorf("Error was expected: %s", tc.expectedError)
				}
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("There were unfulfilled expectations: %s", err)
			}
		})
	}
}
