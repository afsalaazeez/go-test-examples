package main

import (
	"database/sql"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	tests := []struct {
		name             string
		mockDbConnection func() (*sql.DB, sqlmock.Sqlmock)
		expectedRepo     bool
	}{
		{
			name: "Successful Repository Creation with an Active and Connected Database",
			mockDbConnection: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnResult(sqlmock.NewResult(1, nil))
				return db, mock
			},
			expectedRepo: true,
		},
		{
			name: "Failed Repository Creation with an Inactive Database",
			mockDbConnection: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			expectedRepo: false,
		},
		{
			name: "Failed Repository Creation with a Disconnected Database",
			mockDbConnection: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
				return db, mock
			},
			expectedRepo: false,
		},
		{
			name: "Handling Database Errors During Repository Creation",
			mockDbConnection: func() (*sql.DB, sqlmock.Sqlmock) {
				db, mock, _ := sqlmock.New()
				mock.ExpectPing().WillReturnError(sql.ErrTxDone)
				return db, mock
			},
			expectedRepo: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			db, mock := tt.mockDbConnection()
			defer db.Close()

			repo := NewRepository(db)

			if tt.expectedRepo {
				assert.NotNil(t, repo)
			} else {
				assert.Nil(t, repo)
			}

			err := mock.ExpectationsWereMet()
			assert.NoError(t, err)
		})
	}
}
