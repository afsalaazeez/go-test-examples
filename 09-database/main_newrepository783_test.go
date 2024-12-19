package main

import (
	"database/sql"
	"testing"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name           string
		mockDBBehavior func(mock sqlmock.Sqlmock)
		wantErr        bool
	}{
		{
			name: "Repository creation with valid database connection",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
		{
			name: "Repository creation with inactive database connection",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
		{
			name: "Repository creation with a disconnected database",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
		{
			name: "Repository creation with a database that returns errors",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(errors.New("generic error"))
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tc.mockDBBehavior(mock)

			repo := NewRepository(db)

			if tc.wantErr && repo == nil {
				t.Logf("expected no repository, got no repository")
				return
			}

			if !tc.wantErr && repo != nil {
				t.Logf("expected a repository, got a repository")
				return
			}

			t.Errorf("expected '%t' but got '%v'", tc.wantErr, repo == nil)
		})
	}
}
