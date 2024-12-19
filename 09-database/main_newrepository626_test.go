package main

import (
	"database/sql"
	"testing"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {

	scenarios := []struct {
		name        string
		mockDbFunc  func() (*sql.DB, sqlmock.Sqlmock)
		wantRepoNil bool
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDbFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnResult(sqlmock.NewResult(1, nil))
				return db, mock
			},
			wantRepoNil: false,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDbFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			wantRepoNil: true,
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockDbFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			wantRepoNil: true,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDbFunc: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrTxDone)
				return db, mock
			},
			wantRepoNil: true,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {
			db, _ := s.mockDbFunc()
			gotRepo := NewRepository(db)
			if (gotRepo == nil) != s.wantRepoNil {
				t.Errorf("NewRepository() = %v, want %v", gotRepo, s.wantRepoNil)
			}
		})
	}
}
