package main

import (
	"database/sql"
	"errors"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {

	testCases := []struct {
		name           string
		mockDB         func() (*sql.DB, sqlmock.Sqlmock, error)
		expectedResult bool
	}{
		{
			name: "Successful Repository Creation with Active and Connected Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				if err != nil {
					return nil, nil, err
				}
				mock.ExpectPing().WillReturnResult(sqlmock.NewResult(1, nil))
				return db, mock, nil
			},
			expectedResult: true,
		},
		{
			name: "Failed Repository Creation with Inactive Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, mock, err := sqlmock.New()
				if err != nil {
					return nil, nil, err
				}
				mock.ExpectPing().WillReturnError(errors.New("database inactive"))
				return db, mock, nil
			},
			expectedResult: false,
		},
		{
			name: "Failed Repository Creation with Disconnected Database",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, _, err := sqlmock.New()
				if err != nil {
					return nil, nil, err
				}
				db.Close()
				return db, nil, nil
			},
			expectedResult: false,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDB: func() (*sql.DB, sqlmock.Sqlmock, error) {
				db, _, err := sqlmock.New()
				if err != nil {
					return nil, nil, err
				}
				db.Close()
				return db, nil, nil
			},
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			db, _, err := tc.mockDB()
			defer db.Close()
			if err != nil {
				t.Fatalf("failed to mock db connection: %v", err)
			}

			repo := NewRepository(db)

			if tc.expectedResult {
				assert.NotNil(t, repo, "Expected repository to be not nil")
			} else {
				assert.Nil(t, repo, "Expected repository to be nil")
			}
		})
	}
}
