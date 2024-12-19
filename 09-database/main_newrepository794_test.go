package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	testCases := []struct {
		name          string
		mockDBFunc    func()
		expectedError error
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDBFunc: func() {

			},
			expectedError: nil,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDBFunc: func() {

			},
			expectedError: errors.New("database is inactive"),
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockDBFunc: func() {

			},
			expectedError: errors.New("database is disconnected"),
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDBFunc: func() {

			},
			expectedError: errors.New("database error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			tc.mockDBFunc()

			repo := NewRepository(db)

			if tc.expectedError != nil {
				assert.Nil(t, repo)
			} else {
				assert.NotNil(t, repo)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
