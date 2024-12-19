package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {
	var testCases = []struct {
		name      string
		mockDB    func() *sql.DB
		wantError bool
	}{
		{
			name: "Successful creation of a new repository with an active and connected database",
			mockDB: func() *sql.DB {
				db, _, _ := sqlmock.New()
				return db
			},
			wantError: false,
		},
		{
			name: "Failed creation of a new repository with an inactive database",
			mockDB: func() *sql.DB {
				return nil
			},
			wantError: true,
		},
		{
			name: "Failed creation of a new repository with a disconnected database",
			mockDB: func() *sql.DB {
				db, mock, _ := sqlmock.New()
				mock.ExpectClose()
				db.Close()
				return db
			},
			wantError: true,
		},
		{
			name: "Handling database errors during repository creation",
			mockDB: func() *sql.DB {
				db, mock, _ := sqlmock.New()
				mock.ExpectQuery("SELECT 1").WillReturnError(sql.ErrConnDone)
				return db
			},
			wantError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := tc.mockDB()
			repo := NewRepository(db)
			if tc.wantError {
				if repo != nil {
					t.Errorf("Expected error but received none")
				} else {
					t.Log("Test passed: received expected error")
				}
			} else {
				if repo == nil {
					t.Errorf("Expected repository but received nil")
				} else {
					t.Log("Test passed: received expected repository")
				}
			}
		})
	}
}
