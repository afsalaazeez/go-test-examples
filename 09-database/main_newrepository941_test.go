package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name        string
		mockDBFunc  func() (*sql.DB, sqlmock.Sqlmock)
		expectError bool
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDBFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnResult(sqlmock.NewResult(1, nil))
				return db, mock
			},
			expectError: false,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDBFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			expectError: true,
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockDBFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			expectError: true,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDBFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrTxDone)
				return db, mock
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock := tc.mockDBFunc()
			defer db.Close()

			repo := NewRepository(db)

			if (repo == nil) != tc.expectError {
				t.Errorf("Expected error status: %v, but got: %v", tc.expectError, repo == nil)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
