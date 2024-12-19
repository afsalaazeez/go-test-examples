package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name           string
		mockDB         func() (*sql.DB, sqlmock.Sqlmock)
		expectedOutput Repository
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				return db, mock
			},
			expectedOutput: &repository{},
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				db.Close()
				return db, mock
			},
			expectedOutput: nil,
		},
		{
			name:           "Failed Repository Creation with a Disconnected Database",
			mockDB:         nil,
			expectedOutput: nil,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnError(fmt.Errorf("some error"))
				return db, mock
			},
			expectedOutput: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			if tc.mockDB != nil {
				db, mock := tc.mockDB()
				defer db.Close()

				dbConn = func() *sql.DB {
					return db
				}

				mock.ExpectClose()
			}

			got := NewRepository(nil)

			if tc.expectedOutput == nil {
				assert.Nil(t, got)
			} else {
				assert.NotNil(t, got)
			}
		})
	}
}
