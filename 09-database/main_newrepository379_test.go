package main

import (
	"database/sql"
	"testing"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestNewRepository(t *testing.T) {

	scenarios := []struct {
		name           string
		mockDBBehavior func(mock sqlmock.Sqlmock)
		expectedResult bool
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnResult(sqlmock.NewResult(1, 1))
			},
			expectedResult: true,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(sql.ErrNoRows)
			},
			expectedResult: false,
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			},
			expectedResult: false,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDBBehavior: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(sql.ErrTxDone)
			},
			expectedResult: false,
		},
	}

	for _, s := range scenarios {
		t.Run(s.name, func(t *testing.T) {

			db, mock, _ := sqlmock.New()
			defer db.Close()

			s.mockDBBehavior(mock)

			repo := NewRepository(db)

			if s.expectedResult {
				assert.NotNil(t, repo)
			} else {
				assert.Nil(t, repo)
			}
		})
	}
}
