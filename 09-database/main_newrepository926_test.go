package main

import (
	"testing"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name    string
		mockDB  func() (*sql.DB, sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				return db, mock
			},
			wantErr: false,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			wantErr: true,
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnError(sql.ErrConnClosed)
				return db, mock
			},
			wantErr: true,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectQuery("^SELECT (.+) FROM users$").WillReturnError(sql.ErrTxDone)
				return db, mock
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock := tc.mockDB()
			defer db.Close()

			repo := NewRepository(db)

			if tc.wantErr {
				if err := mock.ExpectationsWereMet(); err == nil {
					t.Errorf("Expected an error but didn't get one")
				}
			} else {
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("There was an error '%s' when there shouldn't be", err)
				}
				if repo == nil {
					t.Errorf("Expected a repository but got nil")
				}
			}
		})
	}
}
