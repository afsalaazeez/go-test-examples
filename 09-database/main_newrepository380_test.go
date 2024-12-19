package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name           string
		dbActive       bool
		dbConnected    bool
		expectedResult bool
	}{
		{
			name:           "Successful Repository Creation with an Active and Connected Database",
			dbActive:       true,
			dbConnected:    true,
			expectedResult: true,
		},
		{
			name:           "Failed Repository Creation with an Inactive Database",
			dbActive:       false,
			dbConnected:    true,
			expectedResult: false,
		},
		{
			name:           "Failed Repository Creation with a Disconnected Database",
			dbActive:       true,
			dbConnected:    false,
			expectedResult: false,
		},
		{
			name:           "Handling Database Errors During Repository Creation",
			dbActive:       false,
			dbConnected:    false,
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			if !tc.dbActive {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			}

			if !tc.dbConnected {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			}

			repo := NewRepository(db)

			if tc.expectedResult && repo == nil {
				t.Errorf("Expected a repository, but got nil")
			} else if !tc.expectedResult && repo != nil {
				t.Errorf("Expected nil, but got a repository")
			}
		})
	}
}
